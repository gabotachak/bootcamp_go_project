package buyer

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
	db, errDb := db.Init("unit_test")
	assert.NoError(t, errDb)
	repository := NewRepository(db)

	//declare the expected variable
	expected := domain.Buyer{
		ID:           2,
		CardNumberID: "BuyerCardNumberId",
		FirstName:    "Name",
		LastName:     "LastName",
	}
	expectedID := 2

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method store
	buyerID, err := repository.Save(ctx, expected)

	//assert
	assert.Equal(t, expectedID, buyerID)
	assert.NotEmpty(t, buyerID)
	assert.Nil(t, err)
}

// unit test of repository store method when create is Fail because ordernumber is empty
// repository unit test
func TestCreateConflictRepository(t *testing.T) {

	//initialitation database
	db, errDb := db.Init("unit_test")
	assert.NoError(t, errDb)
	repository := NewRepository(db)

	//declare the expected variable with orderNumber empty
	data := domain.Buyer{
		ID:           3,
		CardNumberID: "",
		FirstName:    "Name",
		LastName:     "LastName",
	}
	expectedError := fmt.Errorf("el CardNumberId está vacio")

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method store
	purchaseOrderID, err := repository.Save(ctx, data)

	//assert
	assert.Empty(t, purchaseOrderID)
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

	flag := repository.Exists(ctx, "BuyerCardNumberId")

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

	flag := repository.Exists(ctx, "NoCardNumberId")

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
	// dataExpected := []domain.Buyer{
	// 	{
	// 		ID:           1,
	// 		CardNumberID: "1",
	// 		FirstName:    "Nombre",
	// 		LastName:     "Apellido",
	// 	},
	// 	{
	// 		ID:           2,
	// 		CardNumberID: "BuyerCardNumberId",
	// 		FirstName:    "Name",
	// 		LastName:     "LastName",
	// 	},
	// }

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
	// dataExpected := domain.Buyer{
	// 	ID:           2,
	// 	CardNumberID: "BuyerCardNumberId",
	// 	FirstName:    "Name",
	// 	LastName:     "LastName",
	// }

	//method GET
	dataObtained, err := repository.Get(ctx, "BuyerCardNumberId")

	//assert
	assert.NoError(t, err)
	assert.NotEmpty(t, dataObtained)
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
	dataExpected := fmt.Errorf("id del buyer no existe")

	dataObtained, err := repository.Get(ctx, "NoCardNumberId")
	assert.Empty(t, dataObtained)
	assert.Error(t, err)
	assert.Equal(t, dataExpected, err)
}

// unit test of repository Update method when Update is successful
// repository unit test
func TestUpdateOkRepository(t *testing.T) {

	//initialitation database
	db, errDb := db.Init("unit_test")
	assert.NoError(t, errDb)
	repository := NewRepository(db)

	//declare the expected variable
	data := domain.Buyer{
		ID:           1,
		CardNumberID: "1",
		FirstName:    "NameUpdated",
		LastName:     "LastNameUpdated",
	}

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method store
	err := repository.Update(ctx, data)

	//assert
	assert.Nil(t, err)
}

// unit test of repository store method when Update is Fail because cardNumberId not exist
// repository unit test
func TestUpdateConflictRepository(t *testing.T) {

	//initialitation database
	db, errDb := db.Init("unit_test")
	assert.NoError(t, errDb)
	repository := NewRepository(db)

	//declare the expected variable with orderNumber empty
	data := domain.Buyer{
		ID:           100,
		CardNumberID: "BuyerCardNumberIdUpdated",
		FirstName:    "NameUpdated",
		LastName:     "LastNameUpdated",
	}
	expectedError := fmt.Errorf("el CardNumberId no existe")

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method store
	err := repository.Update(ctx, data)

	//assert
	assert.Equal(t, expectedError, err)
}

// unit test of repository Delete method when Delete is successful
// repository unit test
func TestDeleteOkRepository(t *testing.T) {

	//initialitation database
	db, errDb := db.Init("unit_test")
	assert.NoError(t, errDb)
	repository := NewRepository(db)

	//declare the expected variable
	CardNumberIDForDelete := "BuyerCardNumberId"

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method store
	err := repository.Delete(ctx, CardNumberIDForDelete)

	//assert
	assert.Nil(t, err)
}

// unit test of repository store method when Delete is Fail because cardNumberId not exist
// repository unit test
func TestDeleteConflictRepository(t *testing.T) {

	//initialitation database
	db, errDb := db.Init("unit_test")
	assert.NoError(t, errDb)
	repository := NewRepository(db)

	//declare the expected variable with orderNumber empty
	CardNumberIDForDelete := "No"
	expectedError := fmt.Errorf("el CardNumberId no existe")

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method store
	err := repository.Delete(ctx, CardNumberIDForDelete)

	//assert
	assert.Equal(t, expectedError, err)
}

// unit test of repository GetPurchaseOrders method when GetPurchaseOrders is successful
// repository unit test
func TestGetPurchaseOrdersOkRepository(t *testing.T) {

	//initialitation database
	db, errDb := db.Init("unit_test")
	assert.NoError(t, errDb)
	repository := NewRepository(db)

	//declare the expected variable
	// expectedBuyerPurchaseOrder := []BuyerPurchaseOrders{
	// 	{
	// 		ID:                  1,
	// 		CardNumberID:        "1",
	// 		FirstName:           "NameUpdated",
	// 		LastName:            "LastNameUpdated",
	// 		PurchaseOrdersCount: 1,
	// 	},
	// }

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method store
	buyerPurchaseOrder, err := repository.GetPurchaseOrders(ctx)

	//assert
	assert.NotEmpty(t, buyerPurchaseOrder)
	assert.Nil(t, err)
}

// unit test of repository GetPurchaseOrdersByBuyer method when ByBuyer is successful
// repository unit test
func TestGetPurchaseOrdersByBuyerOkRepository(t *testing.T) {

	//initialitation database
	db, errDb := db.Init("unit_test")
	assert.NoError(t, errDb)
	repository := NewRepository(db)

	//declare the expected variable
	// expectedCount := BuyerPurchaseOrders{

	// 	ID:                  1,
	// 	CardNumberID:        "1",
	// 	FirstName:           "NameUpdated",
	// 	LastName:            "LastNameUpdated",
	// 	PurchaseOrdersCount: 1,
	// }
	buyerId := 1

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method store
	buyerPurchaseOrderCount, err := repository.GetPurchaseOrdersByBuyer(ctx, buyerId)

	//assert
	assert.NotEmpty(t, buyerPurchaseOrderCount)
	assert.Nil(t, err)
}

// unit test of repository GetPurchaseOrdersByBuyer method when fail because ByBuyer not exist
// repository unit test
func TestGetPurchaseOrdersByBuyerFailRepository(t *testing.T) {

	//initialitation database
	db, errDb := db.Init("unit_test")
	assert.NoError(t, errDb)
	repository := NewRepository(db)

	//declare the expected variable
	expectedError := fmt.Errorf("el CardNumberId ingresado no existe")
	buyerId := 100

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method store
	_, err := repository.GetPurchaseOrdersByBuyer(ctx, buyerId)

	//assert
	assert.Equal(t, err, expectedError)

}
