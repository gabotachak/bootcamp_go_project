package locality

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

/**
* Unit test of save method from repository when create is successful
 */
func TestSave(t *testing.T) {
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	expected := domain.Locality{
		LocalityName: "Prueba",
		ProvinceName: "Prueba",
		CountryName:  "Prueba",
	}

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	res, err1 := repository.Save(ctx, expected)

	if err1 != nil {
		t.Error(err1)
	}

	assert.Equal(t, expected.LocalityName, res.LocalityName, "deben ser iguales")
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

	expected := domain.Locality{
		LocalityName: "Prueba",
		ProvinceName: "Prueba1",
		CountryName:  "Prueba",
	}
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	_, err1 := repository.Save(ctx, expected)

	assert.Error(t, err1)
}

/**
* Unit test of GetGeneralReport method from repository when list is successful
 */
func TestGeneralReport(t *testing.T) {
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	res, err1 := repository.GetGeneralReport(ctx)

	if err1 != nil {
		t.Error(err1)
	}

	assert.NotEmpty(t, res)
}

/**
* Unit test of GetReport method from repository when list is successful
 */
func TestReport(t *testing.T) {
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	id := 1

	res, err1 := repository.GetReport(ctx, id)
	if err1 != nil {
		t.Error(err1)
	}

	assert.Equal(t, id, res.LocalityId)
}

/**
* Unit test of GetReportFail method from repository when list is failed
 */
func TestReportFail(t *testing.T) {
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	id := 3

	_, err1 := repository.GetReport(ctx, id)

	assert.Equal(t, errors.New("sql: no rows in result set"), err1)
	assert.Error(t, err1)
}
