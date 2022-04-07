package locality

import (
	"context"
	"database/sql"
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
type ServiceInterface interface {
	GetReport(ctx context.Context, id int) (LocalityReport, error)
	GetGeneralReport(ctx context.Context) ([]LocalityReport, error)
	Save(ctx context.Context, loc domain.Locality) (domain.Locality, error)
}
*/

type repository_test struct {
	db   *sql.DB
	flag bool
}

func (r *repository_test) GetReport(ctx context.Context, id int) (LocalityReport, error) {
	r.flag = true
	if id == 1 {
		locality := LocalityReport{LocalityId: 1, LocalityName: "Lista 1", Quantity: 3}
		return locality, nil
	}

	return LocalityReport{}, errors.New("el id no existe en la base de datos")
}

func (r *repository_test) GetGeneralReport(ctx context.Context) ([]LocalityReport, error) {
	r.flag = true

	locality := []LocalityReport{
		{LocalityId: 1, LocalityName: "Lista 1", Quantity: 3},
		{LocalityId: 2, LocalityName: "Lista 2", Quantity: 5},
	}
	return locality, nil
}

func (r *repository_test) Save(ctx context.Context, loc domain.Locality) (domain.Locality, error) {
	r.flag = true
	if loc.LocalityName == "" || loc.ProvinceName == "" {
		return domain.Locality{}, errors.New("error: datos faltantes o incorrectos")
	}
	return loc, nil
}

/**
****************************************************************
*  AQUÍ ESTÁ LA COPIA DE SELLER
****************************************************************
 */

func TestGetReport(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}

	var repository = &repository_test{db, false}
	service := NewService(repository)

	localityReportExpected := LocalityReport{LocalityId: 1, LocalityName: "Lista 1", Quantity: 3}

	res, _ := service.GetReport(ctx, 1)

	assert.Equal(t, localityReportExpected, res, "deben ser iguales")
	assert.True(t, repository.flag)

}

func TestGetReportFail(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}

	var repository = &repository_test{db, false}
	service := NewService(repository)

	localityReportExpected := LocalityReport{}

	res, _ := service.GetReport(ctx, 50)

	assert.Equal(t, localityReportExpected, res, "deben ser iguales")
	assert.True(t, repository.flag)

}

func TestGetGeneralReport(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}

	var repository = &repository_test{db, false}
	service := NewService(repository)

	localityReportExpected := []LocalityReport{
		{LocalityId: 1, LocalityName: "Lista 1", Quantity: 3},
		{LocalityId: 2, LocalityName: "Lista 2", Quantity: 5},
	}

	res, _ := service.GetGeneralReport(ctx)

	assert.Equal(t, localityReportExpected, res, "deben ser iguales")
	assert.True(t, repository.flag)

}

func TestSaveLocality(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	var repository = &repository_test{db, false}
	service := NewService(repository)

	newLocality := domain.Locality{LocalityName: "Nombre Prueba", ProvinceName: "Nombre Prueba"}

	res, _ := service.Save(ctx, newLocality)

	assert.Equal(t, newLocality, res, "deben ser iguales")
	assert.True(t, repository.flag)
}

func TestSaveLocalityFail(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	var repository = &repository_test{db, false}
	service := NewService(repository)

	newLocality := domain.Locality{}

	res, _ := service.Save(ctx, newLocality)

	assert.Equal(t, newLocality, res, "deben ser iguales")
	assert.True(t, repository.flag)
}

/*
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
*/
