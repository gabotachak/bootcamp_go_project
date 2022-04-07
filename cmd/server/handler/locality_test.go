package handler

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
* Method that checks if the post endpoint works fine
* End to end test
 */
func TestCreateLocalityOk(t *testing.T) {
	r := CreateServer()

	//body ok
	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/localities/",
		`{            
			"locality_name": "Prueba",
			"province_name": "Prueba",
			"country_name": "PRUEBA"
		}`)

	r.ServeHTTP(rr, req)
	assert.Equal(t, 201, rr.Code)
}

/**
* Method that checks if the post endpoint has failed
* End to end test
 */
func TestCreateLocalityFail(t *testing.T) {
	r := CreateServer()

	//body ok
	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/localities/",
		`{        
			"locality_name": "Pruebañ",    
			"country_name": "Prueba"
		}`)
	reqVal, err := CreateRequestTest(http.MethodPost, "/api/v1/localities/",
		`{
			"locality_name": "Prueba",
			"province_name": "\@",
		}`)

	r.ServeHTTP(rr, req)
	assert.Equal(t, 422, rr.Code)

	r.ServeHTTP(err, reqVal)
	assert.Equal(t, 422, err.Code)
}

// Method that checks if the endpoint get the report of a locality with n sellers
// End to end test
func TestGetReport(t *testing.T) {
	// create server
	r := CreateServer()
	// create request of type get and response for get the result
	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/localities/reportSellers/1", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}

// Method that checks if the endpoint failed to get the report of a locality with n sellers
// End to end test
func TestGetReportFail(t *testing.T) {
	// create server
	r := CreateServer()
	// create request of type get and response for get the result
	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/localities/reportSellers/100", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 404, rr.Code)
}

// Method that checks if the endpoint get the report of all localities with n sellers
// End to end test
func TestGetGeneralReport(t *testing.T) {
	// create server
	r := CreateServer()
	// create request of type get and response for get the result
	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/localities/reportSellers/", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}

// Method that checks if the endpoint get the report of all localities with n sellers
// End to end test
func TestGetGeneralReportFail(t *testing.T) {
	// create server
	r := CreateServer()
	// create request of type get and response for get the result
	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/localities/reportSellers/", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 404, rr.Code)
}
