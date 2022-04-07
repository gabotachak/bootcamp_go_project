package warehouse

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
func TestCreateWarehouseOk(t *testing.T) {

	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	//declare the expected variable
	expected := domain.Warehouse{
		WarehouseCode: "12abc",
		Address:       "Belgrano 1931",
		Telephone:     "358 322432",
		LocalityId:    1,
	}

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method Save
	warehouseId, err := repository.Save(ctx, expected)
	if err != nil {
		t.Error(err)
	}

	//assert
	assert.NotEmpty(t, warehouseId)
}

// unit test of repository store method when create is not successful
// WarehouseCode existent
// repository unit test
func TestCreateWarehouseFail(t *testing.T) {

	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	//declare the expected variable
	expected := domain.Warehouse{
		WarehouseCode: "12abc",
		Address:       "Belgrano 1931",
		Telephone:     "358 322432",
		LocalityId:    1,
	}

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method Save
	warehouseId, _ := repository.Save(ctx, expected)
	//assert
	assert.Equal(t, 0, warehouseId)
}

// unit test of repository store method when create is not successful
// incomplete body
// repository unit test
func TestCreateWarehouseWithoutWarehouseCodeFail(t *testing.T) {

	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	//declare the expected variable
	expected := domain.Warehouse{
		Address:   "Belgrano 1931",
		Telephone: "358 322432",
	}

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method Save
	warehouse, _ := repository.Save(ctx, expected)

	//assert
	assert.Equal(t, 0, warehouse)
}

// unit test of repository cid_exists method when warehouse_code is exist
// repository unit test
func TestExistWarehouseCodeOk(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	//declare the expected variable
	expected := true

	// declare warehouse_code
	warehouse_code := "12abc"
	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method Exists
	res := repository.Exists(ctx, warehouse_code)

	//assert
	assert.Equal(t, expected, res)
}

// unit test of repository cid_exists method when cid is not exist
// repository unit test
func TestExistWarehouseCodeFail(t *testing.T) {

	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	// declare warehouse_code
	warehouseCode := "23pi2sfsdfdaf325345hjtya22"
	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method Exists
	res := repository.Exists(ctx, warehouseCode)

	//assert
	assert.False(t, res)
}

// unit test of repository get_all_report method when data is exist
// repository unit test
func TestGetAllReportOk(t *testing.T) {

	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method get_all
	warehouses, err := repository.GetAll(ctx)
	if err != nil {
		t.Error(err)
	}

	//assert
	assert.NotEmpty(t, warehouses)
}

// unit test of repository get method when data is exist
// repository unit test
func TestGetOneOk(t *testing.T) {

	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	// declare id
	id := 1

	//expected
	expected := 1
	expected1 := "Lavalle 467"

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method get_report
	warehouse, err := repository.Get(ctx, id)
	if err != nil {
		t.Error(err)
	}

	//assert
	assert.Equal(t, expected, warehouse.LocalityId)
	assert.Equal(t, expected1, warehouse.Address)
	assert.NotNil(t, warehouse.ID)
}

// unit test of repository get method when data is not exist
// repository unit test
func TestGetOneFail(t *testing.T) {

	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	// declare id
	id := 100

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method get_report
	_, er := repository.Get(ctx, id)

	//assert
	assert.NotNil(t, er)
}

// unit test of repository update method when data is successful
// repository unit test
func TestUpdateWarehouseOk(t *testing.T) {

	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	//declare the expected variable
	param := domain.Warehouse{
		ID:            1,
		WarehouseCode: "12abcd",
		Address:       "Belgrano 1931",
		Telephone:     "358 322432",
		LocalityId:    1,
	}

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method Update
	erro := repository.Update(ctx, param)
	if err != nil {
		t.Error(err)
	}

	//assert
	assert.Nil(t, erro)
}

// unit test of repository update method when data is not successful
// warehouse_code existent
// repository unit test
func TestUpdateWarehouseFail(t *testing.T) {

	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	//declare the expected variable
	param := domain.Warehouse{
		ID:            2,
		WarehouseCode: "12abc",
		Address:       "Belgrano 1931",
		Telephone:     "358 322432",
		LocalityId:    1,
	}

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method Update
	erro := repository.Update(ctx, param)
	//assert
	assert.NotNil(t, erro)
}

// unit test of repository delete method when id exist
// id existent
// repository unit test
func TestDeleteWarehouseOk(t *testing.T) {

	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	//declare the existent id
	id := 2

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method Update
	erro := repository.Delete(ctx, id)
	//assert
	assert.Nil(t, erro)
}

// unit test of repository delete method when id not exist
// id not existent
// repository unit test
func TestDeleteWarehouseFail(t *testing.T) {

	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	//declare the inexistent id
	id := 100

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method Update
	erro := repository.Delete(ctx, id)
	//assert
	assert.NotNil(t, erro)
}
