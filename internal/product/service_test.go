package product

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	db "github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/pkg/database"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

type repository_test struct {
	db   *sql.DB
	flag bool
}

//Send hardcoded retuns in method "GetAll" for unitary test on service layer
func (r *repository_test) GetAll() ([]domain.Product, error) {
	r.flag = true
	rows := []domain.Product{{
		ID:             1,
		Description:    "Azucar",
		ExpirationRate: 20,
		FreezingRate:   -4,
		Height:         8.3,
		Length:         2.7,
		Netweight:      1,
		ProductCode:    "Prod1",
		RecomFreezTemp: -3,
		Width:          2.8,
		ProductTypeID:  3,
		SellerID:       8,
	}, {
		ID:             2,
		Description:    "Polenta",
		ExpirationRate: 1,
		FreezingRate:   5,
		Height:         5.3,
		Length:         4.5,
		Netweight:      1,
		ProductCode:    "Prod2",
		RecomFreezTemp: 1,
		Width:          2.8,
		ProductTypeID:  4,
		SellerID:       4,
	}}

	return rows, nil
}

//Send hardcoded retuns in method "Get" for unitary test on service layer
func (r *repository_test) Get(ctx context.Context, id int) (domain.Product, error) {

	r.flag = true
	if id == 1 {
		prod := domain.Product{
			ID:             1,
			Description:    "Azucar",
			ExpirationRate: 20,
			FreezingRate:   -4,
			Height:         8.3,
			Length:         2.7,
			Netweight:      1,
			ProductCode:    "Prod1",
			RecomFreezTemp: -3,
			Width:          2.8,
			ProductTypeID:  3,
			SellerID:       8,
		}
		return prod, nil
	} else {
		return domain.Product{}, errors.New("El id ingresado no existe")
	}
}

//Send hardcoded retuns in method "Exist" for unitary test on service layer
func (r *repository_test) Exists(ctx context.Context, productCode string) bool {
	code := "Prod1"
	r.flag = true
	return productCode == code
}

//Send hardcoded retuns in method "Save" for unitary test on service layer
func (r *repository_test) Save(ctx context.Context, p domain.Product) (int, error) {
	if p.ProductCode == "Prod3" {
		return 0, errors.New("no se puede guardar el producto o el codigo de producto ya existe")
	}
	return 1, nil
}

//Send hardcoded retuns in method "Update" for unitary test on service layer
func (r *repository_test) Update(ctx context.Context, p domain.Product) error {
	r.flag = true
	if p.ID == 3 {
		return errors.New("Algo no anduvo bien")
	}
	return nil
}

