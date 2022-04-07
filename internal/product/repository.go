package product

import (
	"context"
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
)

const (
	GetAllSqlProduct = `SELECT * FROM products`
	GetSqlProduct    = `SELECT * FROM products WHERE id=?;`
	SaveSqlProduct   = `INSERT INTO products(description,expiration_rate,freezing_rate,height,length,net_weight,product_code,recommended_freezing_temperature,width,product_type_id,seller_id) VALUES (?,?,?,?,?,?,?,?,?,?,?)`
	ExistsSqlProduct = `SELECT product_code FROM products WHERE product_code=?;`
	UpdateSqlProduct = `UPDATE products SET description=?, expiration_rate=?, freezing_rate=?, height=?, length=?, net_weight=?, product_code=?, recommended_freezing_temperature=?, width=?, product_type_id=?, seller_id=?  WHERE id=?`
	DeleteSqlProduct = `DELETE FROM products WHERE id=?`
)

// Repository encapsulates the storage of a Product.
type Repository interface {
	GetAll() ([]domain.Product, error)
	Get(ctx context.Context, id int) (domain.Product, error)
	Exists(ctx context.Context, productCode string) bool
	Save(ctx context.Context, p domain.Product) (int, error)
	Update(ctx context.Context, p domain.Product) error
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

//Query to database to get all Products
func (r *repository) GetAll() ([]domain.Product, error) {
	rows, err := r.db.Query(GetAllSqlProduct)
	if err != nil {
		return nil, err
	}

	var products []domain.Product

	for rows.Next() {
		p := domain.Product{}
		_ = rows.Scan(&p.ID, &p.Description, &p.ExpirationRate, &p.FreezingRate, &p.Height, &p.Length, &p.Netweight, &p.ProductCode, &p.RecomFreezTemp, &p.Width, &p.ProductTypeID, &p.SellerID)
		products = append(products, p)
	}

	return products, nil
}

//Query to database to get a specific Product and return it
func (r *repository) Get(ctx context.Context, id int) (domain.Product, error) {

	sqlStatement := GetSqlProduct
	row := r.db.QueryRow(sqlStatement, id)
	p := domain.Product{}
	err := row.Scan(&p.ID, &p.Description, &p.ExpirationRate, &p.FreezingRate, &p.Height, &p.Length, &p.Netweight, &p.ProductCode, &p.RecomFreezTemp, &p.Width, &p.ProductTypeID, &p.SellerID)
	if err != nil {
		return domain.Product{}, err
	}

	return p, nil
}

//Query to database to verify if a product code exists
func (r *repository) Exists(ctx context.Context, productCode string) bool {
	sqlStatement := ExistsSqlProduct
	row := r.db.QueryRow(sqlStatement, productCode)
	err := row.Scan(&productCode)
	if err != nil {
		return false
	}
	return true
}

//Query to database to save a product and return the id
func (r *repository) Save(ctx context.Context, p domain.Product) (int, error) {

	stmt, err := r.db.Prepare(SaveSqlProduct)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(p.Description, p.ExpirationRate, p.FreezingRate, p.Height, p.Length, p.Netweight, p.ProductCode, p.RecomFreezTemp, p.Width, p.ProductTypeID, p.SellerID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

//Query to database - Update a Product, and return an error if something goes wrong or doesnt exist
func (r *repository) Update(ctx context.Context, p domain.Product) error {
	stmt, err := r.db.Prepare(UpdateSqlProduct)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(p.Description, p.ExpirationRate, p.FreezingRate, p.Height, p.Length, p.Netweight, p.ProductCode, p.RecomFreezTemp, p.Width, p.ProductTypeID, p.SellerID, p.ID)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect < 1 {
		return errors.New("product not found")
	}

	return nil
}

//Query to database to delete a Product and return nil or error if something goes wrong
func (r *repository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare(DeleteSqlProduct)
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
		return errors.New("product not found")
	}

	return nil
}
