package employee

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

func (r *repository_test) GetAll(ctx context.Context) ([]domain.Employee, error) {
	r.flag = true

	employees := []domain.Employee{
		{
			ID:           1,
			CardNumberID: "12345",
			FirstName:    "Jhon",
			LastName:     "Doe",
			WarehouseID:  2,
		},
		{
			ID:           2,
			CardNumberID: "123456",
			FirstName:    "Adam",
			LastName:     "Wick",
			WarehouseID:  1,
		},
	}

	return employees, nil
}

func (r *repository_test) Get(ctx context.Context, id string) (domain.Employee, error) {
	r.flag = true
	if id == "1" {
		employee := domain.Employee{
			ID:           1,
			CardNumberID: "12345",
			FirstName:    "Jhon",
			LastName:     "Doe",
			WarehouseID:  2,
		}
		return employee, nil
	} else {
		return domain.Employee{}, errors.New("el id no existe en la base de datos")
	}
}

func (r *repository_test) Exists(ctx context.Context, cardNumberID string) bool {
	r.flag = true
	return cardNumberID == "1"
}

func (r *repository_test) Save(ctx context.Context, e domain.Employee) (int, error) {
	r.flag = true

	if e.CardNumberID == "" || e.FirstName == "" || e.LastName == "" || e.WarehouseID == 0 {
		return -1, errors.New("Verificar que todos los campos esten correctos o verifique el tipo de dato")
	}

	return 1, nil
}

func (r *repository_test) Update(ctx context.Context, e domain.Employee) error {
	r.flag = true
	if e.CardNumberID == "" || e.FirstName == "" || e.LastName == "" || e.WarehouseID == 0 {
		return errors.New("Verificar que todos los campos esten correctos o verifique el tipo de dato")
	}

	return nil
}

func (r *repository_test) Delete(ctx context.Context, id string) error {
	r.flag = true
	if id != "1" {
		return errors.New("el id no existe en la base de datos")
	}
	return nil
}
func (r *repository_test) GetInboundOrders(ctx context.Context) ([]employeeInboundOrders, error) {
	r.flag = true

	return []employeeInboundOrders{}, nil
}
func (r *repository_test) GetInboundOrdersByEmployee(ctx context.Context, id int) (employeeInboundOrders, error) {
	r.flag = true
	if id == 1 {
		return employeeInboundOrders{}, errors.New("el id no existe en la base de datos")
	}
	return employeeInboundOrders{}, nil
}

// Method that checks if the service "Save" works fine when the data is fine
// Requirement create_ok
func TestSave(t *testing.T) {

	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	newEmployee := domain.Employee{
		ID:           1,
		CardNumberID: "2",
		FirstName:    "Jhon",
		LastName:     "Doe",
		WarehouseID:  2,
	}

	res, _ := service.Save(ctx, newEmployee)

	assert.Equal(t, 1, res, "deben ser iguales")
	assert.True(t, repository.flag)
}

// Method that checks if the service "Save" check if data is incorrect
// Requirement create_conflict
func TestSaveError(t *testing.T) {

	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	resultadoEsperado := errors.New("el card_number_id ya existe")

	newEmployee := domain.Employee{
		CardNumberID: "1",
		FirstName:    "Jhon",
		LastName:     "Doe",
		WarehouseID:  2,
	}

	_, err := service.Save(ctx, newEmployee)

	assert.Equal(t, resultadoEsperado, err, "deben ser iguales")
	assert.True(t, repository.flag)
}

// Method that checks if the service "GetAll" works fine when is data in the database
// Requirement find_all
func TestGetAll(t *testing.T) {
	resultadoEsperado := []domain.Employee{
		{
			ID:           1,
			CardNumberID: "12345",
			FirstName:    "Jhon",
			LastName:     "Doe",
			WarehouseID:  2,
		},
		{
			ID:           2,
			CardNumberID: "123456",
			FirstName:    "Adam",
			LastName:     "Wick",
			WarehouseID:  1,
		},
	}

	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	res, _ := service.GetAll(ctx)

	assert.Equal(t, resultadoEsperado, res, "deben ser iguales")
	assert.True(t, repository.flag)
}

// Method that checks if the service "Get" works fine when id exist
// Requirement find_by_id_existent
func TestGet(t *testing.T) {
	resultadoEsperado := domain.Employee{
		ID:           1,
		CardNumberID: "12345",
		FirstName:    "Jhon",
		LastName:     "Doe",
		WarehouseID:  2,
	}

	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	res, _ := service.Get(ctx, "1")

	assert.Equal(t, resultadoEsperado, res, "deben ser iguales")
	assert.True(t, repository.flag)
}

// Method that checks if the service "Get" works fine when id does not exist
// Requirement find_by_id_non_existent
func TestGetError(t *testing.T) {
	resultadoEsperado := errors.New("el id no existe en la base de datos")

	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	_, e := service.Get(ctx, "2")

	assert.Equal(t, resultadoEsperado, e, "deben ser iguales")
	assert.True(t, repository.flag)
}

// Method that checks if the service "Update" works fine when the data is correct
// Requirement update_existent
func TestUpdate(t *testing.T) {

	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	newEmployee := domain.Employee{
		ID:           1,
		CardNumberID: "1",
		FirstName:    "Jhon",
		LastName:     "Doe",
		WarehouseID:  2,
	}
	er := service.Update(ctx, newEmployee)

	assert.Equal(t, nil, er, "deben ser iguales")
	assert.True(t, repository.flag)
}

// Method that checks if the service "Update" works fine when the data is incorrect
// Requeriment update_non_existent
func TestUpdateError(t *testing.T) {
	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	resultadoEsperado := errors.New("el card_number_id no existe")

	newEmployee := domain.Employee{
		ID:           1,
		CardNumberID: "5",
		FirstName:    "Jhon",
		LastName:     "Doe",
	}
	er := service.Update(ctx, newEmployee)

	assert.Equal(t, resultadoEsperado, er, "deben ser iguales")
}

// Method that checks if the service "Delete" works fine when id exist
// Requirement delete_ok
func TestDelete(t *testing.T) {

	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	res := service.Delete(ctx, "1")

	assert.Equal(t, nil, res, "deben ser iguales")
	assert.True(t, repository.flag)
}

// Method that checks if the service "Delete" works fine when id does not exist
// Requirement delete_non_existent
func TestDeleteError(t *testing.T) {
	resultadoEsperado := errors.New("el id no existe en la base de datos")

	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	res := service.Delete(ctx, "123")

	assert.Equal(t, resultadoEsperado, res, "deben ser iguales")
	assert.True(t, repository.flag)
}

func TestGetInboundOrder(t *testing.T) {
	resultadoEsperado := []employeeInboundOrders{}
	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	res, err := service.GetInboundOrders(ctx)
	assert.Nil(t, err)
	assert.Equal(t, resultadoEsperado, res, "deben ser iguales")
	assert.True(t, repository.flag)
}

func TestGetInboundOrderByEmployee(t *testing.T) {
	resultadoEsperado := employeeInboundOrders{}
	var repository = &repository_test{false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	res, err := service.GetInboundOrdersByEmployee(ctx, 2)
	assert.Nil(t, err)
	assert.Equal(t, resultadoEsperado, res, "deben ser iguales")
	assert.True(t, repository.flag)
}
