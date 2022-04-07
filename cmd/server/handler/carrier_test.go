package handler

import (
	"net/http"
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

// Method that checks if the post endpoint works fine
// End to end test
func TestCreateOkCarrier(t *testing.T) {
	r := CreateServer()

	//body ok
	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/carriers/",
		`{            
			"cid": "1234",
			"company_name": "carrier 1",
			"address": "Belgrano 1931",
			"telephone": "358 322432",
			"locality_id": 1
		}`)

	r.ServeHTTP(rr, req)
	assert.Equal(t, 201, rr.Code)
}

// Method that checks if the post endpoint works fine when the locality is not exist
// End to end test
func TestCreateWithOutLocalityCarrier(t *testing.T) {
	r := CreateServer()

	// Incomplete body
	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/carriers/",
		`{           
		"cid": "1235", 
		"company_name": "carrier 1",
		"address": "Belgrano 1931",
		"telephone": "324324",
		"locality_id": 100
	}`)

	r.ServeHTTP(rr, req)
	assert.Equal(t, 409, rr.Code)
}

// Method that checks if the post endpoint works fine when the body is not complete
// End to end test
func TestCreateFailCarrier(t *testing.T) {
	r := CreateServer()

	// Incomplete body
	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/carriers/",
		`{            
		"company_name": "carrier 1",
		"address": "Belgrano 1931",
		"telephone": "358 322432",
		"locality_id": 1
	}`)

	r.ServeHTTP(rr, req)
	assert.Equal(t, 422, rr.Code)
}

// // Method that checks if the post endpoint works fine when the cid exist
// // End to end test
func TestCreateConflictCarrier(t *testing.T) {
	r := CreateServer()

	// warehouse_code existent into body
	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/carriers/",
		`{            
		"cid": "1234",
		"company_name": "carrier 1",
		"address": "Belgrano 1931",
		"telephone": "358 322432",
		"locality_id": 1
	}`)

	r.ServeHTTP(rr, req)
	fmt.Println(rr.Body)
	assert.Equal(t, 409, rr.Code)
}

// Method that checks if the endpoint get all report works fine when exist data in the database
// End to end test
func TestGetOkCarrier(t *testing.T) {
	// create server
	r := CreateServer()
	// create request of type get and response for get the result
	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/localities/reportCarriers/", "")
	r.ServeHTTP(rr, req)
	fmt.Println(rr)
	assert.Equal(t, 200, rr.Code)
}

// Method that checks if the endpoint get all report wth details works fine when exist data in the database
// End to end test
func TestGetReportDetailsOkCarrier(t *testing.T) {
	// create server
	r := CreateServer()
	// create request of type get and response for get the result
	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/localities/reportCarriers/1/details", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}

// Method that checks if the endpoint get all report wth details works fine when not exist data in the database
// End to end test
func TestGetReportDetailsFailCarrier(t *testing.T) {
	// create server
	r := CreateServer()
	// create request of type get and response for get the result
	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/localities/reportCarriers/100/details", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 404, rr.Code)
}

// Method that checks if the get endpoint works fine when dont exist data in the database
// End to end test
func TestGetOneInexistentCarrier(t *testing.T) {
	// create server
	r := CreateServer()

	// create request of type get and response for get the result
	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/localities/reportCarriers/100", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 404, rr.Code)
}

// //Method that checks if the get endpoint works fine when exist data in the database
// // End to end test
func TestGetOneOkCarrier(t *testing.T) {
	// create server
	r := CreateServer()

	// create request of type get and response for get the result
	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/localities/reportCarriers/1", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}
