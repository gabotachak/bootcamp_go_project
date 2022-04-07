package employee

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	db "github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

//unit test of repository getAll method when workly correctly
func TestRepositoryEmployeeGetAll(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)
	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//data expected
	dataExpected := []domain.Employee{
		{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Bob",
			LastName:     "Tyson",
			WarehouseID:  1,
		},
		{
			ID:           2,
			CardNumberID: "1234",
			FirstName:    "Bob2",
			LastName:     "Tyson2",
			WarehouseID:  1,
		},
	}

	dataObtained, err := repository.GetAll(ctx)
	assert.NoError(t, err)
	assert.Equal(t, dataExpected, dataObtained)
}

//unit test of repository get method when workly correctly
func TestRepositoryEmployeeGet(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)
	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//data expected
	dataExpected := domain.Employee{
		ID:           1,
		CardNumberID: "123",
		FirstName:    "Bob",
		LastName:     "Tyson",
		WarehouseID:  1,
	}
	cardNumberIDSearch := "123"
	dataObtained, err := repository.Get(ctx, cardNumberIDSearch)
	assert.NoError(t, err)
	assert.Equal(t, dataExpected, dataObtained)
}

//unit test of repository get method when card_number_id dont exist
func TestRepositoryEmployeeGetError(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)
	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//data expected
	dataExpected := domain.Employee{}
	cardNumberIDSearch := "9999999"
	dataObtained, err := repository.Get(ctx, cardNumberIDSearch)
	assert.Equal(t, errors.New("sql: no rows in result set"), err)
	assert.Equal(t, dataExpected, dataObtained)
}

//unit test of GetInboundOrders
func TestRepositoryEmployeeGetInboundOrders(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)
	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//data expected
	dataExpected := []employeeInboundOrders{
		{
			ID:                 1,
			CardNumberID:       "123",
			FirstName:          "Bob",
			LastName:           "Tyson",
			WarehouseID:        1,
			InboundOrdersCount: 0,
		},
		{
			ID:                 2,
			CardNumberID:       "1234",
			FirstName:          "Bob2",
			LastName:           "Tyson2",
			WarehouseID:        1,
			InboundOrdersCount: 0,
		},
	}
	inboundOrders, err := repository.GetInboundOrders(ctx)
	assert.NoError(t, err)
	assert.Equal(t, dataExpected, inboundOrders)
}

//unit test of GetInboundOrdersByEmployee
func TestRepositoryEmployeeGetInboundOrdersByEmployee(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)
	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//data expected
	errExpected := errors.New("sql: no rows in result set")
	dataExpected := employeeInboundOrders{}
	inboundOrders, err := repository.GetInboundOrdersByEmployee(ctx, 99)
	assert.Equal(t, errExpected, err)
	assert.Equal(t, dataExpected, inboundOrders)
}

//unit test of repository store method when workly correctly
func TestRepositoryEmployeeStore(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)
	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//data expected
	dataStorage := domain.Employee{
		CardNumberID: "123456",
		FirstName:    "Bob",
		LastName:     "Tyson",
		WarehouseID:  1,
	}
	dataExpected := 3
	dataObtained, err := repository.Save(ctx, dataStorage)
	assert.NoError(t, err)
	assert.Equal(t, dataExpected, dataObtained)
}

//unit test of repository store method when warehouse dont exist
func TestRepositoryEmployeeStoreError(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)
	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//data expected
	dataStorage := domain.Employee{
		CardNumberID: "99",
		FirstName:    "Bob",
		LastName:     "Tyson",
		WarehouseID:  99,
	}
	dataExpected := 0
	dataObtained, err := repository.Save(ctx, dataStorage)
	assert.Equal(t, errors.New("sql: no rows in result set"), err)
	assert.Equal(t, dataExpected, dataObtained)
}

//unit test of repository update method when workly correctly
func TestRepositoryEmployeeUpdate(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)
	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//data expected
	dataStorage := domain.Employee{
		ID:           3,
		CardNumberID: "123456",
		FirstName:    "Bob UPDATE",
		LastName:     "Tyson UPDATE",
		WarehouseID:  1,
	}

	errCantUpdate := repository.Update(ctx, dataStorage)
	assert.NoError(t, errCantUpdate)
}

//unit test of delete method when workly correctly
func TestRepositoryEmployeeDelete(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)
	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//data expected
	cardNumberIDDelete := "123456"

	errCantDelete := repository.Delete(ctx, cardNumberIDDelete)
	assert.NoError(t, errCantDelete)
}

//unit test of delete method when cardnumberid dont exist
func TestRepositoryEmployeeDeleteError(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)
	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//data expected
	cardNumberIDDelete := "9999999"

	errCantDelete := repository.Delete(ctx, cardNumberIDDelete)
	assert.Equal(t, errors.New("empleado no encontrado"), errCantDelete)
}

//unit test of exist method when cardnumberid exist
func TestRepositoryEmployeeExist(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)
	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//data expected
	cardNumberIDSearch := "123"

	found := repository.Exists(ctx, cardNumberIDSearch)
	assert.True(t, found)
}

//unit test of exist method when cardnumberid dont exist
func TestRepositoryEmployeeExistError(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)
	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//data expected
	cardNumberIDSearch := "99"

	found := repository.Exists(ctx, cardNumberIDSearch)
	assert.False(t, found)
}
