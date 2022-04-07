package carrier

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	db "github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// unit test of repository store method when create is successful
// repository unit test
func TestCreateOk(t *testing.T) {

	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	//declare the expected variable
	expected := domain.Carrier{
		CID:         "321",
		CompanyName: "carrier 2",
		Address:     "Belgrano 1931",
		Telephone:   "358 322432",
		LocalityId:  1,
	}

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method store
	carrier, err := repository.Store(ctx, expected)
	if err != nil {
		t.Error(err)
	}

	//assert
	assert.Equal(t, expected.CompanyName, carrier.CompanyName)
	assert.NotEmpty(t, carrier.ID)
}

// unit test of repository store method when create is not successful
// incomplete body
// repository unit test
func TestCreateFail(t *testing.T) {

	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	//declare the expected variable
	expected := domain.Carrier{
		CID:         "123",
		CompanyName: "carrier 1",
		Address:     "Belgrano 1931",
		Telephone:   "358 322432",
		LocalityId:  1,
	}

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method store
	carrier, _ := repository.Store(ctx, expected)

	//assert
	assert.NotEqual(t, expected.CompanyName, carrier.CompanyName)
	assert.Empty(t, carrier.ID)
}

// unit test of repository cid_exists method when cid is exist
// repository unit test
func TestCIDExistOk(t *testing.T) {
	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	//declare the expected variable
	expected := true

	// declare cid
	cid := "123"
	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method store
	res := repository.CIDExists(ctx, cid)

	//assert
	assert.Equal(t, expected, res)
}

// unit test of repository cid_exists method when cid is not exist
// repository unit test
func TestCIDExistFail(t *testing.T) {

	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	// declare cid
	cid := "23pi2sfsdfdaf325345hjtya22"
	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method store
	res := repository.CIDExists(ctx, cid)

	//assert
	assert.False(t, res)
}

// unit test of repository locality_exists method when id is exist
// repository unit test
func TestLocalityExistOk(t *testing.T) {

	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	// declare id
	id := 1
	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method store
	res := repository.LocalityExists(ctx, id)

	//assert
	assert.True(t, res)
}

// unit test of repository locality_exists method when id is not exist
// repository unit test
func TestLocalityExistFail(t *testing.T) {

	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	// declare id
	id := 100
	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method store
	res := repository.LocalityExists(ctx, id)

	//assert
	assert.False(t, res)
}

// unit test of repository get_all_report method when data is exist
// repository unit test
func TestGetAllReportOk(t *testing.T) {

	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method get_all_report
	carriers, err := repository.GetAllReport(ctx)
	if err != nil {
		t.Error(err)
	}

	//assert
	assert.NotEmpty(t, carriers)
}

// unit test of repository get_report method when data is exist
// repository unit test
func TestGetReportOk(t *testing.T) {

	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	// declare id
	id := 1

	//expected
	expected := "1"
	expected1 := "Rio Cuarto"

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method get_report
	carrier, err := repository.GetReport(ctx, id)
	if err != nil {
		t.Error(err)
	}

	//assert
	assert.Equal(t, expected, carrier.LocalityId)
	assert.Equal(t, expected1, carrier.LocalityName)
	assert.NotNil(t, carrier.CarriersCount)
}

// unit test of repository get_report method when data is not exist
// repository unit test
func TestGetReportFail(t *testing.T) {

	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	// declare id
	id := 100

	//expected
	expected := "1"
	expected1 := "Rio Cuarto"
	expected2 := CarriersByLocality{}

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method get_report
	carrier, err := repository.GetReport(ctx, id)
	if err != nil {
		t.Error(err)
	}

	//assert
	assert.NotEqual(t, expected, carrier.LocalityId)
	assert.NotEqual(t, expected1, carrier.LocalityName)
	assert.Equal(t, expected2, carrier)
}

// unit test of repository get_report_detail method when data is exist
// repository unit test
func TestGetReportDetailOk(t *testing.T) {

	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	// declare id
	id := 1

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method get_report_detail
	carriers, err := repository.GetReportDetails(ctx, id)
	if err != nil {
		t.Error(err)
	}

	//assert
	assert.NotEmpty(t, carriers)

}

// unit test of repository get_report_detail method when data is not exist
// repository unit test
func TestGetReportDetailFail(t *testing.T) {

	//initialitation database
	db, err := db.Init("unit_test")
	if err != nil {
		fmt.Println(err)
	}
	repository := NewRepository(db)

	// declare id
	id := 100
	var expected []domain.Carrier

	//initialitation context
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	//call method get_report_detail
	carriers, err := repository.GetReportDetails(ctx, id)
	if err != nil {
		t.Error(err)
	}

	//assert
	assert.Empty(t, carriers)
	assert.Equal(t, expected, carriers)

}
