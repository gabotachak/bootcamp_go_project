package carrier

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

type repository_test struct {
	db   *sql.DB
	flag bool
}

/**
I define to do the service test by building a mock repository
**/

func (r *repository_test) GetAllReport(ctx context.Context) ([]CarriersByLocality, error) {
	r.flag = true

	carriers := []CarriersByLocality{{
		LocalityId:    "1",
		LocalityName:  "San Francisco",
		CarriersCount: "1",
	},
		{
			LocalityId:    "1",
			LocalityName:  "Rio Cuarto",
			CarriersCount: "3",
		}}
	return carriers, nil
}

func (r *repository_test) GetReportDetails(ctx context.Context, id int) ([]domain.Carrier, error) {
	r.flag = true

	carriers := []domain.Carrier{
		{
			ID:          1,
			CID:         "123",
			CompanyName: "Carry 1",
			Address:     "Lavalle 123",
			Telephone:   "38274324",
			LocalityId:  1,
		},
		{
			ID:          2,
			CID:         "2343r24ss2siCb1",
			CompanyName: "carrier 1",
			Address:     "Belgrano 1931",
			Telephone:   "358 322432",
			LocalityId:  1,
		},
	}
	if id == 1 {
		return carriers, nil
	}
	return []domain.Carrier{}, errors.New("no existe una localidad con ese id")
}

func (r *repository_test) GetReport(ctx context.Context, id int) (CarriersByLocality, error) {
	r.flag = true

	if id == 1 {
		carr := CarriersByLocality{
			LocalityId:    "1",
			LocalityName:  "San Francisco",
			CarriersCount: "1",
		}
		return carr, nil
	}
	return CarriersByLocality{}, errors.New("no existe una localidad con ese id")

}

func (r *repository_test) Store(ctx context.Context, c domain.Carrier) (domain.Carrier, error) {
	r.flag = true
	if c.CID == "123" {
		return domain.Carrier{}, errors.New("ya existe un transportista con cid igual")
	}
	if c.LocalityId == 100 {
		return domain.Carrier{}, errors.New("no existe una localidad con ese id")
	}
	res := domain.Carrier{
		ID:          100,
		CID:         "123",
		CompanyName: "Carry 1",
		Address:     "Lavalle 123",
		Telephone:   "38274324",
		LocalityId:  1,
	}

	return res, nil
}

func (r *repository_test) LocalityExists(ctx context.Context, id int) bool {
	r.flag = true

	return id == 1
}

func (r *repository_test) CIDExists(ctx context.Context, cid string) bool {
	r.flag = true

	return cid == "1"
}

//Method that checks if the service get all works fine when data exist
func TestGetAllOk(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}

	expected := []CarriersByLocality{{
		LocalityId:    "1",
		LocalityName:  "San Francisco",
		CarriersCount: "1",
	},
		{
			LocalityId:    "1",
			LocalityName:  "Rio Cuarto",
			CarriersCount: "3",
		}}

	var repository_t = &repository_test{db, false}
	service := NewService(repository_t)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	res, _ := service.GetAllReport(ctx)

	assert.Equal(t, expected, res, "deben ser iguales")
	assert.True(t, repository_t.flag)
}

//Method that checks if the service get one works fine when data exist
func TestGetOk(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}

	expected := CarriersByLocality{
		LocalityId:    "1",
		LocalityName:  "San Francisco",
		CarriersCount: "1",
	}

	var repository_t = &repository_test{db, false}
	service := NewService(repository_t)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	res, _ := service.GetReport(ctx, 1)

	assert.Equal(t, expected, res, "deben ser iguales")
	assert.True(t, repository_t.flag)
}

//Method that checks if the service get one works fine when data not exist
func TestGetFail(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}

	expected := CarriersByLocality{}
	err_expected := errors.New("no existe una localidad con ese id")

	var repository_t = &repository_test{db, false}
	service := NewService(repository_t)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	res, err := service.GetReport(ctx, 2)

	assert.Equal(t, expected, res, "deben ser iguales")
	assert.Equal(t, err_expected, err, "deben ser iguales")
	assert.True(t, repository_t.flag)
}

//Method that checks if the service get one report works fine when data exist
func TestGetRepDetailsOk(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}

	expected := []domain.Carrier{
		{
			ID:          1,
			CID:         "123",
			CompanyName: "Carry 1",
			Address:     "Lavalle 123",
			Telephone:   "38274324",
			LocalityId:  1,
		},
		{
			ID:          2,
			CID:         "2343r24ss2siCb1",
			CompanyName: "carrier 1",
			Address:     "Belgrano 1931",
			Telephone:   "358 322432",
			LocalityId:  1,
		},
	}

	var repository_t = &repository_test{db, false}
	service := NewService(repository_t)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	res, _ := service.GetReportDetails(ctx, 1)

	assert.Equal(t, expected, res, "deben ser iguales")
	assert.True(t, repository_t.flag)
}

//Method that checks if the service get one report works fine when data not exist
func TestGetReportDetailsFail(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}

	expected := []domain.Carrier{}
	err_expected := errors.New("no existe una localidad con ese id")

	var repository_t = &repository_test{db, false}
	service := NewService(repository_t)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	res, err := service.GetReportDetails(ctx, 2)

	assert.Equal(t, expected, res, "deben ser iguales")
	assert.Equal(t, err_expected, err, "deben ser iguales")
	assert.True(t, repository_t.flag)
}

//Method that checks if the service create works fine when the data have a cid existent
func TestCreateCIDExistent(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	var repository_t = &repository_test{db, false}
	service := NewService(repository_t)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	newCarrier := domain.Carrier{
		Address:     "Lavalle 467",
		Telephone:   "358 4329528",
		CID:         "123",
		CompanyName: "test",
		LocalityId:  1,
	}

	res, err := service.Store(ctx, newCarrier)
	expected := domain.Carrier{}
	err_expected := errors.New("ya existe un transportista con cid igual")
	assert.Equal(t, expected, res, "deben ser iguales")
	assert.Equal(t, err_expected, err, "deben ser iguales")
	assert.True(t, repository_t.flag)
}

//Method that checks if the service create works fine when the data is ok
func TestCreateWorkOk(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	var repository_t = &repository_test{db, false}
	service := NewService(repository_t)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	newCarrier := domain.Carrier{
		CID:         "1234567",
		CompanyName: "Carry 1",
		Address:     "Lavalle 123",
		Telephone:   "38274324",
		LocalityId:  1,
	}
	res, err := service.Store(ctx, newCarrier)

	assert.Equal(t, 100, res.ID)
	assert.NotEmpty(t, res)
	assert.Nil(t, err)
	assert.True(t, repository_t.flag)
}

//Method that checks if the service create works fine when the data have a locality is not existent
func TestCreateLocalityNotExistent(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	var repository_t = &repository_test{db, false}
	service := NewService(repository_t)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	newCarrier := domain.Carrier{
		Address:     "Lavalle 467",
		Telephone:   "358 4329528",
		CID:         "123",
		CompanyName: "test",
		LocalityId:  100,
	}

	res, err := service.Store(ctx, newCarrier)
	expected := domain.Carrier{}
	err_expected := errors.New("no existe una localidad con ese id")

	assert.Equal(t, expected, res, "deben ser iguales")
	assert.Equal(t, err_expected, err, "deben ser iguales")
	assert.True(t, repository_t.flag)
}
