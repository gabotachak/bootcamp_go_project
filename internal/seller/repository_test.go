package seller

import (
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	db "github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

/*
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Seller, error)
	Get(ctx context.Context, id int) (domain.Seller, error)
	Save(ctx context.Context, s domain.Seller) (int, error)
	Update(ctx context.Context, s domain.Seller) error
	Delete(ctx context.Context, id int) error
	Exists(ctx context.Context, cid int) bool
}
*/
/**
* Unit test of save method from repository when create is successful
 */
func TestSaveSeller(t *testing.T) {
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	expected := domain.Seller{
		CID:         101,
		CompanyName: "Nombre Prueba",
		Address:     "Direccion Prueba",
		Telephone:   "1234567890",
		LocalityID:  1,
	}

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	res, err1 := repository.Save(ctx, expected)
	fmt.Println(err1)

	if err1 != nil {
		t.Error(err1)
	}

	assert.NotEmpty(t, res)
}

/**
* Unit test of save method from repository when create is failed
 */
func TestSaveFail(t *testing.T) {
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	expected := domain.Seller{
		CID:         1,
		CompanyName: "Nombre Prueba",
		Address:     "Direccion Prueba",
		Telephone:   "1234567890",
		LocalityID:  1,
	}
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	_, err1 := repository.Save(ctx, expected)

	assert.Error(t, err1)
}

/**
* Unit test of GetGeneralReport method from repository when list is successful
 */
func TestGetAllSellers(t *testing.T) {
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	res, err1 := repository.GetAll(ctx)

	if err1 != nil {
		t.Error(err1)
	}

	assert.NotEmpty(t, res)
}

func TestReport(t *testing.T) {
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	id := 1

	res, err1 := repository.Get(ctx, id)
	fmt.Println(err1)
	if err1 != nil {
		t.Error(err1)
	}

	assert.NotEmpty(t, res)
}

func TestGetFail(t *testing.T) {
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	id := 50

	_, err1 := repository.Get(ctx, id)

	assert.Equal(t, errors.New("sql: no rows in result set"), err1)
	assert.Error(t, err1)
}

/**
* Unit test of save method from repository when create is successful
 */
func TestUpdateSeller(t *testing.T) {
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	expected := domain.Seller{
		ID:          2,
		CID:         141,
		CompanyName: "Nombre Prueba",
		Address:     "Direccion Prueba",
		Telephone:   "1234567890",
		LocalityID:  1,
	}

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	err1 := repository.Update(ctx, expected)
	fmt.Println(err1)

	if err1 != nil {
		t.Error(err1)
	}

	assert.Equal(t, nil, err1, "deben ser iguales")
	assert.Nil(t, err1)
}

/**
* Unit test of save method from repository when create is successful
 */
func TestUpdateSellerFail(t *testing.T) {
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	expected := domain.Seller{
		ID:          120,
		CID:         111,
		CompanyName: "Nombre Prueba",
		Address:     "Direccion Prueba",
		Telephone:   "1234567890",
		LocalityID:  1,
	}

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	err1 := repository.Update(ctx, expected)

	assert.Equal(t, errors.New("seller not found"), err1, "deben ser iguales")
	assert.NotNil(t, err1)
}

/**
* Unit test of save method from repository when create is successful
 */
func TestDeleteSellerOk(t *testing.T) {
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	expected := 2

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	err1 := repository.Delete(ctx, expected)

	assert.Equal(t, nil, err1, "deben ser iguales")
	assert.Nil(t, err1)
}

/**
* Unit test of save method from repository when create is successful
 */
func TestDeleteSellerFail(t *testing.T) {
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	expected := 50

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	err1 := repository.Delete(ctx, expected)

	assert.Equal(t, errors.New("seller not found"), err1, "deben ser iguales")
	assert.NotNil(t, err1)
}

/**
* Unit test of save method from repository when create is successful
 */
func TestExistsOk(t *testing.T) {
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	cidExpected := 1

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	res := repository.Exists(ctx, cidExpected)

	assert.Equal(t, true, res, "deben ser iguales")
}

/**
* Unit test of save method from repository when create is successful
 */
func TestExistsFail(t *testing.T) {
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	cidExpected := 500

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	res := repository.Exists(ctx, cidExpected)

	assert.Equal(t, false, res, "deben ser iguales")
}