//Send hardcoded retuns in method "delete" for unitary test on service layer
func (r *repository_test) Delete(ctx context.Context, id int) error {
	r.flag = true
	if id != 3 {
		return errors.New("No existe el ID")
	}
	return nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////
//								Methods to test the service layer
///////////////////////////////////////////////////////////////////////////////////////////////////

//Create test database
func createDataBase() (*sql.DB, error) {
	// Method that initializes the database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	return db, nil
}

func TestGetAll(t *testing.T) {
	expectedResult := []domain.Product{{
		ID:             1,
		Description:    "Azucar",
		ExpirationRate: 20,
		FreezingRate:   -4,
		Height:         8.3,
		Length:         2.7,
		Netweight:      1,
		ProductCode:    "Prod1",
		RecomFreezTemp: -3,
		Width:          2.8,
		ProductTypeID:  3,
		SellerID:       8,
	}, {
		ID:             2,
		Description:    "Polenta",
		ExpirationRate: 1,
		FreezingRate:   5,
		Height:         5.3,
		Length:         4.5,
		Netweight:      1,
		ProductCode:    "Prod2",
		RecomFreezTemp: 1,
		Width:          2.8,
		ProductTypeID:  4,
		SellerID:       4,
	}}
	dbTest, err := createDataBase()
	if err != nil {
		fmt.Println(err)
	}
	var repositoryAux = &repository_test{dbTest, false}
	service := NewService(repositoryAux)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	result, err1 := service.GetAll(ctx)
	assert.Equal(t, expectedResult, result, "Deben ser iguales")
	assert.Nil(t, err1)
}

func TestGet(t *testing.T) {

	expectedResult := domain.Product{
		ID:             1,
		Description:    "Azucar",
		ExpirationRate: 20,
		FreezingRate:   -4,
		Height:         8.3,
		Length:         2.7,
		Netweight:      1,
		ProductCode:    "Prod1",
		RecomFreezTemp: -3,
		Width:          2.8,
		ProductTypeID:  3,
		SellerID:       8,
	}

	dbTest, err := createDataBase()
	if err != nil {
		fmt.Println(err)
	}
	var repositoryAux = &repository_test{dbTest, false}
	service := NewService(repositoryAux)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	result, _ := service.Get(ctx, 1)
	assert.Equal(t, expectedResult, result, "Deben ser iguales")
	assert.True(t, repositoryAux.flag)
}

func TestGetNotFound(t *testing.T) {

	expectedResult := errors.New("El id ingresado no existe")

	dbTest, err := createDataBase()
	if err != nil {
		fmt.Println(err)
	}
	var repositoryAux = &repository_test{dbTest, false}
	service := NewService(repositoryAux)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	_, err1 := service.Get(ctx, 3)
	assert.Equal(t, expectedResult, err1, "Deben ser iguales")
	assert.True(t, repositoryAux.flag)
}

func TestDelete(t *testing.T) {
	dbTest, err := createDataBase()
	if err != nil {
		fmt.Println(err)
	}
	var repositoryAux = &repository_test{dbTest, false}
	service := NewService(repositoryAux)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	result := service.Delete(ctx, 3)
	assert.Equal(t, nil, result, "Deben ser iguales")
}

func TestDeleteIdNotFound(t *testing.T) {
	dbTest, err := createDataBase()
	if err != nil {
		fmt.Println(err)
	}
	expectedResult := errors.New("No existe el ID")
	var repositoryAux = &repository_test{dbTest, false}
	service := NewService(repositoryAux)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	result := service.Delete(ctx, 1)
	assert.Equal(t, expectedResult, result, "No existe el ID")
	assert.True(t, repositoryAux.flag)
}

func TestSave(t *testing.T) {
	dbTest, err := createDataBase()
	if err != nil {
		fmt.Println(err)
	}
	prod := domain.Product{
		ID:             1,
		Description:    "Azucar",
		ExpirationRate: 20,
		FreezingRate:   -4,
		Height:         8.3,
		Length:         2.7,
		Netweight:      1,
		ProductCode:    "Prod2",
		RecomFreezTemp: -3,
		Width:          2.8,
		ProductTypeID:  3,
		SellerID:       8,
	}
	var repositoryAux = &repository_test{dbTest, false}
	service := NewService(repositoryAux)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	result, _ := service.Save(ctx, prod)
	assert.Equal(t, 1, result, "El id debe ser igual")
	assert.True(t, repositoryAux.flag)
}

func TestSaveError(t *testing.T) {
	dbTest, err := createDataBase()
	if err != nil {
		fmt.Println(err)
	}
	prod := domain.Product{
		ID:             1,
		Description:    "Azucar",
		ExpirationRate: 20,
		FreezingRate:   -4,
		Height:         8.3,
		Length:         2.7,
		Netweight:      1,
		ProductCode:    "Prod1",
		RecomFreezTemp: -3,
		Width:          2.8,
		ProductTypeID:  3,
		SellerID:       8,
	}
	var repositoryAux = &repository_test{dbTest, false}
	service := NewService(repositoryAux)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	expectedResult := errors.New("no se puede guardar el producto o el codigo de producto ya existe")
	_, err1 := service.Save(ctx, prod)
	assert.Equal(t, expectedResult, err1, "El id debe ser igual")
	assert.True(t, repositoryAux.flag)
}

func TestUpdate(t *testing.T) {
	dbTest, err := createDataBase()
	if err != nil {
		fmt.Println(err)
	}
	prod := domain.Product{
		ID:             1,
		Description:    "Azucar",
		ExpirationRate: 20,
		FreezingRate:   -4,
		Height:         8.3,
		Length:         2.7,
		Netweight:      1,
		ProductCode:    "Prod1",
		RecomFreezTemp: -3,
		Width:          2.8,
		ProductTypeID:  3,
		SellerID:       8,
	}
	var repositoryAux = &repository_test{dbTest, false}
	service := NewService(repositoryAux)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	result := service.Update(ctx, prod)
	assert.Equal(t, nil, result, "El id debe ser igual")
	assert.True(t, repositoryAux.flag)
}

func TestUpdateNotFound(t *testing.T) {
	dbTest, err := createDataBase()
	if err != nil {
		fmt.Println(err)
	}
	prod := domain.Product{
		ID:             3,
		Description:    "Azucar",
		ExpirationRate: 20,
		FreezingRate:   -4,
		Height:         8.3,
		Length:         2.7,
		Netweight:      1,
		ProductCode:    "Prod2",
		RecomFreezTemp: -3,
		Width:          2.8,
		ProductTypeID:  3,
		SellerID:       8,
	}
	expectedResult := errors.New("Algo no anduvo bien")
	var repositoryAux = &repository_test{dbTest, false}
	service := NewService(repositoryAux)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	result := service.Update(ctx, prod)
	assert.Equal(t, expectedResult, result, "El id es incorrecto")
	assert.True(t, repositoryAux.flag)
}
