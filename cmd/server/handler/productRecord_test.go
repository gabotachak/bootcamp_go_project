package handler

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

func TestGetAllEmpty(t *testing.T) {
	//Create server and define routes
	r := CreateServer()
	//Create a Create and Response Request to get the result
	req, respRecorder := CreateRequestTest(http.MethodGet, "/api/v1/productsRecords/", "")
	//Identify server that could serve the request
	r.ServeHTTP(respRecorder, req)
	assert.Equal(t, 404, respRecorder.Code)
}

//Test correct insertion of a product record
func TestCreateOkProductRecord(t *testing.T) {
	//Create server and define routes
	r := CreateServer()
	//Create a Create and Response Request to get the result
	req, respRecorder := CreateRequestTest(http.MethodPost, "/api/v1/productsRecords/", `
	{
			"last_update_date": "2023-07-19",
			"purchase_price":  10,
			"sale_price":      15,
			"product_id":      1
		
	}`)
	//Identify server that could serve the request
	r.ServeHTTP(respRecorder, req)
	assert.Equal(t, 201, respRecorder.Code)
}

//Test failure of insertion when Product Record has an invalid date
func TestCreateConflictProdRecord(t *testing.T) {
	//Create server and define routes
	r := CreateServer()
	//Create a Create and Response Request to get the result
	req, respRecorder := CreateRequestTest(http.MethodPost, "/api/v1/productsRecords/", `
	{
		"last_update_date": "2019-07-19",
		"purchase_price":  10,
		"sale_price":      15,
		"product_id":      1
	
}`)
	//Identify server that could serve the request
	r.ServeHTTP(respRecorder, req)
	assert.Equal(t, 422, respRecorder.Code)
}

func TestCreateFailProductRecord(t *testing.T) {
	r := CreateServer()

	// Incomplete body
	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/productsRecords/",
		`{
			"last_update_date": gg,
			"purchase_price":  10,
			"product_id":      1
		}`)

	r.ServeHTTP(rr, req)
	assert.Equal(t, 422, rr.Code)
}

func TestCreateFailProductIdNotExist(t *testing.T) {
	r := CreateServer()

	// Incomplete body
	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/productsRecords/",
		`{            
			"last_update_date": "2027-07-19",
			"purchase_price":  10,
			"sale_price":      15,
			"product_id":      999
		}`)

	r.ServeHTTP(rr, req)
	assert.Equal(t, 404, rr.Code)
}

func TestGetAllOkProductRecord(t *testing.T) {
	//Create server and define routes
	r := CreateServer()
	//Create a Create and Response Request to get the result
	req, respRecorder := CreateRequestTest(http.MethodGet, "/api/v1/productsRecords/", "")
	//Identify server that could serve the request
	r.ServeHTTP(respRecorder, req)
	assert.Equal(t, 200, respRecorder.Code)
}

func TestFindByIdExistentProductRecord(t *testing.T) {
	//Create server and define routes
	r := CreateServer()
	//Create a Create and Response Request to get the result
	req, respRecorder := CreateRequestTest(http.MethodGet, "/api/v1/productsRecords/1", "")
	//Identify server that could serve the request
	r.ServeHTTP(respRecorder, req)
	assert.Equal(t, 200, respRecorder.Code)
}

func TestFindByNonExistentProductRecord(t *testing.T) {
	//Create server and define routes
	r := CreateServer()
	//Create a Create and Response Request to get the result
	req, respRecorder := CreateRequestTest(http.MethodGet, "/api/v1/productsRecords/55556", "")
	//Identify server that could serve the request
	r.ServeHTTP(respRecorder, req)
	assert.Equal(t, 404, respRecorder.Code)
}
