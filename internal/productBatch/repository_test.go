package productBatch

import (
	"errors"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	db "github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateOk(t *testing.T) {
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	productBatch := domain.ProductBatch{
		BatchNumber:        "xyz",
		CurrentQuantity:    1,
		CurrentTemperature: 1,
		DueDate:            time.Now(),
		InitialQuantity:    1,
		ManufacturingDate:  time.Now(),
		ManufacturingHour:  time.Now(),
		MinimumTemperature: 1,
		ProductId:          1,
		SectionId:          1,
	}

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	res, err := repository.Save(ctx, productBatch)

	expectedId := 2
	assert.Equal(t, expectedId, res, "must be equal")
	assert.NotNil(t, res)
	assert.NoError(t, err)
}

func TestCreateConflictBatchNumber(t *testing.T) {
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	productBatch := domain.ProductBatch{
		BatchNumber:        "abc",
		CurrentQuantity:    1,
		CurrentTemperature: 1,
		DueDate:            time.Now(),
		InitialQuantity:    1,
		ManufacturingDate:  time.Now(),
		ManufacturingHour:  time.Now(),
		MinimumTemperature: 1,
		ProductId:          1,
		SectionId:          1,
	}

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	res, err := repository.Save(ctx, productBatch)

	expErr := errors.New("product batch already exists")
	assert.Equal(t, -1, res, "must be equal")
	assert.Equal(t, expErr, err, "must be equal")
}

func TestCreateSectionIdNotExist(t *testing.T) {
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	productBatch := domain.ProductBatch{
		BatchNumber:        "jkl",
		CurrentQuantity:    1,
		CurrentTemperature: 1,
		DueDate:            time.Now(),
		InitialQuantity:    1,
		ManufacturingDate:  time.Now(),
		ManufacturingHour:  time.Now(),
		MinimumTemperature: 1,
		ProductId:          1,
		SectionId:          9,
	}

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	res, err := repository.Save(ctx, productBatch)

	expErr := errors.New("section id does not exist")
	assert.Equal(t, -1, res, "must be equal")
	assert.Equal(t, expErr, err, "must be equal")
}

func TestCreateProductIdNotExist(t *testing.T) {
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	productBatch := domain.ProductBatch{
		BatchNumber:        "jkl",
		CurrentQuantity:    1,
		CurrentTemperature: 1,
		DueDate:            time.Now(),
		InitialQuantity:    1,
		ManufacturingDate:  time.Now(),
		ManufacturingHour:  time.Now(),
		MinimumTemperature: 1,
		ProductId:          9,
		SectionId:          1,
	}

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	res, err := repository.Save(ctx, productBatch)

	expErr := errors.New("product id does not exist")
	assert.Equal(t, -1, res, "must be equal")
	assert.Equal(t, expErr, err, "must be equal")
}

func TestReportAll(t *testing.T) {
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	// reports := []domain.ReportProductBatch{
	// 	{
	// 		SectionId:     1,
	// 		SectionNumber: "1",
	// 		Quantity:      2,
	// 	},
	// }

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	res, err := repository.ReportAll(ctx)

	assert.NotEmpty(t, res)
	assert.NotNil(t, res)
	assert.NoError(t, err)
}

func TestReportBySection(t *testing.T) {
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	reports := domain.ReportProductBatch{
		SectionId:     1,
		SectionNumber: "1",
		Quantity:      2,
	}

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	res, err := repository.ReportBySection(ctx, 1)

	assert.Equal(t, reports, res, "must be equal")
	assert.NotNil(t, res)
	assert.NoError(t, err)
}
