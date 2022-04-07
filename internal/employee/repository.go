package employee

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
)

// Repository encapsulates the storage of a employee.
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Employee, error)
	Get(ctx context.Context, cardNumberID string) (domain.Employee, error)
	Exists(ctx context.Context, cardNumberID string) bool
	Save(ctx context.Context, e domain.Employee) (int, error)
	Update(ctx context.Context, e domain.Employee) error
	Delete(ctx context.Context, cardNumberID string) error
	GetInboundOrders(ctx context.Context) ([]employeeInboundOrders, error)
	GetInboundOrdersByEmployee(ctx context.Context, id int) (employeeInboundOrders, error)
}

type repository struct {
	db *sql.DB
}
type employeeInboundOrders struct {
	ID                 int    `json:"id"`
	CardNumberID       string `json:"card_number_id"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	WarehouseID        int    `json:"warehouse_id"`
	InboundOrdersCount int    `json:"inbound_orders_count"`
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Employee, error) {
	rows, err := r.db.Query(`SELECT * FROM employees`)
	if err != nil {
		return nil, err
	}

	var employees []domain.Employee

	for rows.Next() {
		e := domain.Employee{}
		_ = rows.Scan(&e.ID, &e.CardNumberID, &e.FirstName, &e.LastName, &e.WarehouseID)
		employees = append(employees, e)
	}

	return employees, nil
}

func (r *repository) Get(ctx context.Context, cardNumberID string) (domain.Employee, error) {

	sqlStatement := `SELECT * FROM employees WHERE card_number_id=?;`
	row := r.db.QueryRow(sqlStatement, cardNumberID)
	e := domain.Employee{}
	err := row.Scan(&e.ID, &e.CardNumberID, &e.FirstName, &e.LastName, &e.WarehouseID)
	if err != nil {
		return domain.Employee{}, err
	}

	return e, nil
}

func (r *repository) Exists(ctx context.Context, cardNumberID string) bool {
	sqlStatement := `SELECT card_number_id FROM employees WHERE card_number_id=?;`
	row := r.db.QueryRow(sqlStatement, cardNumberID)
	err := row.Scan(&cardNumberID)
	if err != nil {
		return false

	}
	return true
}

func (r *repository) Save(ctx context.Context, e domain.Employee) (int, error) {

	//validation if warehouse_id exist
	sqlStatement := `SELECT id FROM warehouses WHERE id = ?`
	row := r.db.QueryRow(sqlStatement, &e.WarehouseID)
	err := row.Scan(&e.WarehouseID)
	if err != nil {
		return 0, err
	}

	stmt, err := r.db.Prepare(`INSERT INTO employees(card_number_id,first_name,last_name,warehouse_id) VALUES (?,?,?,?)`)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&e.CardNumberID, &e.FirstName, &e.LastName, &e.WarehouseID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, e domain.Employee) error {
	//validation if warehouse_id exist
	sqlStatement := `SELECT id FROM warehouses WHERE id = ?`
	row := r.db.QueryRow(sqlStatement, &e.WarehouseID)
	err := row.Scan(&e.WarehouseID)
	if err != nil {
		return err
	}
	stmt, err := r.db.Prepare(`UPDATE employees SET first_name=?, last_name=?, warehouse_id=?  WHERE card_number_id=?`)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&e.FirstName, &e.LastName, &e.WarehouseID, &e.CardNumberID)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect < 1 {
		return errors.New("empleado no encontrado")
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, cardNumberID string) error {
	stmt, err := r.db.Prepare(`DELETE FROM employees WHERE card_number_id=?`)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(cardNumberID)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect < 1 {
		return errors.New("empleado no encontrado")
	}

	return nil
}

func (r *repository) GetInboundOrders(ctx context.Context) ([]employeeInboundOrders, error) {
	rows, err := r.db.Query(`
	SELECT e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id, count(io.id)
	FROM employees AS e
	LEFT JOIN mydb.inbound_orders as io
	ON io.employee_id = e.id
	GROUP BY e.id;`)
	if err != nil {
		return nil, err
	}

	var employees []employeeInboundOrders

	for rows.Next() {
		e := employeeInboundOrders{}
		_ = rows.Scan(&e.ID, &e.CardNumberID, &e.FirstName, &e.LastName, &e.WarehouseID, &e.InboundOrdersCount)
		employees = append(employees, e)
	}

	return employees, nil
}
func (r *repository) GetInboundOrdersByEmployee(ctx context.Context, id int) (employeeInboundOrders, error) {
	sqlStatement := `
	SELECT e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id, count(io.id)
	FROM employees AS e
	LEFT JOIN inbound_orders as io
	ON io.employee_id = e.id
	WHERE io.employee_id = ?
	GROUP BY e.id;`
	fmt.Println(id)
	row := r.db.QueryRow(sqlStatement, id)
	e := employeeInboundOrders{}
	err := row.Scan(&e.ID, &e.CardNumberID, &e.FirstName, &e.LastName, &e.WarehouseID, &e.InboundOrdersCount)
	if err != nil {
		return employeeInboundOrders{}, err
	}

	return e, nil
}
