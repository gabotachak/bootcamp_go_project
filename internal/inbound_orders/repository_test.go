package inboundOrders

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
func TestCreateOK(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	//declare the expected variable
	dataStorage := domain.Inbound_Orders{
		OrderDate:      "20201223",
		OrderNumber:    "xxyy22",
		EmployeeID:     1,
		ProductBatchID: 1,
		WarehouseID:    1,
	}

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method store
	inboundOrder, err := repository.Save(ctx, dataStorage)
	expectedID := 1

	//assert
	assert.Equal(t, expectedID, inboundOrder.ID)
	assert.NotEmpty(t, inboundOrder.ID)
	assert.Nil(t, err)
}

//unit test of repository store method when order_number is empty
func TestCreateConflict(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	//declare the expected variable
	dataStorage := domain.Inbound_Orders{
		OrderDate:      "20201223",
		EmployeeID:     1,
		ProductBatchID: 1,
		WarehouseID:    1,
	}

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method store
	inboundOrder, err := repository.Save(ctx, dataStorage)
	expectedError := fmt.Errorf("el order_number esta vacio")

	//assert
	assert.Empty(t, inboundOrder.ID)
	assert.Equal(t, expectedError, err)
}

//unit test of repository store method when employee_id not exist
func TestCreateConflictEmployee(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	//declare the expected variable
	dataStorage := domain.Inbound_Orders{
		OrderDate:      "20201223",
		OrderNumber:    "xxyy22",
		EmployeeID:     99,
		ProductBatchID: 1,
		WarehouseID:    1,
	}

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method store
	inboundOrder, err := repository.Save(ctx, dataStorage)
	expectedError := fmt.Errorf("el empleado no existe")

	//assert
	assert.Empty(t, inboundOrder.ID)
	assert.Equal(t, expectedError, err)
}

//unit test of repository store method when warehouse_id not exist
func TestCreateConflictWarehouse(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	//declare the expected variable
	dataStorage := domain.Inbound_Orders{
		OrderDate:      "20201223",
		OrderNumber:    "xxyy22",
		EmployeeID:     1,
		ProductBatchID: 1,
		WarehouseID:    99,
	}

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method store
	inboundOrder, err := repository.Save(ctx, dataStorage)
	expectedError := fmt.Errorf("el warehouse no existe")

	//assert
	assert.Empty(t, inboundOrder.ID)
	assert.Equal(t, expectedError, err)
}

//unit test of repository store method when product_batches not exist
func TestCreateConflictProductBatches(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	//declare the expected variable
	dataStorage := domain.Inbound_Orders{
		OrderDate:      "20201223",
		OrderNumber:    "xxyy22",
		EmployeeID:     1,
		ProductBatchID: 9999,
		WarehouseID:    1,
	}

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method store
	inboundOrder, err := repository.Save(ctx, dataStorage)
	expectedError := fmt.Errorf("el product_batches no existe")

	//assert
	assert.Empty(t, inboundOrder.ID)
	assert.Equal(t, expectedError, err)
}

//unit test of repository exist method when workly correctly
func TestExist(t *testing.T) {
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
func TestExistFailed(t *testing.T) {
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
func TestGetAll(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//data expected
	dataExpected := []domain.Inbound_Orders{
		{
			ID:             1,
			OrderDate:      "2020-12-23 00:00:00",
			OrderNumber:    "xxyy22",
			EmployeeID:     1,
			ProductBatchID: 1,
			WarehouseID:    1,
		},
	}

	dataObtained, err := repository.GetAll(ctx)
	assert.NoError(t, err)
	assert.Equal(t, dataExpected, dataObtained)

}

//unit test of repository get method when workly correctly
func TestGet(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//data expected
	dataExpected := domain.Inbound_Orders{
		ID:             1,
		OrderDate:      "2020-12-23 00:00:00",
		OrderNumber:    "xxyy22",
		EmployeeID:     1,
		ProductBatchID: 1,
		WarehouseID:    1,
	}

	//method GET
	dataObtained, err := repository.Get(ctx, 1)

	//assert
	assert.NoError(t, err)
	assert.Equal(t, dataExpected, dataObtained)
}

//unit test of repository get method when failed
func TestGetFailed(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//data expected
	dataExpected := fmt.Errorf("inbound order id no encontrado")

	dataObtained, err := repository.Get(ctx, 5)
	assert.Empty(t, dataObtained)
	assert.Error(t, err)
	assert.Equal(t, dataExpected, err)
}
