package product

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	db "github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// type Product struct {
// 	ID             int     `json:"id"`
// 	Description    string  `json:"description"`
// 	ExpirationRate int     `json:"expiration_rate"`
// 	FreezingRate   int     `json:"freezing_rate"`
// 	Height         float32 `json:"height"`
// 	Length         float32 `json:"length"`
// 	Netweight      float32 `json:"net_weight"`
// 	ProductCode    string  `json:"product_code"`
// 	RecomFreezTemp float32 `json:"recommended_freezing_temperature"`
// 	Width          float32 `json:"width"`
// 	ProductTypeID  int     `json:"product_type_id"`
// 	SellerID       int     `json:"seller_id"`
// }

func TestGetAllRepository(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	//initialitation context
	//ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//data expected
	// expectedResult := []domain.Product{{
	// 	Description:    "Palmitos",
	// 	ExpirationRate: 10,
	// 	FreezingRate:   1,
	// 	Height:         15,
	// 	Length:         10,
	// 	Netweight:      0.5,
	// 	ProductCode:    "Prodr7r",
	// 	RecomFreezTemp: 3,
	// 	Width:          5,
	// 	ProductTypeID:  1,
	// 	SellerID:       1,
	// }}
	result, err := repository.GetAll()
	assert.NoError(t, err)
	assert.NotEmpty(t, result)

}

// unit test of repository store method when create is successful
func TestCreateOK(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	//declare expected variable
	dataStorage := domain.Product{
		Description:    "Palmitos",
		ExpirationRate: 10,
		FreezingRate:   1,
		Height:         15,
		Length:         10,
		Netweight:      0.5,
		ProductCode:    "Prodr7r",
		RecomFreezTemp: 3,
		Width:          5,
		ProductTypeID:  1,
		SellerID:       1,
	}
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method store
	prod, err := repository.Save(ctx, dataStorage)
	//expectedResult := 5

	//veriry if it is successful
	assert.NotEmpty(t, prod)
	assert.Nil(t, err)
}

// func TestCreateConflict(t *testing.T) {
// 	//initialitation database
// 	db, err := db.Init("unit_test")
// 	assert.NoError(t, err)
// 	repository := NewRepository(db)

// 	//declare expected variable
// 	dataStorage := domain.Product{
// 		Description:    "Palmitos",
// 		ExpirationRate: 10,
// 		FreezingRate:   1,
// 		Height:         15,
// 		Length:         10,
// 		Netweight:      0.5,
// 		ProductCode:    "xxxx",
// 		RecomFreezTemp: 3,
// 		Width:          5,
// 		ProductTypeID:  1,
// 		SellerID:       1,
// 	}
// 	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

// 	//call method store
// 	prod, err := repository.Save(ctx, dataStorage)
// 	expectedResult := errors.New("la fecha debe ser mayor a la actual")
// 	fmt.Println(prod)
// 	//veriry if it is successful
// 	assert.Equal(t, expectedResult, err)

// }

func TestExistId(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	//declare expected variable
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	//call method store
	exist := repository.Exists(ctx, "abc1")
	expectedResult := true
	//veriry if it is successful
	assert.Equal(t, expectedResult, exist)

}

func TestNotExistId(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	//declare expected variable
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	//call method store
	exist := repository.Exists(ctx, "10000")
	expectedResult := false
	//veriry if it is successful
	assert.Equal(t, expectedResult, exist)

}

func TestGetRepositoryProd(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//data expected
	expectedResult := domain.Product{
		ID:             1,
		Description:    "aaa",
		ExpirationRate: 10,
		FreezingRate:   10,
		Height:         500,
		Length:         300,
		Netweight:      400,
		ProductCode:    "abc1",
		RecomFreezTemp: 20,
		Width:          50,
		ProductTypeID:  1,
		SellerID:       1,
	}

	result, err := repository.Get(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)

}

func TestUpdateRepositoryProd(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	updateProd := domain.Product{
		ID:             3,
		Description:    "Palmitooooos",
		ExpirationRate: 10,
		FreezingRate:   1,
		Height:         15,
		Length:         10,
		Netweight:      0.5,
		ProductCode:    "iluyikuyku",
		RecomFreezTemp: 3,
		Width:          5,
		ProductTypeID:  1,
		SellerID:       1,
	}

	//declare expected variable
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	//call method store
	upProd := repository.Update(ctx, updateProd)
	//veriry if it is successful
	assert.Equal(t, nil, upProd)

}

// func TestUpdateProdNotFoundRepositoryProd(t *testing.T) {
// 	//initialitation database
// 	db, err := db.Init("unit_test")
// 	assert.NoError(t, err)
// 	repository := NewRepository(db)

// 	updateProd := domain.Product{
// 		ID:             2,
// 		Description:    "Palmitooooos",
// 		ExpirationRate: 10,
// 		FreezingRate:   1,
// 		Height:         15,
// 		Length:         10,
// 		Netweight:      0.5,
// 		RecomFreezTemp: 3,
// 		Width:          5,
// 		ProductTypeID:  1,
// 		SellerID:       1,
// 	}

// 	//declare expected variable
// 	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
// 	//call method store
// 	upProd := repository.Update(ctx, updateProd)
// 	fmt.Println("upprod", upProd)
// 	//veriry if it is successful
// 	assert.NotNil(t, upProd)

// }

func TestUpdateEmptyFieldRepositoryProd(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	updateProd := domain.Product{
		ID:             4848,
		Description:    "Palmitooooos",
		ExpirationRate: 10,
		FreezingRate:   1,
		Height:         15,
		Length:         10,
		Netweight:      0.5,
		ProductCode:    "Prodghjghj",
		RecomFreezTemp: 3,
		Width:          5,
		ProductTypeID:  1,
		SellerID:       1,
	}

	//declare expected variable
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	//call method store
	upProd := repository.Update(ctx, updateProd)
	//veriry if it is successful
	assert.Equal(t, errors.New("product not found"), upProd)

}

func TestDeleteRepositoryProd(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	//declare expected variable
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	//call method store
	deleteOk := repository.Delete(ctx, 3)
	//veriry if it is successful
	assert.Equal(t, nil, deleteOk)

}

func TestDeleteProdNotFoundRepositoryProd(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	//declare expected variable
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	//call method store
	deleteOk := repository.Delete(ctx, 5668)
	//veriry if it is successful
	assert.Equal(t, errors.New("product not found"), deleteOk)

}
