package inboundOrders

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

func (r *repository_test) GetAll(ctx context.Context) ([]domain.Inbound_Orders, error) {
	r.flag = true
	return []domain.Inbound_Orders{}, nil
}

func (r *repository_test) Get(ctx context.Context, id int) (domain.Inbound_Orders, error) {

	r.flag = true
	if id == 0 {
		return domain.Inbound_Orders{}, errors.New("el id no existe en la base de datos")
	}
	return domain.Inbound_Orders{}, nil
}

func (r *repository_test) Exists(ctx context.Context, id int) bool {
	r.flag = true
	if id == 0 {
		return false
	}
	return true
}

func (r *repository_test) Save(ctx context.Context, io domain.Inbound_Orders) (domain.Inbound_Orders, error) {
	r.flag = true
	//validation if order_number not empty
	if io.OrderNumber == "" {
		return domain.Inbound_Orders{}, errors.New("el order_number esta vacio")
	}

	//validation if employee exist
	if io.EmployeeID == 0 {
		return domain.Inbound_Orders{}, errors.New("el empleado no existe")
	}

	//validation if warehouse exist
	if io.WarehouseID == 0 {
		return domain.Inbound_Orders{}, errors.New("el warehouse no existe")
	}

	//validation if product_batch exist
	if io.ProductBatchID == 0 {
		return domain.Inbound_Orders{}, errors.New("el product_batches no existe")
	}

	io.ID = 1
	return io, nil
}

// Method that checks if the service "GetAll" works fine when the data is fine
func TestServiceInboundOrdersGetAll(t *testing.T) {
	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	dataExpected := []domain.Inbound_Orders{}

	dataObtained, err := service.GetAll(ctx)

	assert.NoError(t, err)
	assert.Equal(t, dataExpected, dataObtained)
	assert.True(t, repository.flag)
}

// Method that checks if the service "Get" works fine when id exist
func TestServiceInboundOrdersGet(t *testing.T) {
	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	dataExpected := domain.Inbound_Orders{}

	dataObtained, err := service.Get(ctx, 1)

	assert.NoError(t, err)
	assert.Equal(t, dataExpected, dataObtained)
	assert.True(t, repository.flag)
}

// Method that checks if the service "Get" works fine when id dont exist
func TestServiceInboundOrdersGetError(t *testing.T) {
	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	errorExpected := errors.New("el id no existe en la base de datos")

	dataObtained, err := service.Get(ctx, 0)

	assert.Equal(t, errorExpected, err)
	assert.Empty(t, dataObtained)
	assert.True(t, repository.flag)
}

// Method that checks if the service "Exists" works fine when id exist
func TestServiceInboundOrdersExists(t *testing.T) {
	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	found := service.Exists(ctx, 1)

	assert.True(t, found)
	assert.True(t, repository.flag)
}

// Method that checks if the service "Exists" works fine when id dont exist
func TestServiceInboundOrdersExistsError(t *testing.T) {
	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	found := service.Exists(ctx, 0)

	assert.False(t, found)
	assert.True(t, repository.flag)
}

// Method that checks if the service "Save" works fine
func TestServiceInboundOrdersSave(t *testing.T) {
	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	dataExpected := domain.Inbound_Orders{
		ID:             1,
		OrderDate:      "aaa",
		OrderNumber:    "xxyy22",
		EmployeeID:     1,
		ProductBatchID: 1,
		WarehouseID:    1,
	}

	dataStorage := domain.Inbound_Orders{
		OrderDate:      "aaa",
		OrderNumber:    "xxyy22",
		EmployeeID:     1,
		ProductBatchID: 1,
		WarehouseID:    1,
	}

	dataObtained, err := service.Save(ctx, dataStorage)

	assert.NoError(t, err)
	assert.Equal(t, dataExpected, dataObtained)
	assert.True(t, repository.flag)
}
