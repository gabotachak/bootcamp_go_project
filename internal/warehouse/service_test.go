package warehouse

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

type repository_test struct {
	db   *sql.DB
	flag bool
}

func (r *repository_test) GetAll(ctx context.Context) ([]domain.Warehouse, error) {
	r.flag = true

	warehouses := []domain.Warehouse{{
		ID:            1,
		Address:       "Lavalle 467",
		Telephone:     "358 4329528",
		WarehouseCode: "IT123",
		LocalityId:    1,
	}, {
		ID:            2,
		Address:       "Colombia 1467",
		Telephone:     "358 4337528",
		WarehouseCode: "OP456",
		LocalityId:    5,
	}}
	return warehouses, nil
}

func (r *repository_test) Get(ctx context.Context, id int) (domain.Warehouse, error) {
	r.flag = true
	//find_by_id_existent
	if id == 1 {
		w := domain.Warehouse{
			ID:            1,
			Address:       "Lavalle 467",
			Telephone:     "358 4329528",
			WarehouseCode: "123WC",
			LocalityId:    1,
		}
		return w, nil
	} else {
		if id == 2 {
			w1 := domain.Warehouse{
				ID:            2,
				Address:       "Corrientes 63",
				Telephone:     "358 4329528",
				WarehouseCode: "111WC",
			}
			return w1, nil
		}
		// find_by_id_non_existent
		return domain.Warehouse{}, errors.New("el id no existe en la base de datos")
	}
}

func (r *repository_test) Exists(ctx context.Context, warehouseCode string) bool {
	//i set the flag to true in case the warehousecode exists in the database
	r.flag = true
	// create_conflict in case the result is true
	return warehouseCode == "123WC"
}

func (r *repository_test) Save(ctx context.Context, w domain.Warehouse) (int, error) {
	r.flag = true

	if w.Address == "" || w.Telephone == "" || w.WarehouseCode == "" {
		return -1, errors.New("falta enviar un campo o algun tipo de dato no es el correcto")
	}

	// create_ok
	return 1, nil
}

func (r *repository_test) Update(ctx context.Context, w domain.Warehouse) error {
	r.flag = true
	if w.Address == "" || w.Telephone == "" || w.WarehouseCode == "" {
		return errors.New("falta enviar un campo o algun tipo de dato no es el correcto")
	}
	// update_existent
	return nil
}

func (r *repository_test) Delete(ctx context.Context, id int) error {
	r.flag = true
	//delete_non_existent
	if id != 1 {
		return errors.New("el id es inexistente")
	}
	// delete_ok
	return nil
}

//Method that initializes the database
func createDataBase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./meli_test.db")
	if err != nil {
		return nil, errors.New("error inicializacion base de datos")
	}
	return db, nil
}

