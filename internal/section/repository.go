package section

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
)

const (
	SECTIONS_GET_ALL_QUERY   = `SELECT id, section_number, current_temperature, minimun_temperature, current_capacity, minimun_capacity, maximun_capacity, warehouse_id, product_type_id FROM sections`
	SECTIONS_GET_BY_ID_QUERY = `SELECT id, section_number, current_temperature, minimun_temperature, current_capacity, minimun_capacity, maximun_capacity, warehouse_id, product_type_id FROM sections WHERE id=?;`
	SECTIONS_EXISTS_QUERY    = `SELECT section_number FROM sections WHERE section_number=?;`
	SECTIONS_SAVE_QUERY      = `INSERT INTO sections (section_number, current_temperature, minimun_temperature, current_capacity, minimun_capacity, maximun_capacity, warehouse_id, product_type_id) VALUES (?,?,?,?,?,?,?,?);`
	SECTIONS_UPDATE_QUERY    = `UPDATE sections SET section_number=?, current_temperature=?, minimun_temperature=?, current_capacity=?, minimun_capacity=?, maximun_capacity=?, warehouse_id=?, product_type_id=?  WHERE id=?`
	SECTIONS_DELETE_QUERY    = `DELETE FROM sections WHERE id=?`
)

// Repository encapsulates the storage of a section.
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Section, error)
	Get(ctx context.Context, id int) (domain.Section, error)
	Exists(ctx context.Context, sectionNumber string) bool
	Save(ctx context.Context, s domain.Section) (int, error)
	Update(ctx context.Context, s domain.Section) error
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Section, error) {
	rows, err := r.db.Query(SECTIONS_GET_ALL_QUERY)
	if err != nil {
		return nil, err
	}

	var sections []domain.Section

	for rows.Next() {
		s := domain.Section{}
		_ = rows.Scan(&s.ID, &s.SectionNumber, &s.CurrentTemperature, &s.MinimumTemperature, &s.CurrentCapacity, &s.MinimumCapacity, &s.MaximumCapacity, &s.WarehouseID, &s.ProductTypeID)
		sections = append(sections, s)
	}

	return sections, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Section, error) {

	fmt.Println(id)
	sqlStatement := SECTIONS_GET_BY_ID_QUERY
	row := r.db.QueryRow(sqlStatement, id)
	s := domain.Section{}
	err := row.Scan(&s.ID, &s.SectionNumber, &s.CurrentTemperature, &s.MinimumTemperature, &s.CurrentCapacity, &s.MinimumCapacity, &s.MaximumCapacity, &s.WarehouseID, &s.ProductTypeID)
	if err != nil {
		fmt.Print(err)
		return domain.Section{}, err
	}

	return s, nil
}

func (r *repository) Exists(ctx context.Context, sectionNumber string) bool {
	sqlStatement := SECTIONS_EXISTS_QUERY
	row := r.db.QueryRow(sqlStatement, sectionNumber)
	err := row.Scan(&sectionNumber)
	return err == nil
}

func (r *repository) Save(ctx context.Context, s domain.Section) (int, error) {

	stmt, err := r.db.Prepare(SECTIONS_SAVE_QUERY)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&s.SectionNumber, &s.CurrentTemperature, &s.MinimumTemperature, &s.CurrentCapacity, &s.MinimumCapacity, &s.MaximumCapacity, &s.WarehouseID, &s.ProductTypeID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, s domain.Section) error {
	stmt, err := r.db.Prepare(SECTIONS_UPDATE_QUERY)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&s.SectionNumber, &s.CurrentTemperature, &s.MinimumTemperature, &s.CurrentCapacity, &s.MinimumCapacity, &s.MaximumCapacity, &s.WarehouseID, &s.ProductTypeID, &s.ID)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect < 1 {
		return errors.New("section not found")
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare(SECTIONS_DELETE_QUERY)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect < 1 {
		return errors.New("section not found")
	}

	return nil
}
