package inboundOrders

import (
	"context"
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
)

// Repository encapsulates the storage of a employee.
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Inbound_Orders, error)
	Get(ctx context.Context, id int) (domain.Inbound_Orders, error)
	Exists(ctx context.Context, id int) bool
	Save(ctx context.Context, io domain.Inbound_Orders) (domain.Inbound_Orders, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Inbound_Orders, error) {
	rows, err := r.db.Query(`SELECT * FROM inbound_orders`)
	if err != nil {
		return nil, err
	}

	var inbound_orders []domain.Inbound_Orders

	for rows.Next() {
		io := domain.Inbound_Orders{}
		_ = rows.Scan(&io.ID, &io.OrderDate, &io.OrderNumber, &io.EmployeeID, &io.WarehouseID, &io.ProductBatchID)
		inbound_orders = append(inbound_orders, io)
	}

	return inbound_orders, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Inbound_Orders, error) {

	sqlStatement := `SELECT * FROM inbound_orders WHERE id=?;`
	row := r.db.QueryRow(sqlStatement, id)
	io := domain.Inbound_Orders{}
	err := row.Scan(&io.ID, &io.OrderDate, &io.OrderNumber, &io.EmployeeID, &io.WarehouseID, &io.ProductBatchID)
	if err != nil {
		return domain.Inbound_Orders{}, errors.New("inbound order id no encontrado")
	}

	return io, nil
}

func (r *repository) Exists(ctx context.Context, id int) bool {
	sqlStatement := `SELECT id FROM inbound_orders WHERE id=?;`
	row := r.db.QueryRow(sqlStatement, id)
	err := row.Scan(&id)
	if err != nil {
		return false

	}
	return true
}

func (r *repository) Save(ctx context.Context, io domain.Inbound_Orders) (domain.Inbound_Orders, error) {

	//validation if order_number not empty
	if io.OrderNumber == "" {
		return domain.Inbound_Orders{}, errors.New("el order_number esta vacio")
	}

	//validation if employee exist
	sqlStatement := `SELECT id FROM employees WHERE id = ?`
	row := r.db.QueryRow(sqlStatement, &io.EmployeeID)
	err := row.Scan(&io.EmployeeID)
	if err != nil {
		return domain.Inbound_Orders{}, errors.New("el empleado no existe")
	}

	//validation if warehouse exist
	sqlStatement = `SELECT id FROM warehouses WHERE id = ?`
	row = r.db.QueryRow(sqlStatement, &io.WarehouseID)
	err = row.Scan(&io.WarehouseID)
	if err != nil {
		return domain.Inbound_Orders{}, errors.New("el warehouse no existe")
	}

	//validation if product_batch exist
	sqlStatement = `SELECT id FROM product_batches WHERE id = ?`
	row = r.db.QueryRow(sqlStatement, &io.ProductBatchID)
	err = row.Scan(&io.ProductBatchID)
	if err != nil {
		return domain.Inbound_Orders{}, errors.New("el product_batches no existe")
	}

	stmt, err := r.db.Prepare(`INSERT INTO inbound_orders(order_date,order_number,employee_id,warehouse_id,product_batch_id) VALUES (?,?,?,?,?)`)
	if err != nil {
		return domain.Inbound_Orders{}, err
	}

	res, err := stmt.Exec(&io.OrderDate, &io.OrderNumber, &io.EmployeeID, &io.WarehouseID, &io.ProductBatchID)
	if err != nil {
		return domain.Inbound_Orders{}, err
	}
	rows_affected, err := res.RowsAffected()
	if rows_affected == 0 {
		return domain.Inbound_Orders{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return domain.Inbound_Orders{}, err
	}

	io.ID = int(id)

	return io, nil
}
