package buyer

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
)

// Repository encapsulates the storage of a buyer.
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Buyer, error)
	Get(ctx context.Context, cardNumberID string) (domain.Buyer, error)
	Exists(ctx context.Context, cardNumberID string) bool
	Save(ctx context.Context, b domain.Buyer) (int, error)
	Update(ctx context.Context, b domain.Buyer) error
	Delete(ctx context.Context, cardNumberID string) error
	GetPurchaseOrders(ctx context.Context) ([]BuyerPurchaseOrders, error)
	GetPurchaseOrdersByBuyer(ctx context.Context, id int) (BuyerPurchaseOrders, error)
}

type repository struct {
	db *sql.DB
}

type BuyerPurchaseOrders struct {
	ID                  int    `json:"id"`
	CardNumberID        string `json:"card_number_id"`
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	PurchaseOrdersCount int    `json:"Purchase_orders_count"`
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Buyer, error) {
	rows, err := r.db.Query(`SELECT * FROM buyers`)
	if err != nil {
		return nil, err
	}

	var buyers []domain.Buyer

	for rows.Next() {
		b := domain.Buyer{}
		_ = rows.Scan(&b.ID, &b.CardNumberID, &b.FirstName, &b.LastName)
		buyers = append(buyers, b)
	}

	return buyers, nil
}

func (r *repository) Get(ctx context.Context, cardNumberID string) (domain.Buyer, error) {

	sqlStatement := `SELECT * FROM buyers WHERE id_card_number=?;`
	row := r.db.QueryRow(sqlStatement, cardNumberID)
	b := domain.Buyer{}
	err := row.Scan(&b.ID, &b.CardNumberID, &b.FirstName, &b.LastName)
	if err != nil {
		return domain.Buyer{}, errors.New("id del buyer no existe")
	}

	return b, nil
}

func (r *repository) Exists(ctx context.Context, cardNumberID string) bool {
	sqlStatement := `SELECT id_card_number FROM buyers WHERE id_card_number=?;`
	row := r.db.QueryRow(sqlStatement, cardNumberID)
	err := row.Scan(&cardNumberID)
	if err != nil {
		return false
	}
	return true
}

func (r *repository) Save(ctx context.Context, b domain.Buyer) (int, error) {

	if len(b.CardNumberID) <= 0 {
		return 0, errors.New("el CardNumberId está vacio")
	}

	stmt, err := r.db.Prepare(`INSERT INTO buyers(id_card_number,first_name,last_name) VALUES (?,?,?)`)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&b.CardNumberID, &b.FirstName, &b.LastName)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, b domain.Buyer) error {
	stmt, err := r.db.Prepare(`UPDATE buyers SET first_name=?, last_name=?  WHERE id_card_number=?`)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&b.FirstName, &b.LastName, &b.CardNumberID)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect < 1 {
		return errors.New("el CardNumberId no existe")
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, cardNumberID string) error {
	stmt, err := r.db.Prepare(`DELETE FROM buyers WHERE id_card_number=?`)
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
		return errors.New("el CardNumberId no existe")
	}

	return nil
}

func (r *repository) GetPurchaseOrders(ctx context.Context) ([]BuyerPurchaseOrders, error) {
	rows, err := r.db.Query(`
	SELECT b.id, b.id_card_number, b.first_name, b.last_name, count(p.id)
	FROM buyers AS b
	LEFT JOIN purchase_orders as p
	ON p.buyer_id = b.id
	GROUP BY b.id;`)
	if err != nil {
		return nil, err
	}

	var buyers []BuyerPurchaseOrders

	for rows.Next() {
		b := BuyerPurchaseOrders{}
		_ = rows.Scan(&b.ID, &b.CardNumberID, &b.FirstName, &b.LastName, &b.PurchaseOrdersCount)
		buyers = append(buyers, b)
	}

	return buyers, nil
}
func (r *repository) GetPurchaseOrdersByBuyer(ctx context.Context, id int) (BuyerPurchaseOrders, error) {
	sqlStatement := `
	SELECT b.id, b.id_card_number, b.first_name, b.last_name,count(p.id)
	FROM buyers AS b
	LEFT JOIN purchase_orders as p
	ON p.buyer_id = b.id
	WHERE p.buyer_id = ?
	GROUP BY b.id;`
	fmt.Println(id)
	row := r.db.QueryRow(sqlStatement, id)
	b := BuyerPurchaseOrders{}
	err := row.Scan(&b.ID, &b.CardNumberID, &b.FirstName, &b.LastName, &b.PurchaseOrdersCount)
	if err != nil {
		return BuyerPurchaseOrders{}, errors.New("el CardNumberId ingresado no existe")
	}

	return b, nil
}
