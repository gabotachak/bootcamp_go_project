package productRecord

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
)

const (
	layoutISO           = "2006-01-02"
	ExistIdSqlStatement = `SELECT * FROM products WHERE id=?;`
	GetAllSqlStatement  = `SELECT pr.product_id, p.description , count(*) as records_count FROM products p inner join product_records pr on p.id = pr.product_id group by pr.product_id;`
	GetsqlStatement     = `SELECT pr.product_id, p.description , count(*) as records_count FROM products p inner join product_records pr on p.id = pr.product_id WHERE product_id = ? group by pr.product_id;`
	InsertSqlStatement  = `INSERT INTO product_records(last_update_date,purchase_price,sale_price,product_id) VALUES (?,?,?,?)`
)

// Repository encapsulates the storage of a Product.
type Repository interface {
	GetAll(ctx context.Context) ([]ProductRecordByProduct, error)
	Get(ctx context.Context, id int) (ProductRecordByProduct, error)
	Save(ctx context.Context, pr domain.ProductRecord) (ProdRecSave, error)
	ExistId(ctx context.Context, id int) bool
	ValidateDate(ctx context.Context, myDate string) bool
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

type ProductRecordByProduct struct {
	ProdID          int    `json:"product_id"`
	ProdDescription string `json:"description"`
	Records_Count   int    `json:"records_count"`
}

type ProdRecSave struct {
	LastUpdateDate string  `json:"last_update_date"`
	PurchasePrice  float32 `json:"purchase_price"`
	SalePrice      float32 `json:"sale_price"`
	ProductId      int     `json:"product_id"`
}

// This method ask for all product records grouped by product id, and returns Description, ID and records count of each of them
func (r *repository) GetAll(ctx context.Context) ([]ProductRecordByProduct, error) {
	rows, err := r.db.Query(GetAllSqlStatement)
	if err != nil {
		return nil, err
	}
	var prodRecords []ProductRecordByProduct

	for rows.Next() {
		oneRecord := ProductRecordByProduct{}
		_ = rows.Scan(&oneRecord.ProdID, &oneRecord.ProdDescription, &oneRecord.Records_Count)
		prodRecords = append(prodRecords, oneRecord)
	}
	return prodRecords, nil
}

//This method ask for a specific product id, and return Description, ID and records count asociated with it
func (r *repository) Get(ctx context.Context, id int) (ProductRecordByProduct, error) {

	rows, err := r.db.Query(GetsqlStatement, id)
	var prodRecords ProductRecordByProduct
	if err != nil {
		return ProductRecordByProduct{}, err
	}
	for rows.Next() {
		if err := rows.Scan(&prodRecords.ProdID, &prodRecords.ProdDescription, &prodRecords.Records_Count); err != nil {
			log.Println(err.Error())
			return ProductRecordByProduct{}, err
		}
	}
	return prodRecords, nil
}

//This method saves a Product Record asociated whit an existing Product, and returns the product record
func (r *repository) Save(ctx context.Context, pr domain.ProductRecord) (ProdRecSave, error) {
	if r.ValidateDate(ctx, pr.LastUpdateDate) {
		stmt, err := r.db.Prepare(InsertSqlStatement)
		if err != nil {
			return ProdRecSave{}, err
		}

		res, err := stmt.Exec(pr.LastUpdateDate, pr.PurchasePrice, pr.SalePrice, pr.ProductId)
		if err != nil {
			return ProdRecSave{}, err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return ProdRecSave{}, err
		}
		fmt.Println(id)
		prodRSave := ProdRecSave{}
		prodRSave.LastUpdateDate = pr.LastUpdateDate
		prodRSave.PurchasePrice = pr.PurchasePrice
		prodRSave.SalePrice = pr.SalePrice
		prodRSave.ProductId = pr.ProductId
		return prodRSave, nil
	}
	return ProdRecSave{}, errors.New("la fecha debe ser mayor a la actual")
}

// This method verify into the database if a Product id exist and return true if it does
func (r *repository) ExistId(ctx context.Context, id int) bool {
	myProd := domain.Product{}

	row := r.db.QueryRow(ExistIdSqlStatement, id)
	err := row.Scan(&myProd.ID, &myProd.Description, &myProd.ExpirationRate, &myProd.FreezingRate, &myProd.Height, &myProd.Length, &myProd.Netweight, &myProd.ProductCode, &myProd.RecomFreezTemp, &myProd.Width, &myProd.ProductTypeID, &myProd.SellerID)
	return nil == err
}

// This methos validates if the date inserted in a product record is after the current date, return true
func (r *repository) ValidateDate(ctx context.Context, myDate string) bool {
	//get actual date
	now := time.Now()
	//y, m, d := now.Date()
	//convert string to date
	t, _ := time.Parse(layoutISO, myDate)
	//compare dates and return boolean
	return t.After(now)
}
