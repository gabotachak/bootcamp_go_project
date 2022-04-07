package carrier

import (
	"context"
	"database/sql"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
)

// Repository encapsulates the storage of a warehouse.
type Repository interface {
	GetAllReport(ctx context.Context) ([]CarriersByLocality, error)
	GetReportDetails(ctx context.Context, idLocality int) ([]domain.Carrier, error)
	GetReport(ctx context.Context, idLocality int) (CarriersByLocality, error)
	LocalityExists(ctx context.Context, id int) bool
	Store(ctx context.Context, w domain.Carrier) (domain.Carrier, error)
	CIDExists(ctx context.Context, cid string) bool
}

type repository struct {
	db *sql.DB
}

type CarriersByLocality struct {
	LocalityId    string `json:"locality_id"`
	LocalityName  string `json:"locality_name"`
	CarriersCount string `json:"carriers_count"`
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

/**
Methods in the repository layer.
In the repository we communicate with the database, in each method we consult different resources.
**/
const (
	STORE             = `INSERT INTO carriers(cid, company_name, address, telephone, locality_id) VALUES (?,?,?,?,?);`
	GET_REPORT_DETAIL = `SELECT * FROM carriers WHERE locality_id=?;`
	GET_ALL_REPORT    = `SELECT localities.locality_name, carriers.locality_id, count(*) as carriers_count
		FROM localities inner join carriers 
		on localities.id = carriers.locality_id
		group by locality_id;`
	GET_REPORT = `SELECT localities.locality_name, carriers.locality_id , count(*) as carriers_count
		FROM localities inner join carriers 
		on localities.id = carriers.locality_id
		where locality_id = ?
		group by locality_id;`
	LOCALITY_EXIST = `SELECT id FROM localities WHERE id=?;`
	CID_EXIST      = `SELECT id FROM carriers WHERE cid=?;`
)

func (r *repository) GetAllReport(ctx context.Context) ([]CarriersByLocality, error) {
	rows, err := r.db.Query(GET_ALL_REPORT)
	if err != nil {
		return nil, err
	}

	var localities []CarriersByLocality

	for rows.Next() {
		l := CarriersByLocality{}
		_ = rows.Scan(&l.LocalityName, &l.LocalityId, &l.CarriersCount)
		localities = append(localities, l)
	}
	return localities, nil
}

func (r *repository) GetReportDetails(ctx context.Context, id int) ([]domain.Carrier, error) {
	var carriers []domain.Carrier
	rows, err := r.db.Query(GET_REPORT_DETAIL, id)
	if err != nil {
		return []domain.Carrier{}, err
	}
	for rows.Next() {
		c := domain.Carrier{}
		if err := rows.Scan(&c.ID, &c.CID, &c.CompanyName, &c.Address, &c.Telephone, &c.LocalityId); err != nil {
			return []domain.Carrier{}, err
		}
		carriers = append(carriers, c)
	}
	return carriers, nil
}

func (r *repository) GetReport(ctx context.Context, id int) (CarriersByLocality, error) {
	rows, err := r.db.Query(GET_REPORT, id)
	if err != nil {
		return CarriersByLocality{}, err
	}
	var locality CarriersByLocality
	for rows.Next() {
		if err := rows.Scan(&locality.LocalityName, &locality.LocalityId, &locality.CarriersCount); err != nil {
			return CarriersByLocality{}, err
		}
	}
	return locality, nil
}

func (r *repository) Store(ctx context.Context, c domain.Carrier) (domain.Carrier, error) {
	stmt, err := r.db.Prepare(STORE)
	if err != nil {
		return domain.Carrier{}, err
	}

	res, err := stmt.Exec(&c.CID, &c.CompanyName, &c.Address, &c.Telephone, &c.LocalityId)
	if err != nil {
		return domain.Carrier{}, err
	}

	id, er := res.LastInsertId()
	if er != nil {
		return domain.Carrier{}, err
	}
	c.ID = int(id)

	return c, nil
}

func (r *repository) LocalityExists(ctx context.Context, id int) bool {
	sqlStatement := LOCALITY_EXIST
	row := r.db.QueryRow(sqlStatement, id)
	err := row.Scan(&id)
	return err == nil
}

func (r *repository) CIDExists(ctx context.Context, cid string) bool {
	sqlStatement := CID_EXIST
	row := r.db.QueryRow(sqlStatement, cid)
	err := row.Scan(&cid)

	return err == nil

}
