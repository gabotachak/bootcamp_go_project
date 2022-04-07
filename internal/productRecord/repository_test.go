package productRecord

import (
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	db "github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// ID              int     `json:"id"`
// 	LastUpdateDate string  `json:"last_update_date"`
// 	PurchasePrice  float32 `json:"purchase_price"`
// 	SalePrice      float32 `json:"sale_price"`
// 	ProductId      int     `json:"product_id"`

// unit test of repository store method when create is successful
func TestCreateOK(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	//declare expected variable
	dataStorage := domain.ProductRecord{
		LastUpdateDate: "2023-07-19",
		PurchasePrice:  10,
		SalePrice:      15,
		ProductId:      1,
	}
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method store
	productRecord, err := repository.Save(ctx, dataStorage)
	// expectedResult := ProdRecSave{
	// 	LastUpdateDate: "2023-07-19",
	// 	PurchasePrice:  10,
	// 	SalePrice:      15,
	// 	ProductId:      1,
	// }

	// //veriry if it is successful
	// assert.Equal(t, expectedResult, productRecord)
	assert.NotEmpty(t, productRecord)
	assert.Nil(t, err)
}

func TestCreateConflict(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	//declare expected variable
	dataStorage := domain.ProductRecord{
		LastUpdateDate: "2019-07-19",
		PurchasePrice:  10,
		SalePrice:      15,
		ProductId:      1,
	}
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method store
	productRecord, err := repository.Save(ctx, dataStorage)
	expectedResult := errors.New("la fecha debe ser mayor a la actual")
	fmt.Println(productRecord)
	//veriry if it is successful
	assert.Equal(t, expectedResult, err)

}

func TestExistId(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	//declare expected variable
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	//call method store
	exist := repository.ExistId(ctx, 10000)
	expectedResult := false
	//veriry if it is successful
	assert.Equal(t, expectedResult, exist)

}

func TestGetAll(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	result, err := repository.GetAll(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)

}

func TestGet(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	//data expected
	// expectedResult := ProductRecordByProduct{
	// 	ProdID:          1,
	// 	ProdDescription: "aaa",
	// 	Records_Count:   1,
	// }

	result, err := repository.Get(ctx, 1)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)

}
