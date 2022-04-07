package section

import (
	"context"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type repository_test struct {
	flag bool
}

func (r *repository_test) GetAll(ctx context.Context) ([]domain.Section, error) {
	r.flag = true

	sections := []domain.Section{
		{
			ID:                 1,
			SectionNumber:      "abc",
			CurrentTemperature: 10,
			MinimumTemperature: 5,
			CurrentCapacity:    2,
			MinimumCapacity:    1,
			MaximumCapacity:    10,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
		{
			ID:                 2,
			SectionNumber:      "xyz",
			CurrentTemperature: 7,
			MinimumTemperature: 2,
			CurrentCapacity:    3,
			MinimumCapacity:    2,
			MaximumCapacity:    20,
			WarehouseID:        2,
			ProductTypeID:      2,
		},
	}
	return sections, nil
}

func (r *repository_test) Get(ctx context.Context, id int) (domain.Section, error) {
	r.flag = true

	sections := []domain.Section{
		{
			ID:                 1,
			SectionNumber:      "abc",
			CurrentTemperature: 10,
			MinimumTemperature: 5,
			CurrentCapacity:    2,
			MinimumCapacity:    1,
			MaximumCapacity:    10,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
		{
			ID:                 2,
			SectionNumber:      "xyz",
			CurrentTemperature: 7,
			MinimumTemperature: 2,
			CurrentCapacity:    3,
			MinimumCapacity:    2,
			MaximumCapacity:    20,
			WarehouseID:        2,
			ProductTypeID:      2,
		},
	}

	if 0 < id && id < len(sections) {
		return sections[id-1], nil
	}
	return domain.Section{}, errors.New("ID does not exist")
}

func (r *repository_test) Exists(ctx context.Context, sectionNumber string) bool {
	r.flag = true
	return sectionNumber == "abc" || sectionNumber == "xyz"
}

func (r *repository_test) Save(ctx context.Context, d domain.Section) (int, error) {
	r.flag = true
	if d.SectionNumber != "" && d.CurrentCapacity >= 0 && d.MinimumCapacity >= 0 && d.MaximumCapacity >= 0 && d.WarehouseID > 0 && d.ProductTypeID > 0 {
		return d.ID, nil
	}
	return 0, errors.New("Fields invalid")
}

func (r *repository_test) Update(ctx context.Context, d domain.Section) error {
	r.flag = true
	if d.SectionNumber != "" && d.CurrentCapacity >= 0 && d.MinimumCapacity >= 0 && d.MaximumCapacity >= 0 && d.WarehouseID > 0 && d.ProductTypeID > 0 {
		if d.ID == 1 || d.ID == 2 {
			return nil
		}
		return errors.New("Not found")
	}
	return errors.New("Fields invalid")
}

func (r *repository_test) Delete(ctx context.Context, id int) error {
	r.flag = true
	if id == 1 || id == 2 {
		return nil
	}
	return errors.New("Not found")
}

func Test_create_ok(t *testing.T) {
	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	newSection := domain.Section{
		SectionNumber:      "def",
		CurrentTemperature: 10,
		MinimumTemperature: 5,
		CurrentCapacity:    2,
		MinimumCapacity:    1,
		MaximumCapacity:    10,
		WarehouseID:        1,
		ProductTypeID:      1,
	}

	_, err := service.Save(ctx, newSection)

	assert.Nil(t, err)
	assert.True(t, repository.flag)
}

func Test_create_fail(t *testing.T) {
	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	newSection := domain.Section{
		SectionNumber:      "",
		CurrentTemperature: 10,
		MinimumTemperature: 5,
		CurrentCapacity:    2,
		MinimumCapacity:    1,
		MaximumCapacity:    10,
		WarehouseID:        1,
		ProductTypeID:      1,
	}

	_, err := service.Save(ctx, newSection)

	assert.NotNil(t, err)
	assert.True(t, repository.flag)
}

func Test_create_confict(t *testing.T) {
	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	newSection := domain.Section{
		SectionNumber:      "abc",
		CurrentTemperature: 10,
		MinimumTemperature: 5,
		CurrentCapacity:    2,
		MinimumCapacity:    1,
		MaximumCapacity:    10,
		WarehouseID:        1,
		ProductTypeID:      1,
	}

	_, err := service.Save(ctx, newSection)

	assert.Equal(t, errors.New("SectionNumber already exists"), err, "Must be equal")
	assert.True(t, repository.flag)
}

func Test_find_all(t *testing.T) {
	expected := []domain.Section{
		{
			ID:                 1,
			SectionNumber:      "abc",
			CurrentTemperature: 10,
			MinimumTemperature: 5,
			CurrentCapacity:    2,
			MinimumCapacity:    1,
			MaximumCapacity:    10,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
		{
			ID:                 2,
			SectionNumber:      "xyz",
			CurrentTemperature: 7,
			MinimumTemperature: 2,
			CurrentCapacity:    3,
			MinimumCapacity:    2,
			MaximumCapacity:    20,
			WarehouseID:        2,
			ProductTypeID:      2,
		},
	}

	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	res, _ := service.GetAll(ctx)

	assert.Equal(t, expected, res, "Must be equal")
	assert.True(t, repository.flag)
}

func Test_find_by_id_non_existent(t *testing.T) {

	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	_, err := service.Get(ctx, 10)

	assert.NotNil(t, err)
	assert.True(t, repository.flag)
}

func Test_find_by_id_existent(t *testing.T) {
	expected := domain.Section{
		ID:                 1,
		SectionNumber:      "abc",
		CurrentTemperature: 10,
		MinimumTemperature: 5,
		CurrentCapacity:    2,
		MinimumCapacity:    1,
		MaximumCapacity:    10,
		WarehouseID:        1,
		ProductTypeID:      1,
	}

	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	res, _ := service.Get(ctx, 1)

	assert.Equal(t, expected, res, "Must be equal")
	assert.True(t, repository.flag)
}

func Test_update_existent(t *testing.T) {
	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	updateSection := domain.Section{
		ID:                 1,
		SectionNumber:      "abc",
		CurrentTemperature: 10,
		MinimumTemperature: 5,
		CurrentCapacity:    2,
		MinimumCapacity:    1,
		MaximumCapacity:    10,
		WarehouseID:        1,
		ProductTypeID:      1,
	}

	res := service.Update(ctx, updateSection)

	assert.Nil(t, res)
	assert.True(t, repository.flag)
}

func Test_update_non_existent(t *testing.T) {
	expected := errors.New("Not found")

	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	updateSection := domain.Section{
		ID:                 9,
		SectionNumber:      "jkl",
		CurrentTemperature: 10,
		MinimumTemperature: 5,
		CurrentCapacity:    2,
		MinimumCapacity:    1,
		MaximumCapacity:    10,
		WarehouseID:        1,
		ProductTypeID:      1,
	}

	err := service.Update(ctx, updateSection)

	assert.Equal(t, expected, err, "Must be equal")
	assert.True(t, repository.flag)
}

func Test_update_conflict(t *testing.T) {
	expected := errors.New("section number already used")

	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	updateSection := domain.Section{
		ID:                 1,
		SectionNumber:      "xyz",
		CurrentTemperature: 10,
		MinimumTemperature: 5,
		CurrentCapacity:    2,
		MinimumCapacity:    1,
		MaximumCapacity:    10,
		WarehouseID:        1,
		ProductTypeID:      1,
	}

	err := service.Update(ctx, updateSection)

	assert.Equal(t, expected, err, "Must be equal")
	assert.True(t, repository.flag)
}

func Test_delete_non_existent(t *testing.T) {
	expected := errors.New("Not found")

	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	res := service.Delete(ctx, 9)

	assert.Equal(t, expected, res, "Must be equal")
	assert.True(t, repository.flag)
}

func Test_delete_ok(t *testing.T) {
	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	res := service.Delete(ctx, 1)

	assert.Nil(t, res)
	assert.True(t, repository.flag)
}
