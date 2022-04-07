package productRecord

import (
	// "context"
	// "database/sql"
	// "errors"
	// "fmt"
	// "net/http/httptest"
	// "testing"

	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	db "github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

type repository_test struct {
	db   *sql.DB
	flag bool
}

//Send hardcoded retuns in method "GetAll" for unitary test on service layer
func (r *repository_test) GetAll(ctx context.Context) ([]ProductRecordByProduct, error) {
	r.flag = true
	rows := []ProductRecordByProduct{{
		ProdID:          1,
		ProdDescription: "prodTest",
		Records_Count:   1,
	}, {
		ProdID:          1,
		ProdDescription: "prodTest2",
		Records_Count:   4,
	}}

	return rows, nil
}

func (r *repository_test) ExistId(ctx context.Context, id int) bool {
	myProd := domain.Product{}
	sqlStatement := `SELECT * FROM products WHERE id=?;`
	row := r.db.QueryRow(sqlStatement, id)
	err := row.Scan(&myProd.ID, &myProd.Description, &myProd.ExpirationRate, &myProd.FreezingRate, &myProd.Height, &myProd.Length, &myProd.Netweight, &myProd.ProductCode, &myProd.RecomFreezTemp, &myProd.Width, &myProd.ProductTypeID, &myProd.SellerID)
	return nil == err
}

//Send hardcoded retuns in method "Get" for unitary test on service layer
func (r *repository_test) Get(ctx context.Context, id int) (ProductRecordByProduct, error) {
	r.flag = true
	if id == 1 {
		prodRec := ProductRecordByProduct{
			ProdID:          1,
			ProdDescription: "prodTest",
			Records_Count:   1,
		}
		return prodRec, nil
	} else {
		return ProductRecordByProduct{}, errors.New("El id ingresado no existe")
	}
}

//Send hardcoded retuns in method "Save" for unitary test on service layer
func (r *repository_test) Save(ctx context.Context, req domain.ProductRecord) (ProdRecSave, error) {
	r.flag = true
	var myProRec ProdRecSave
	myProRec = ProdRecSave{
		LastUpdateDate: "2026-12-12",
		PurchasePrice:  100,
		SalePrice:      200,
		ProductId:      1,
	}
	if req.ProductId != 1 {
		return ProdRecSave{}, errors.New("no se puede guardar, el product id no existe")
	}
	return myProRec, nil
}

func (r *repository_test) ValidateDate(ctx context.Context, myDate string) bool {
	//get actual date
	now := time.Now()
	//y, m, d := now.Date()
	//convert string to date
	t, _ := time.Parse(layoutISO, myDate)
	//compare dates and return boolean
	return t.After(now)
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////
// //								Methods to test the service layer
// ///////////////////////////////////////////////////////////////////////////////////////////////////

//Create test database
func createDataBase() (*sql.DB, error) {
	// Method that initializes the database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	return db, nil
}

func TestGetAllServiceProRecord(t *testing.T) {
	expectedResult := []ProductRecordByProduct{{
		ProdID:          1,
		ProdDescription: "prodTest",
		Records_Count:   1,
	}, {
		ProdID:          1,
		ProdDescription: "prodTest2",
		Records_Count:   4,
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

func TestGetPrdRecordService(t *testing.T) {

	expectedResult := ProductRecordByProduct{
		ProdID:          1,
		ProdDescription: "prodTest",
		Records_Count:   1,
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

	expectedResult := errors.New("el id de producto ingresado no existe")

	dbTest, err := createDataBase()
	if err != nil {
		fmt.Println(err)
	}
	var repositoryAux = &repository_test{dbTest, false}
	service := NewService(repositoryAux)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	_, err1 := service.Get(ctx, 8)
	assert.Equal(t, expectedResult, err1, "Deben ser iguales")
}

func TestSaveProdRecordService(t *testing.T) {
	dbTest, err := createDataBase()
	if err != nil {
		fmt.Println(err)
	}
	prod := domain.ProductRecord{
		ID: 1,

		LastUpdateDate: "2026-12-12",
		PurchasePrice:  100,
		SalePrice:      200,
		ProductId:      1,
	}
	expectedResult := ProdRecSave{
		LastUpdateDate: "2026-12-12",
		PurchasePrice:  100,
		SalePrice:      200,
		ProductId:      1,
	}

	var repositoryAux = &repository_test{dbTest, false}
	service := NewService(repositoryAux)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	result, _ := service.Save(ctx, prod)
	assert.Equal(t, expectedResult, result, "El id debe ser igual")
	assert.True(t, repositoryAux.flag)
}

func TestSaveErrorProdRecordService(t *testing.T) {
	dbTest, err := createDataBase()
	if err != nil {
		fmt.Println(err)
	}
	prod := domain.ProductRecord{
		ID: 1,

		LastUpdateDate: "2026-12-12",
		PurchasePrice:  100,
		SalePrice:      200,
		ProductId:      696969,
	}
	expectedResult := errors.New("el id de producto no existe")

	var repositoryAux = &repository_test{dbTest, false}
	service := NewService(repositoryAux)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	result, err := service.Save(ctx, prod)
	fmt.Println(result)
	assert.Equal(t, expectedResult, err, "El id debe ser igual")

}
