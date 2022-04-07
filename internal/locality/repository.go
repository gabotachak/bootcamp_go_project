package locality

import (
	"context"
	"database/sql"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
)

// Repository encapsulates the storage of a Locality.
type Repository interface {
	GetGeneralReport(ctx context.Context) ([]LocalityReport, error)
	GetReport(ctx context.Context, id int) (LocalityReport, error)
	Save(ctx context.Context, s domain.Locality) (domain.Locality, error)
}

type repository struct {
	db *sql.DB
}

type LocalityReport struct {
	LocalityId   int    `json:"locality_id"`
	LocalityName string `json:"locality_name"`
	Quantity     int    `json:"cantidad"`
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

const (
	GET_ONE = `SELECT localities.id, localities.locality_name, provinces.province_name, countries.country_name 
			FROM localities 
			INNER JOIN provinces 
			ON localities.province_id=provinces.id 
			INNER JOIN countries 
			ON provinces.country_id=countries.id
			WHERE localities.id=?;`
	SAVE_ONE = `INSERT INTO localities 
			SET locality_name = ?,
			province_id = (
			SELECT id
		  	FROM provinces 
		 	WHERE province_name = ?)`
	GET_GENERAL_REPORT = `SELECT s.locality_id, l.locality_name, COUNT(s.id) as Cantidad
			FROM sellers s
			INNER JOIN localities l 
			ON s.locality_id = l.id 
			GROUP BY s.locality_id`
	GET_REPORT = `SELECT s.locality_id, l.locality_name, COUNT(s.id) as Cantidad
			FROM sellers s
			INNER JOIN localities l 
			ON s.locality_id = l.id 
			WHERE s.locality_id = ?
			GROUP BY s.locality_id`
)

/**
* Methods in the repository layer.
* In the repository we communicate with the database, in each method we consult different resources.
**/
func (r *repository) GetGeneralReport(ctx context.Context) ([]LocalityReport, error) {
	rows, err := r.db.Query(GET_GENERAL_REPORT)

	if err != nil {
		return nil, err
	}

	var localities []LocalityReport

	for rows.Next() {
		s := LocalityReport{}
		_ = rows.Scan(&s.LocalityId, &s.LocalityName, &s.Quantity)
		localities = append(localities, s)
	}

	return localities, nil
}

func (r *repository) GetReport(ctx context.Context, id int) (LocalityReport, error) {
	sqlStatement := GET_REPORT
	row := r.db.QueryRow(sqlStatement, id)
	l := LocalityReport{}
	err := row.Scan(&l.LocalityId, &l.LocalityName, &l.Quantity)
	if err != nil {
		return LocalityReport{}, err
	}

	return l, nil
}

func (r *repository) Save(ctx context.Context, s domain.Locality) (domain.Locality, error) {

	stmt, err := r.db.Prepare(SAVE_ONE)
	if err != nil {
		return domain.Locality{}, err
	}

	res, err := stmt.Exec(&s.LocalityName, &s.ProvinceName)
	if err != nil {
		return domain.Locality{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return domain.Locality{}, err
	}
	s.ID = int(id)
	return s, nil
}
