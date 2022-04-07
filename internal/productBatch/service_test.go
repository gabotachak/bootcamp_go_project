package productBatch

import (
	"context"
	"errors"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type repository_test struct {
	flag bool
}

func (r *repository_test) Save(ctx context.Context, pb domain.ProductBatch) (int, error) {
	r.flag = true
	if r.BatchExists(ctx, pb.BatchNumber) {
		return -1, errors.New("product batch already exists")
	}
	if !r.SectionExists(ctx, pb.SectionId) {
		return -1, errors.New("section id does not exist")
	}
	if !r.ProductExists(ctx, pb.ProductId) {
		return -1, errors.New("product id does not exist")
	}
	return 2, nil
}

func (r *repository_test) ReportAll(ctx context.Context) ([]domain.ReportProductBatch, error) {
	r.flag = true

	reports := []domain.ReportProductBatch{
		{
			SectionId:     1,
			SectionNumber: "1",
			Quantity:      1,
		},
	}

	return reports, nil
}

func (r *repository_test) ReportBySection(ctx context.Context, sectionId int) (domain.ReportProductBatch, error) {
	r.flag = true

	reports := []domain.ReportProductBatch{
		{
			SectionId:     1,
			SectionNumber: "1",
			Quantity:      1,
		},
	}

	if sectionId-1 < 0 || sectionId-1 >= len(reports) {
		return domain.ReportProductBatch{}, errors.New("not found section")
	}

	return reports[sectionId-1], nil
}

func (r *repository_test) BatchExists(ctx context.Context, batchNumber string) bool {
	return batchNumber == "1"
}

func (r *repository_test) SectionExists(ctx context.Context, sectionId int) bool {
	return sectionId == 1
}

func (r *repository_test) ProductExists(ctx context.Context, productId int) bool {
	return productId == 1
}

func Test_create(t *testing.T) {
	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	newProductBatch := domain.ProductBatch{
		BatchNumber:        "2",
		CurrentQuantity:    22,
		CurrentTemperature: 22,
		DueDate:            time.Now(),
		InitialQuantity:    22,
		ManufacturingDate:  time.Now(),
		ManufacturingHour:  time.Now(),
		MinimumTemperature: 22,
		ProductId:          1,
		SectionId:          1,
	}
	_, err := service.Save(ctx, newProductBatch)
	assert.NoError(t, err)
	assert.True(t, repository.flag)
}

func Test_report_all(t *testing.T) {
	expected := []domain.ReportProductBatch{
		{
			SectionId:     1,
			SectionNumber: "1",
			Quantity:      1,
		},
	}

	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	res, _ := service.ReportAll(ctx)

	assert.Equal(t, expected, res, "Must be equal")
	assert.True(t, repository.flag)
}

func Test_report_by_section_non_existent(t *testing.T) {

	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	_, err := service.ReportBySection(ctx, 10)

	assert.NotNil(t, err)
	assert.False(t, repository.flag)
}

func Test_report_by_section_existent(t *testing.T) {
	expected := domain.ReportProductBatch{
		SectionId:     1,
		SectionNumber: "1",
		Quantity:      1,
	}

	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	res, _ := service.ReportBySection(ctx, 1)

	assert.Equal(t, expected, res, "Must be equal")
	assert.True(t, repository.flag)
}
