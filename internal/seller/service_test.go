package seller

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type repository_test struct {
	db   *sql.DB
	flag bool
}

func (r *repository_test) GetAll(ctx context.Context) ([]domain.Seller, error) {
	r.flag = true
	sellers := []domain.Seller{
		{ID: 1, CID: 1, CompanyName: "Meli", Address: "CDMX-MX", Telephone: "1234567890", LocalityID: 1},
		{ID: 2, CID: 2, CompanyName: "Meli", Address: "CDMX-MX", Telephone: "1234567890", LocalityID: 1},
	}
	return sellers, nil
}

func (r *repository_test) Get(ctx context.Context, id int) (domain.Seller, error) {
	r.flag = true
	if id == 1 {
		seller := domain.Seller{ID: 1, CID: 2, CompanyName: "Meli", Address: "CDMX-MX", Telephone: "1234567890", LocalityID: 1}
		return seller, nil
	}

	return domain.Seller{}, errors.New("el id no existe en la base de datos")
}

func (r *repository_test) Exists(ctx context.Context, cid int) bool {
	r.flag = true
	return cid == 1
}

func (r *repository_test) Save(ctx context.Context, sel domain.Seller) (int, error) {
	r.flag = true
	if sel.CID < 0 || sel.CompanyName == "" || sel.Address == "" || sel.LocalityID < 0 {
		return 0, errors.New("error: datos faltantes o incorrectos")
	}
	if sel.CID == 117 {
		return -1, errors.New("el cid ya existe")
	}
	return 1, nil
}

func (r *repository_test) Update(ctx context.Context, sel domain.Seller) error {
	r.flag = true
	if sel.ID == -1 {
		return errors.New("el id no existe")
	}
	return nil
}

func (r *repository_test) Delete(ctx context.Context, id int) error {
	r.flag = true
	if id < 0 {
		return errors.New("el id no existe")
	}
	return nil
}

type storeDummy struct {
	db_test *sql.DB
}

func (sd *storeDummy) openDbTest() {
	var err error
	sd.db_test, err = sql.Open("sqlite3", "../.././meli_test.db")
	if err != nil {
		fmt.Println("error al inicializar la base de datos", err)
	}
}

func TestSave(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	store := storeDummy{}
	store.openDbTest()

	var repository = &repository_test{store.db_test, false}
	service := NewService(repository)

	newSeller := domain.Seller{CID: 2, CompanyName: "Meli", Address: "CDMX-MX", Telephone: "1234567890", LocalityID: 1}

	res, _ := service.Save(ctx, newSeller)

	assert.Equal(t, 1, res, "deben ser iguales")
	assert.True(t, repository.flag)
}

func TestSaveError(t *testing.T) {

	errorEsperado := errors.New("el cid ya existe")
	store := storeDummy{}
	store.openDbTest()
	var repository = &repository_test{store.db_test, false}
	service := NewService(repository)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	newSeller := domain.Seller{CID: 117, CompanyName: "Meli", Address: "CDMX-MX", Telephone: "1234567890", LocalityID: 1}

	_, err := service.Save(ctx, newSeller)
	assert.Equal(t, errorEsperado, err, "deben ser iguales")
	assert.True(t, repository.flag)

}

func TestGetAll(t *testing.T) {
	sellerEsperado := []domain.Seller{
		{ID: 1, CID: 1, CompanyName: "Meli", Address: "CDMX-MX", Telephone: "1234567890", LocalityID: 1},
		{ID: 2, CID: 2, CompanyName: "Meli", Address: "CDMX-MX", Telephone: "1234567890", LocalityID: 1},
	}
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	store := storeDummy{}
	store.openDbTest()

	var repository = &repository_test{store.db_test, false}
	service := NewService(repository)
	res, err := service.GetAll(ctx)

	assert.Equal(t, sellerEsperado, res, "deben ser iguales")
	assert.True(t, repository.flag)
	assert.Nil(t, err)

}

func TestGetNotFound(t *testing.T) {
	errorEsperado := errors.New("el id no existe en la base de datos")
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	store := storeDummy{}
	store.openDbTest()

	var repository = &repository_test{store.db_test, false}
	service := NewService(repository)

	_, err := service.Get(ctx, 2)

	assert.Equal(t, errorEsperado, err, "deben ser iguales")
	assert.True(t, repository.flag)

}

func TestGet(t *testing.T) {
	sellerEsperado := domain.Seller{ID: 1, CID: 2, CompanyName: "Meli", Address: "CDMX-MX", Telephone: "1234567890", LocalityID: 1}
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	store := storeDummy{}
	store.openDbTest()

	var repository = &repository_test{store.db_test, false}
	service := NewService(repository)

	res, _ := service.Get(ctx, 1)

	assert.Equal(t, sellerEsperado, res, "deben ser iguales")
	assert.True(t, repository.flag)

}

func TestUpdate(t *testing.T) {
	newSeller := domain.Seller{ID: 23, CID: 1000, CompanyName: "Meli", Address: "CDMX-MX", Telephone: "1234567890", LocalityID: 1}
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	store := storeDummy{}
	store.openDbTest()

	var repository = &repository_test{store.db_test, false}
	service := NewService(repository)

	err := service.Update(ctx, newSeller)

	assert.Equal(t, nil, err, "deben ser iguales")
	assert.True(t, repository.flag)

}

func TestUpdateNotFound(t *testing.T) {
	errorEsperado := errors.New("el id no existe")
	newSeller := domain.Seller{ID: -1, CID: -111, CompanyName: "Meli", Address: "CDMX-MX", Telephone: "1234567890", LocalityID: 1}
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	store := storeDummy{}
	store.openDbTest()

	var repository = &repository_test{store.db_test, false}
	service := NewService(repository)

	err := service.Update(ctx, newSeller)

	assert.Equal(t, errorEsperado, err, "deben ser iguales")
	assert.True(t, repository.flag)

}

func TestDeleteNotFound(t *testing.T) {
	errorEsperado := errors.New("el id no existe")
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	store := storeDummy{}
	store.openDbTest()

	var repository = &repository_test{store.db_test, false}
	service := NewService(repository)

	err := service.Delete(ctx, -1)

	assert.Equal(t, errorEsperado, err, "deben ser iguales")
	assert.True(t, repository.flag)

}

func TestDelete(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	store := storeDummy{}
	store.openDbTest()

	var repository = &repository_test{store.db_test, false}
	service := NewService(repository)

	err := service.Delete(ctx, 1)

	assert.Equal(t, nil, err, "deben ser iguales")
	assert.True(t, repository.flag)

}
