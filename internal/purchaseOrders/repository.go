package purchaseOrders

import (
	"context"
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
)

type Repository interface {
	Get(ctx context.Context, purchaseOrderId int) (domain.PurchaseOrder, error)
	GetAll(ctx context.Context) ([]domain.PurchaseOrder, error)
	Store(ctx context.Context, p domain.PurchaseOrder) (domain.PurchaseOrder, error)
	Exists(ctx context.Context, purchaseOrderId int) bool
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

//Conections with the database
const (
	STORE                 = `INSERT INTO purchase_orders(id, order_number, order_date, tracking_code, buyer_id, order_status_id, carrier_id, warehouse_id) VALUES (?,?,?,?,?,?,?,?)`
	GET                   = `SELECT * FROM purchase_orders WHERE id=?;`
	GET_ALL               = `SELECT * FROM purchase_orders`
	GET_ID_BUYER          = `SELECT id FROM buyers WHERE id=?`
	GET_ID_PURCHASE_ORDER = `SELECT id FROM purchase_orders WHERE id=?`
)

func (r *repository) Get(ctx context.Context, purchaseOrderId int) (domain.PurchaseOrder, error) {
	rows := r.db.QueryRow(GET, purchaseOrderId)

	p := domain.PurchaseOrder{}
	err := rows.Scan(&p.ID, &p.OrderNumber, &p.OrderDate, &p.TrackingCode, &p.BuyerId, &p.OrderStatusId, &p.CarrierId, &p.WarehouseId)
	if err != nil {
		return domain.PurchaseOrder{}, errors.New("id de la orden de compra no fue encontrado")
	}

	return p, nil
}

func (r *repository) GetAll(ctx context.Context) ([]domain.PurchaseOrder, error) {
	rows, err := r.db.Query(GET_ALL)
	if err != nil {
		return nil, err
	}

	var purchaseOrders []domain.PurchaseOrder

	for rows.Next() {
		p := domain.PurchaseOrder{}
		_ = rows.Scan(&p.ID, &p.OrderNumber, &p.OrderDate, &p.TrackingCode, &p.BuyerId, &p.OrderStatusId, &p.CarrierId, &p.WarehouseId)
		purchaseOrders = append(purchaseOrders, p)
	}
	return purchaseOrders, nil
}

func (r *repository) Store(ctx context.Context, p domain.PurchaseOrder) (domain.PurchaseOrder, error) {

	if len(p.OrderNumber) <= 0 {
		return domain.PurchaseOrder{}, errors.New("el número de la orden está vacio")
	}

	//validation if buyer exist
	row := r.db.QueryRow(GET_ID_BUYER, &p.BuyerId)
	err := row.Scan(&p.BuyerId)
	if err != nil {
		return domain.PurchaseOrder{}, errors.New("el comprador no existe")
	}

	stmt, err := r.db.Prepare(STORE)
	if err != nil {
		return domain.PurchaseOrder{}, err
	}

	res, err := stmt.Exec(
		&p.ID,
		&p.OrderNumber,
		&p.OrderDate,
		&p.TrackingCode,
		&p.BuyerId,
		&p.OrderStatusId,
		&p.CarrierId,
		&p.WarehouseId,
	)
	if err != nil {
		return domain.PurchaseOrder{}, err
	}

	id, er := res.LastInsertId()
	if er != nil {
		return domain.PurchaseOrder{}, err
	}
	p.ID = int(id)

	return p, nil
}

func (r *repository) Exists(ctx context.Context, purchaseOrderId int) bool {
	row := r.db.QueryRow(GET_ID_PURCHASE_ORDER, purchaseOrderId)
	err := row.Scan(&purchaseOrderId)
	return err == nil
}
