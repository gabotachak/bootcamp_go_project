package purchaseOrders

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	db "github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// unit test of repository store method when create is successful
// repository unit test
func TestCreateOkRepository(t *testing.T) {

	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	//declare the expected variable
	expected := domain.PurchaseOrder{
		OrderNumber:   "order#2",
		OrderDate:     "2021-04-04",
		TrackingCode:  "trackingCode2",
		BuyerId:       1,
		OrderStatusId: 1,
		CarrierId:     1,
		WarehouseId:   1,
	}
	expectedID := "order#2"

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method store
	purchaseOrder, err := repository.Store(ctx, expected)

	//assert
	assert.Equal(t, expectedID, purchaseOrder.OrderNumber)
	assert.NotEmpty(t, purchaseOrder.OrderNumber)
	assert.Nil(t, err)
}

// unit test of repository store method when create is Fail because ordernumber is empty
// repository unit test
func TestCreateConflictRepository(t *testing.T) {

	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	//declare the expected variable with orderNumber empty
	data := domain.PurchaseOrder{

		OrderDate:     "2021-04-04",
		TrackingCode:  "trackingCode",
		BuyerId:       1,
		OrderStatusId: 1,
		CarrierId:     1,
		WarehouseId:   1,
	}
	expectedError := fmt.Errorf("el número de la orden está vacio")

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method store
	purchaseOrder, err := repository.Store(ctx, data)

	//assert
	assert.Empty(t, purchaseOrder.ID)
	assert.Equal(t, expectedError, err)
}

// unit test of repository store method when buyer not exist
// repository unit test
func TestCreateConflictBuyerRepository(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	//declare the expected variable
	dataStorage := domain.PurchaseOrder{
		OrderNumber:   "order#3",
		OrderDate:     "2021-04-04",
		TrackingCode:  "trackingCode",
		BuyerId:       100,
		OrderStatusId: 1,
		CarrierId:     1,
		WarehouseId:   1,
	}

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method store
	purchaseOrder, err := repository.Store(ctx, dataStorage)
	expectedError := fmt.Errorf("el comprador no existe")

	//assert
	assert.Empty(t, purchaseOrder.ID)
	assert.Equal(t, expectedError, err)
}

// unit test of repository store method when purchaseOrder exist
// repository unit test
func TestExistRepository(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	flag := repository.Exists(ctx, 1)

	resultExpected := true

	assert.Equal(t, resultExpected, flag)
}

//unit test of repository exist method when failed
//repository unit test
func TestExistFailedRepository(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	flag := repository.Exists(ctx, 999)

	resultExpected := false

	assert.Equal(t, resultExpected, flag)
}

//unit test of repository getAll method when workly correctly
// repository unit test
func TestGetAllRepository(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	//data expected
	dataObtained, err := repository.GetAll(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, dataObtained)

}

//unit test of repository get method when workly correctly
// repository unit test
func TestGetRepository(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//data expected
	dataExpected := domain.PurchaseOrder{
		ID:            1,
		OrderNumber:   "order#1",
		OrderDate:     "2021-04-04 00:00:00",
		TrackingCode:  "trackingCode",
		BuyerId:       1,
		OrderStatusId: 1,
		CarrierId:     1,
		WarehouseId:   1,
	}

	//method GET
	dataObtained, err := repository.Get(ctx, 1)

	//assert
	assert.NoError(t, err)
	assert.Equal(t, dataExpected, dataObtained)
}

//unit test of repository get method when failed
// repository unit test
func TestGetFailedRepository(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//data expected
	dataExpected := fmt.Errorf("id de la orden de compra no fue encontrado")

	dataObtained, err := repository.Get(ctx, 999)
	assert.Empty(t, dataObtained)
	assert.Error(t, err)
	assert.Equal(t, dataExpected, err)
}
