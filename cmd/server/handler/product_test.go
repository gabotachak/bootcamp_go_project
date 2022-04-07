package handler

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

//Test correct insertion of a product
func TestCreateOkProduct(t *testing.T) {
	//Create server and define routes
	r := CreateServer()
	//Create a Create and Response Request to get the result
	req, respRecorder := CreateRequestTest(http.MethodPost, "/api/v1/products/", `
	{
		"description": "Queso",
		"expiration_rate": 100,
		"freezing_rate": 3,
		"height": 5,
		"length": 1,
		"net_weight": 8.1,
		"product_code": "PROD078",
		"recommended_freezing_temperature": 1,
		"width": 4,
		"product_type_id": 1,
		"seller_id": 1
	}`)
	//Identify server that could serve the request
	r.ServeHTTP(respRecorder, req)
	assert.Equal(t, 201, respRecorder.Code)
}

//Test failure of insertion when a required field is missing
func TestCreateFailProduct(t *testing.T) {
	//Create server and define routes
	r := CreateServer()
	//Create a Create and Response Request to get the result
	req, respRecorder := CreateRequestTest(http.MethodPost, "/api/v1/products/", `
	{
		"description": "",
		"expiration_rate": 100,
		"freezing_rate": 3,
		"height": 5,
		"length": 1,
		"net_weight": 8.1,
		"product_code": "PROD078",
		"recommended_freezing_temperature": 1,
		"width": 4,
		"product_type_id": 1,
		"seller_id": 1
	}`)
	//Identify server that could serve the request
	r.ServeHTTP(respRecorder, req)
	assert.Equal(t, 422, respRecorder.Code)
}

//Test failure of insertion when Product Code already exists
func TestCreateConflict(t *testing.T) {
	//Create server and define routes
	r := CreateServer()
	//Create a Create and Response Request to get the result
	req, respRecorder := CreateRequestTest(http.MethodPost, "/api/v1/products/", `
	{
		"description": "Queso",
		"expiration_rate": 100,
		"freezing_rate": 3,
		"height": 5,
		"length": 1,
		"net_weight": 8.1,
		"product_code": "PROD078",
		"recommended_freezing_temperature": 1,
		"width": 4,
		"product_type_id": 1,
		"seller_id": 1
	}`)
	//Identify server that could serve the request
	r.ServeHTTP(respRecorder, req)
	assert.Equal(t, 409, respRecorder.Code)
}

func TestGetAllOkProduct(t *testing.T) {
	//Create server and define routes
	r := CreateServer()
	//Create a Create and Response Request to get the result
	req, respRecorder := CreateRequestTest(http.MethodGet, "/api/v1/products/", "")
	//Identify server that could serve the request
	r.ServeHTTP(respRecorder, req)
	assert.Equal(t, 200, respRecorder.Code)
}

func TestFindByIdNonExistentProduct(t *testing.T) {
	//Create server and define routes
	r := CreateServer()
	//Create a Create and Response Request to get the result
	req, respRecorder := CreateRequestTest(http.MethodGet, "/api/v1/products/888", "")
	//Identify server that could serve the request
	r.ServeHTTP(respRecorder, req)
	assert.Equal(t, 404, respRecorder.Code)
}

func TestFindByIdExistentProduct(t *testing.T) {
	//Create server and define routes
	r := CreateServer()
	//Create a Create and Response Request to get the result
	req, respRecorder := CreateRequestTest(http.MethodGet, "/api/v1/products/1", "")
	//Identify server that could serve the request
	r.ServeHTTP(respRecorder, req)
	assert.Equal(t, 200, respRecorder.Code)
}

func TestUpdateOkProduct(t *testing.T) {
	//Create server and define routes
	r := CreateServer()
	//Create a Create and Response Request to get the result
	req, respRecorder := CreateRequestTest(http.MethodPatch, "/api/v1/products/2", `
	{
		"description": "Queso",
		"expiration_rate": 100,
		"freezing_rate": 3,
		"height": 99,
		"length": 1,
		"net_weight": 8.1,
		"product_code": "PRODXXXX",
		"recommended_freezing_temperature": 1,
		"width": 4,
		"product_type_id": 1,
		"seller_id": 1
	}`)
	//Identify server that could serve the request
	r.ServeHTTP(respRecorder, req)
	assert.Equal(t, 200, respRecorder.Code)
}

func TestUpdateNonExistenteProduct(t *testing.T) {
	//Create server and define routes
	r := CreateServer()
	//Create a Create and Response Request to get the result
	req, respRecorder := CreateRequestTest(http.MethodPatch, "/api/v1/products/99", `
	{
		"description": "LuegoUpdate",
		"expiration_rate": 100,
		"freezing_rate": -3,
		"height": 5,
		"length": 20.5,
		"netweight": 8.1,
		"product_code": "PROD01",
		"recommended_freezing_temperature": 1,
		"width": 4,
		"product_type_id": 8,
		"seller_id": 3
	}`)
	//Identify server that could serve the request
	r.ServeHTTP(respRecorder, req)
	assert.Equal(t, 404, respRecorder.Code)
}

func TestUpdateEmptyFieldProduct(t *testing.T) {
	//Create server and define routes
	r := CreateServer()
	//Create a Create and Response Request to get the result
	req, respRecorder := CreateRequestTest(http.MethodPatch, "/api/v1/products/2", `
	{
		"description": 4,
		"expiration_rate": 100,
		"freezing_rate": -3,
		"height": 5,
		"length": 20.5,
		"netweight": 8.1,
		"product_code": "PROD01",
		"recommended_freezing_temperature": 1,
		"width": 4,
		"product_type_id": 8,
		"seller_id": 3
	}`)
	//Identify server that could serve the request
	r.ServeHTTP(respRecorder, req)
	assert.Equal(t, 404, respRecorder.Code)
}

func TestDeleteNonExistent(t *testing.T) {
	//Create server and define routes
	r := CreateServer()
	//Create a Create and Response Request to get the result
	req, respRecorder := CreateRequestTest(http.MethodDelete, "/api/v1/products/5555", "")
	//Identify server that could serve the request
	r.ServeHTTP(respRecorder, req)
	assert.Equal(t, 404, respRecorder.Code)
}

func TestDeleteWrongTypeProduct(t *testing.T) {
	//Create server and define routes
	r := CreateServer()
	//Create a Create and Response Request to get the result
	req, respRecorder := CreateRequestTest(http.MethodDelete, "/api/v1/products/j", "")
	//Identify server that could serve the request
	r.ServeHTTP(respRecorder, req)
	assert.Equal(t, 404, respRecorder.Code)
}

func TestDeleteOkProduct(t *testing.T) {
	//Create server and define routes
	r := CreateServer()
	//Create a Create and Response Request to get the result
	req, respRecorder := CreateRequestTest(http.MethodDelete, "/api/v1/products/2", "")
	//Identify server that could serve the request
	r.ServeHTTP(respRecorder, req)
	assert.Equal(t, 204, respRecorder.Code)
}
