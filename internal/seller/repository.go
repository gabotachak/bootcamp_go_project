package seller

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
)

// Repository encapsulates the storage of a Seller.
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Seller, error)
	Get(ctx context.Context, id int) (domain.Seller, error)
	Save(ctx context.Context, s domain.Seller) (int, error)
	Update(ctx context.Context, s domain.Seller) error
	Delete(ctx context.Context, id int) error
	Exists(ctx context.Context, cid int) bool
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

// SQL Queries
const (
	GET_ALL = `SELECT * FROM sellers`
	GET_ONE = `SELECT * FROM sellers WHERE id=?;`
	SAVE    = `INSERT INTO sellers(cid,company_name,address,telephone,locality_id) VALUES (?,?,?,?,?)`
	UPDATE  = `UPDATE sellers SET cid=?, company_name=?, address=?, telephone=?, locality_id=?  WHERE id=?`
	DELETE  = `DELETE FROM sellers WHERE id=?`
	EXISTS  = `SELECT cid FROM mydb.sellers WHERE cid=?;`
)

func (r *repository) GetAll(ctx context.Context) ([]domain.Seller, error) {
	rows, err := r.db.Query(GET_ALL)
	if err != nil {
		return nil, err
	}

	var sellers []domain.Seller

	for rows.Next() {
		s := domain.Seller{}
		_ = rows.Scan(&s.ID, &s.CID, &s.CompanyName, &s.Address, &s.Telephone, &s.LocalityID)
		sellers = append(sellers, s)
	}

	return sellers, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Seller, error) {

	sqlStatement := GET_ONE
	row := r.db.QueryRow(sqlStatement, id)
	s := domain.Seller{}
	err := row.Scan(&s.ID, &s.CID, &s.CompanyName, &s.Address, &s.Telephone, &s.LocalityID)
	fmt.Println(err)
	if err != nil {
		return domain.Seller{}, err
	}

	return s, nil
}

func (r *repository) Save(ctx context.Context, s domain.Seller) (int, error) {

	stmt, err := r.db.Prepare(SAVE)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(s.CID, s.CompanyName, s.Address, s.Telephone, s.LocalityID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, s domain.Seller) error {
	stmt, err := r.db.Prepare(UPDATE)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(s.CID, s.CompanyName, s.Address, s.Telephone, s.LocalityID, s.ID)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect < 1 {
		return errors.New("seller not found")
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare(DELETE)
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
		return errors.New("seller not found")
	}

	return nil
}

func (r *repository) Exists(ctx context.Context, cid int) bool {
	sqlStatement := EXISTS
	row := r.db.QueryRow(sqlStatement, cid)
	err := row.Scan(&cid)

	return err == nil
}
