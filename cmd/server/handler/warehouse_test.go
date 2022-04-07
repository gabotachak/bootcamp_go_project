package handler

import (
	"net/http"
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

// Method that checks if the post endpoint works fine
// End to end test
func TestCreateOkWarehouse(t *testing.T) {
	r := CreateServer()

	//body ok
	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/warehouses/",
		`{"address": "Colombia 55",
		"telephone": "474343240",
		"warehouse_code": "DrewaspsdsyrY00541266630",
		"locality_id": 1}`)
	r.ServeHTTP(rr, req)
	assert.Equal(t, 201, rr.Code)
}

// Method that checks if the post endpoint fails when type is wrong
// End to end test
func TestCreateConflictW(t *testing.T) {
	r := CreateServer()

	//body ok
	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/warehouses/",
		`{"address": "ñññ",
		"telephone": "474343240",
		"warehouse_code": "DrewaspsdsyrY00541266630",
		"locality_id": 1}`)
	r.ServeHTTP(rr, req)
	assert.Equal(t, 400, rr.Code)
}

// Method that checks if the post endpoint works fine when the body is not complete
// End to end test
func TestCreateFailWarehouse(t *testing.T) {
	r := CreateServer()

	// Incomplete body
	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/warehouses/",
		`{"address": "Colombia 55",
		"telephone": "474343240",
		"locality_id": 1}`)

	r.ServeHTTP(rr, req)
	assert.Equal(t, 422, rr.Code)
}

// Method that checks if the post endpoint works fine when the warehouse_code exist
// End to end test
func TestCreateConflictWarehouse(t *testing.T) {
	r := CreateServer()

	// warehouse_code existent into body
	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/warehouses/",
		`{"address": "Colombia 55",
		"telephone": "474343240",
		"warehouse_code": "23435",
		"locality_id": 1}`)

	r.ServeHTTP(rr, req)
	fmt.Println(rr.Body)
	assert.Equal(t, 409, rr.Code)
}

// Method that checks if the endpoint getAll works fine when exist data in the database
// End to end test
func TestGetOkWarehouse(t *testing.T) {
	// create server
	r := CreateServer()

	// create request of type get and response for get the result
	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/warehouses/", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}

//Method that checks if the get endpoint works fine when dont exist data in the database
// End to end test
func TestGetOneInexistentWarehouse(t *testing.T) {
	// create server
	r := CreateServer()

	// create request of type get and response for get the result
	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/warehouses/100", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 404, rr.Code)
}

//Method that checks if the get endpoint works fine when exist data in the database
// End to end test
func TestGetOneOkWarehouse(t *testing.T) {
	// create server
	r := CreateServer()

	// create request of type get and response for get the result
	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/warehouses/1", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}

// Method that checks if the put endpoint works fine
// End to end test
func TestUpdateOkWarehouse(t *testing.T) {
	r := CreateServer()

	//body ok
	req, rr := CreateRequestTest(http.MethodPatch, "/api/v1/warehouses/3",
		`{"address": "Colombia 55",
		"telephone": "474343333240",
		"warehouse_code": "DY005po1266630",
		"locality_id": 1}`)
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}

// Method that checks if the put endpoint works fine when warehouse dont exist
// End to end test
func TestUpdateFailWarehouse(t *testing.T) {
	r := CreateServer()

	//warehouse not exist
	req, rr := CreateRequestTest(http.MethodPatch, "/api/v1/warehouses/100",
		`{"address": "Colombia 55",
		"telephone": "474343333240",
		"warehouse_code": "DY005po1266630",
		"locality_id": 1}`)
	r.ServeHTTP(rr, req)
	assert.Equal(t, 404, rr.Code)
}

// Method that checks if the put endpoint works fine when Some parameter does not exist
// End to end test
func TestUpdateErrorBodyWarehouse(t *testing.T) {
	r := CreateServer()

	// Some parameter does not exist
	req, rr := CreateRequestTest(http.MethodPatch, "/api/v1/warehouses/1",
		`{"address": "Colombia 55",
	"telephone": "474343333240",
	"locality_id": 1}`)
	r.ServeHTTP(rr, req)
	assert.Equal(t, 422, rr.Code)
}

//Method that checks if the delete endpoint works fine when exist data in the database
// End to end test
func TestDeleteOKWarehouse(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodDelete, "/api/v1/warehouses/3", "")

	r.ServeHTTP(rr, req)
	assert.Equal(t, 204, rr.Code)
}

//Method that checks if the delete endpoint works fine when dont exist data in the database
// End to end test
func TestDeleteInexistentWarehouse(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodDelete, "/api/v1/warehouses/100", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 404, rr.Code)
}

//Method that checks if the delete endpoint works fine when pass a invalid id
// End to end test
func TestDeleteInvalidIdWarehouse(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodDelete, "/api/v1/warehouses/sdy", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 400, rr.Code)
}
