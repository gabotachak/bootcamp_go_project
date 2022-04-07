package section

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	db "github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	// sections := []domain.Section{
	// 	{
	// 		ID:                 1,
	// 		SectionNumber:      "1",
	// 		CurrentTemperature: 11,
	// 		MinimumTemperature: 11,
	// 		CurrentCapacity:    11,
	// 		MinimumCapacity:    11,
	// 		MaximumCapacity:    11,
	// 		WarehouseID:        1,
	// 		ProductTypeID:      1,
	// 	},
	// }

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	res, err := repository.GetAll(ctx)

	assert.NotEmpty(t, res)
	assert.NoError(t, err)
}

func TestGetByIDExistent(t *testing.T) {
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	section := domain.Section{
		ID:                 1,
		SectionNumber:      "1",
		CurrentTemperature: 11,
		MinimumTemperature: 11,
		CurrentCapacity:    11,
		MinimumCapacity:    11,
		MaximumCapacity:    11,
		WarehouseID:        1,
		ProductTypeID:      1,
	}

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	res, err := repository.Get(ctx, section.ID)

	assert.Equal(t, section, res, "must be equal")
	assert.NoError(t, err)
}

func TestGetByIDNotExistent(t *testing.T) {
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	section := domain.Section{
		ID:                 9,
		SectionNumber:      "99",
		CurrentTemperature: 99,
		MinimumTemperature: 99,
		CurrentCapacity:    99,
		MinimumCapacity:    99,
		MaximumCapacity:    99,
		WarehouseID:        9,
		ProductTypeID:      9,
	}

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	res, err := repository.Get(ctx, section.ID)

	assert.Equal(t, domain.Section{}, res, "must be equal")
	assert.Error(t, err)
}

func TestExistsExistent(t *testing.T) {
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	res := repository.Exists(ctx, "1")

	assert.True(t, res)
}

func TestExistsNotExistent(t *testing.T) {
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	res := repository.Exists(ctx, "9")

	assert.False(t, res)
}

func TestCreateOk(t *testing.T) {
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	section := domain.Section{
		SectionNumber:      "558",
		CurrentTemperature: 22,
		MinimumTemperature: 22,
		CurrentCapacity:    22,
		MinimumCapacity:    22,
		MaximumCapacity:    22,
		WarehouseID:        1,
		ProductTypeID:      1,
	}

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	res, err := repository.Save(ctx, section)
	assert.NotEmpty(t, res)
	assert.NotNil(t, res)
	assert.NoError(t, err)
}

func TestUpdateNotExistent(t *testing.T) {
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	section := domain.Section{
		ID:                 9,
		SectionNumber:      "99",
		CurrentTemperature: 99,
		MinimumTemperature: 99,
		CurrentCapacity:    99,
		MinimumCapacity:    99,
		MaximumCapacity:    99,
		WarehouseID:        9,
		ProductTypeID:      9,
	}

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	err = repository.Update(ctx, section)

	assert.Equal(t, errors.New("section not found"), err, "must be equal")
}

func TestUpdateOk(t *testing.T) {
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	section := domain.Section{
		ID:                 3,
		SectionNumber:      "22",
		CurrentTemperature: 222,
		MinimumTemperature: 222,
		CurrentCapacity:    222,
		MinimumCapacity:    222,
		MaximumCapacity:    222,
		WarehouseID:        1,
		ProductTypeID:      1,
	}

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	err = repository.Update(ctx, section)

	assert.NoError(t, err)
}

func TestDeleteNotExistent(t *testing.T) {
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	err = repository.Delete(ctx, 9)

	assert.Equal(t, errors.New("section not found"), err, "must be equal")
}

func TestDeleteExistent(t *testing.T) {
	db, err := db.Init("unit_test")
	assert.NoError(t, err)
	repository := NewRepository(db)

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	err = repository.Delete(ctx, 3)

	assert.NoError(t, err)
}
