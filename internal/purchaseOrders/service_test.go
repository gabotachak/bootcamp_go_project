package purchaseOrders

import (
	"context"
	"database/sql"
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

type repository_test struct {
	db   *sql.DB
	flag bool
}

func (r *repository_test) GetAll(ctx context.Context) ([]domain.PurchaseOrder, error) {
	r.flag = true
	purchaseOrder := []domain.PurchaseOrder{
		{
			ID:            1,
			OrderNumber:   "order#1",
			OrderDate:     "2021-04-04 00:00:00",
			TrackingCode:  "trackingCode1",
			BuyerId:       1,
			OrderStatusId: 1,
			CarrierId:     1,
			WarehouseId:   1,
		},

		{
			ID:            2,
			OrderNumber:   "order#2",
			OrderDate:     "2021-04-04 00:00:00",
			TrackingCode:  "trackingCode2",
			BuyerId:       1,
			OrderStatusId: 1,
			CarrierId:     1,
			WarehouseId:   1,
		},
	}

	return purchaseOrder, nil
}

func (r *repository_test) Get(ctx context.Context, purchaseOrderId int) (domain.PurchaseOrder, error) {
	r.flag = true

	if purchaseOrderId == 1 {
		return domain.PurchaseOrder{
			ID:            1,
			OrderNumber:   "order#1",
			OrderDate:     "2021-04-04 00:00:00",
			TrackingCode:  "trackingCode1",
			BuyerId:       1,
			OrderStatusId: 1,
			CarrierId:     1,
			WarehouseId:   1,
		}, nil
	}

	return domain.PurchaseOrder{}, errors.New("el id no existe en la base de datos")

}

func (r *repository_test) Exists(ctx context.Context, purchaseOrderId int) bool {
	r.flag = true
	return purchaseOrderId == 1
}

func (r *repository_test) Store(ctx context.Context, p domain.PurchaseOrder) (domain.PurchaseOrder, error) {
	r.flag = true

	if p.OrderNumber == "" || p.TrackingCode == "" {
		return domain.PurchaseOrder{OrderNumber: "order#NO"}, errors.New("falta enviar un campo o algun tipo de dato no es el correcto")
	}

	if p.OrderNumber == "order#NO" {
		return domain.PurchaseOrder{OrderNumber: "order#NO"}, errors.New("ya existe una orden de compra")
	}

	return p, nil
}

type storeMock struct {
	db_test *sql.DB
}

func (sm *storeMock) openBdTest() {
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println("error inicializacion base de datos", err)
	}
	sm.db_test = db
}

func TestStoreOkService(t *testing.T) {

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	store := storeMock{}
	store.openBdTest()

	var repository = &repository_test{store.db_test, false}
	service := NewService(repository)

	newOrder := domain.PurchaseOrder{
		ID:            2,
		OrderNumber:   "order#2",
		OrderDate:     "2021-04-04 00:00:00",
		TrackingCode:  "trackingCode2",
		BuyerId:       1,
		OrderStatusId: 1,
		CarrierId:     1,
		WarehouseId:   1,
	}

	res, err := service.Store(c, newOrder)

	assert.Equal(t, newOrder, res, "deben ser iguales")
	assert.True(t, repository.flag)
	assert.Nil(t, err)
}

func TestStoreExistentService(t *testing.T) {
	errorEsperado := errors.New("la orden de compra ya existe")

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	store := storeMock{}
	store.openBdTest()

	var repository = &repository_test{store.db_test, false}
	service := NewService(repository)

	newOrder := domain.PurchaseOrder{
		ID:            1,
		OrderNumber:   "order#NO",
		OrderDate:     "2021-04-04 00:00:00",
		TrackingCode:  "trackingCode1",
		BuyerId:       1,
		OrderStatusId: 1,
		CarrierId:     1,
		WarehouseId:   1,
	}

	res, err := service.Store(c, newOrder)

	assert.Equal(t, errorEsperado, err, "deben ser iguales")
	assert.Equal(t, domain.PurchaseOrder{}, res, "deben ser iguales")
	assert.True(t, repository.flag)
}

func TestGetAllService(t *testing.T) {

	purchaseOrderExpected := []domain.PurchaseOrder{
		{
			ID:            1,
			OrderNumber:   "order#1",
			OrderDate:     "2021-04-04 00:00:00",
			TrackingCode:  "trackingCode1",
			BuyerId:       1,
			OrderStatusId: 1,
			CarrierId:     1,
			WarehouseId:   1,
		},

		{
			ID:            2,
			OrderNumber:   "order#2",
			OrderDate:     "2021-04-04 00:00:00",
			TrackingCode:  "trackingCode2",
			BuyerId:       1,
			OrderStatusId: 1,
			CarrierId:     1,
			WarehouseId:   1,
		},
	}

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	store := storeMock{}
	store.openBdTest()

	var repository = &repository_test{store.db_test, false}
	service := NewService(repository)

	res, err := service.GetAll(c)

	assert.Equal(t, purchaseOrderExpected, res, "deben ser iguales")
	assert.True(t, repository.flag)
	assert.Nil(t, err)

}

func TestGetByOrderNonExistentService(t *testing.T) {

	errorEsperado := errors.New("el id no existe en la base de datos")
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	store := storeMock{}

	var repository = &repository_test{store.db_test, false}
	service := NewService(repository)

	_, err := service.Get(c, -1)

	assert.Equal(t, errorEsperado, err, "deben ser iguales")
	assert.True(t, repository.flag)

}

func TestGetByIdExistentService(t *testing.T) {

	purchaseOrderExpected := domain.PurchaseOrder{

		ID:            1,
		OrderNumber:   "order#1",
		OrderDate:     "2021-04-04 00:00:00",
		TrackingCode:  "trackingCode1",
		BuyerId:       1,
		OrderStatusId: 1,
		CarrierId:     1,
		WarehouseId:   1,
	}

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	store := storeMock{}
	store.openBdTest()

	var repository = &repository_test{store.db_test, false}
	service := NewService(repository)

	res, err := service.Get(c, 1)

	assert.Equal(t, purchaseOrderExpected, res, "deben ser iguales")
	assert.True(t, repository.flag)
	assert.Nil(t, err)
}

func TestExistService(t *testing.T) {

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	store := storeMock{}
	store.openBdTest()

	var repository = &repository_test{store.db_test, false}
	service := NewService(repository)

	res := service.Exists(c, 1)
	assert.True(t, res)
	assert.True(t, repository.flag)
}
