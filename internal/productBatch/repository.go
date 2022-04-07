package productBatch

import (
	"context"
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
)

const (
	SAVE_QUERY              = `INSERT INTO product_batches(batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimun_temperature, product_id, section_id) VALUES (?,?,?,?,?,?,?,?,?,?);`
	REPORT_ALL_QUERY        = `SELECT s.id, s.section_number, count(pb.section_id) FROM sections AS s LEFT JOIN product_batches AS pb ON pb.section_id = s.id GROUP BY s.section_number;`
	REPORT_BY_SECTION_QUERY = `SELECT s.id, s.section_number, count(pb.section_id) FROM sections AS s LEFT JOIN product_batches AS pb ON pb.section_id = s.id AND s.id=? GROUP BY s.section_number;`
	BATCH_EXISTS_QUERY      = `SELECT id FROM product_batches WHERE batch_number=?;`
	SECTION_EXISTS_QUERY    = `SELECT id FROM sections WHERE id=?;`
	PRODUCT_EXISTS_QUERY    = `SELECT id FROM products WHERE id=?;`
)

type Repository interface {
	Save(ctx context.Context, pb domain.ProductBatch) (int, error)
	ReportAll(ctx context.Context) ([]domain.ReportProductBatch, error)
	ReportBySection(ctx context.Context, sectionId int) (domain.ReportProductBatch, error)
	BatchExists(ctx context.Context, batchNumber string) bool
	SectionExists(ctx context.Context, sectionId int) bool
	ProductExists(ctx context.Context, productId int) bool
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(ctx context.Context, pb domain.ProductBatch) (int, error) {

	if r.BatchExists(ctx, pb.BatchNumber) {
		return -1, errors.New("product batch already exists")
	}
	if !r.SectionExists(ctx, pb.SectionId) {
		return -1, errors.New("section id does not exist")
	}
	if !r.ProductExists(ctx, pb.ProductId) {
		return -1, errors.New("product id does not exist")
	}
	stmt, err := r.db.Prepare(SAVE_QUERY)
	if err != nil {
		return -1, err
	}

	res, err := stmt.Exec(&pb.BatchNumber, &pb.CurrentQuantity, &pb.CurrentTemperature, &pb.DueDate, &pb.InitialQuantity, &pb.ManufacturingDate, &pb.ManufacturingHour, &pb.MinimumTemperature, &pb.ProductId, &pb.SectionId)
	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(id), nil
}

func (r *repository) ReportAll(ctx context.Context) ([]domain.ReportProductBatch, error) {
	sqlStatement := REPORT_ALL_QUERY
	rows, err := r.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	var reports []domain.ReportProductBatch

	for rows.Next() {
		rpb := domain.ReportProductBatch{}
		_ = rows.Scan(&rpb.SectionId, &rpb.SectionNumber, &rpb.Quantity)
		reports = append(reports, rpb)
	}

	return reports, nil
}

func (r *repository) ReportBySection(ctx context.Context, sectionId int) (domain.ReportProductBatch, error) {

	sqlStatement := REPORT_BY_SECTION_QUERY
	row := r.db.QueryRow(sqlStatement, sectionId)
	rpb := domain.ReportProductBatch{}
	err := row.Scan(&rpb.SectionId, &rpb.SectionNumber, &rpb.Quantity)
	if err != nil {
		return domain.ReportProductBatch{}, err
	}

	return rpb, nil
}

func (r *repository) BatchExists(ctx context.Context, batchNumber string) bool {
	sqlStatement := BATCH_EXISTS_QUERY
	row := r.db.QueryRow(sqlStatement, batchNumber)
	err := row.Scan(&batchNumber)
	return err == nil
}

func (r *repository) SectionExists(ctx context.Context, sectionId int) bool {
	sqlStatement := SECTION_EXISTS_QUERY
	row := r.db.QueryRow(sqlStatement, sectionId)
	err := row.Scan(&sectionId)
	return err == nil
}

func (r *repository) ProductExists(ctx context.Context, productId int) bool {
	sqlStatement := PRODUCT_EXISTS_QUERY
	row := r.db.QueryRow(sqlStatement, productId)
	err := row.Scan(&productId)
	return err == nil
}