//Method that checks if the service get works fine when id exist
func TestGetOk(t *testing.T) {
	expected_res := domain.Warehouse{
		ID:            1,
		Address:       "Lavalle 467",
		Telephone:     "358 4329528",
		WarehouseCode: "123WC",
		LocalityId:    1,
	}

	db, err := createDataBase()
	if err != nil {
		fmt.Println(err)
	}
	var repository = &repository_test{db, false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	res, _ := service.Get(ctx, 1)

	assert.Equal(t, expected_res, res, "deben ser iguales")
	assert.True(t, repository.flag)
}

//Method that checks if the service get works fine when id does not exist
func TestGetIdNotExist(t *testing.T) {
	expected_res := errors.New("el id no existe en la base de datos")

	db, err := createDataBase()
	if err != nil {
		fmt.Println(err)
	}

	var repository = &repository_test{db, false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	_, e := service.Get(ctx, 3)

	assert.Equal(t, expected_res, e, "deben ser iguales")
	assert.True(t, repository.flag)
}

//Method that checks if the service getAll works fine when is data in the database
func TestGetAllOk(t *testing.T) {
	expected_res := []domain.Warehouse{{
		ID:            1,
		Address:       "Lavalle 467",
		Telephone:     "358 4329528",
		WarehouseCode: "IT123",
		LocalityId:    1,
	}, {
		ID:            2,
		Address:       "Colombia 1467",
		Telephone:     "358 4337528",
		WarehouseCode: "OP456",
		LocalityId:    5,
	}}

	db, err := createDataBase()
	if err != nil {
		fmt.Println(err)
	}

	var repository = &repository_test{db, false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	res, _ := service.GetAll(ctx)

	assert.Equal(t, expected_res, res, "deben ser iguales")
	assert.True(t, repository.flag)
}

//Method that checks if the service delete works fine when id exist
func TestDeleteOk(t *testing.T) {

	db, err := createDataBase()
	if err != nil {
		fmt.Println(err)
	}

	var repository = &repository_test{db, false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	res := service.Delete(ctx, 1)

	assert.Equal(t, nil, res, "deben ser iguales")
	assert.True(t, repository.flag)
}

//Method that checks if the service delete works fine when id does not exist
func TestDeleteIdNotExist(t *testing.T) {
	expected_res := errors.New("el id es inexistente")
	db, err := createDataBase()
	if err != nil {
		fmt.Println(err)
	}

	var repository = &repository_test{db, false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	res := service.Delete(ctx, 2)

	assert.Equal(t, expected_res, res, "deben ser iguales")
	assert.True(t, repository.flag)
}

// //Method that checks if the service create works fine when the data is fine
func TestCreateOk(t *testing.T) {

	db, err := createDataBase()
	if err != nil {
		fmt.Println(err)
	}

	var repository = &repository_test{db, false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	newW := domain.Warehouse{
		Address:       "Lavalle 467",
		Telephone:     "358 4329528",
		WarehouseCode: "IT1234&",
	}

	res, _ := service.Save(ctx, newW)

	assert.Equal(t, 1, res, "deben ser iguales")
	assert.True(t, repository.flag)
}

// //Method that checks if the service create works fine when the data have a warehousecode existent
func TestCreateCodeWarehouseExistent(t *testing.T) {

	db, err := createDataBase()
	if err != nil {
		fmt.Println(err)
	}

	var repository = &repository_test{db, false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	newW := domain.Warehouse{
		Address:       "Lavalle 467",
		Telephone:     "358 4329528",
		WarehouseCode: "123WC",
	}

	res, _ := service.Save(ctx, newW)

	assert.Equal(t, -1, res, "deben ser iguales")
	assert.True(t, repository.flag)
}

// //Method that checks if the service update works fine when the data is fine
func TestUpdateOk(t *testing.T) {

	db, err := createDataBase()
	if err != nil {
		fmt.Println(err)
	}

	var repository = &repository_test{db, false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	newW := domain.Warehouse{
		Address:       "Corrientes 63",
		Telephone:     "358 4329528",
		WarehouseCode: "123WC",
	}
	newW.ID = 1
	er := service.Update(ctx, newW)

	assert.Equal(t, nil, er, "deben ser iguales")
	assert.True(t, repository.flag)
}

//Method that checks if the service update works fine when the id is inexistent
func TestUpdateIdNotExist(t *testing.T) {

	db, err := createDataBase()
	if err != nil {
		fmt.Println(err)
	}

	var repository = &repository_test{db, false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	newW := domain.Warehouse{
		Address:       "Corrientes 63",
		Telephone:     "358 4329528",
		WarehouseCode: "123WC",
	}
	newW.ID = 3
	er := service.Update(ctx, newW)
	expected_err := errors.New("item inexistente")
	assert.Equal(t, expected_err, er, "deben ser iguales")
	assert.True(t, repository.flag)
}

// //Method that checks if the service update works fine when the warehousecode exist
func TestUpdateWarehouseExist(t *testing.T) {

	db, err := createDataBase()
	if err != nil {
		fmt.Println(err)
	}

	var repository = &repository_test{db, false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	newW := domain.Warehouse{
		Address:       "Corrientes 63",
		Telephone:     "358 4329528",
		WarehouseCode: "123WC",
	}
	newW.ID = 2
	er := service.Update(ctx, newW)
	expected_err := errors.New("ya existe un elemento con ese warehouse code")
	assert.Equal(t, expected_err, er, "deben ser iguales")
	assert.True(t, repository.flag)
}
