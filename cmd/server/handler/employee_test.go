package handler

import (
	"fmt"
	"net/http"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

// Method check if data is sucessfully store in database
func Test_create_ok_employee(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/employees/", `
	{
		"card_number_id": "1",
		"first_name":"Jhon",
		"last_name":"Doe",
		"warehouse_id":1 
	}
	`)
	messageExpected := "{\"code\":\"201\",\"data\":\"El empleado ha sido creado correctamente con id: 3\"}"
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, messageExpected, rr.Body.String())
}

// Method check if data is validate by type data card_number_id not accept "space"
func Test_create_failed_validate_type_employee(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/employees/", `
	{
		"card_number_id": "1 ", 
		"first_name":"Jhon",
		"last_name":"Doe",
		"warehouse_id":1 
	}
	`)
	messageExpected := "{\"code\":\"422\",\"error\":\"el campo CardNumberID es requerido o su tipo de dato no es el correcto alphanum\"}"
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	assert.Equal(t, messageExpected, rr.Body.String())
}

//Method check if any field is empty o dont send in body request
func Test_create_failed_employee(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/employees/", `
	{
		"card_number_id": "99",
		"first_name":"aaa",
		"last_name":"Doe",
	}
	`)
	r.ServeHTTP(rr, req)
	fmt.Println(rr.Body.String())
	assert.Equal(t, http.StatusOK, rr.Code)
}

// Method check id card_number_id exist
func Test_create_conflict_employee(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/employees/", `
	{
		"card_number_id": "123",
		"first_name":"Jhon",
		"last_name":"Doe",
		"warehouse_id":1 
	}
	`)
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusConflict, rr.Code)
}

//Method check if dara in body is correctly typed
func Test_create_failed_json_employee(t *testing.T) {
	r := CreateServer()
	result_expected := `{"code":"422","error":"Algun dato en el body esta mal ingresado"}`
	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/employees/", `
	{
		"card_number_id": 12345,
		"first_name":"Jhon",
		"last_name":"Doe",
		"warehouse_id":1 
	}
	`)
	r.ServeHTTP(rr, req)

	assert.Equal(t, result_expected, rr.Body.String())
}

// Method check if GetAll workly correctly
func Test_find_all_employee(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/employees/", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

//Method check if id non exists
func Test_find_by_id_non_existent_employee(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/employees/2", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

//Method check if id exists
func Test_find_by_id_existent_employee(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/employees/123", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

// Method check if Update workly correctly
func Test_update_ok_employee(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodPatch, "/api/v1/employees/123", `
	{
		"first_name":"Jhon2",
		"last_name":"Doe2",
		"warehouse_id":1 
	}
	`)
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

// Method check if id not exist
func Test_update_non_existent_employee(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodPatch, "/api/v1/employees/99999", `
	{
		"first_name":"Jhon2",
		"last_name":"Doe2",
		"warehouse_id":1 
	}
	`)
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

// Method check if data correctly typed
func Test_updated_failed_json_employee(t *testing.T) {
	r := CreateServer()
	result_expected := `{"code":"400","error":"Algun dato en el body esta mal ingresado"}`
	req, rr := CreateRequestTest(http.MethodPatch, "/api/v1/employees/1", `
	{
		"first_name":"Jhon",
		"last_name":1234*?
	}
	`)
	r.ServeHTTP(rr, req)

	assert.Equal(t, result_expected, rr.Body.String())
}

// Method check if data is validate by type data card_number_id not accept "space"
func Test_updated_failed_validate_type_employee(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/employees/", `
	{
		"card_number_id": "1 ", 
		"first_name":"Jhon",
		"last_name":"Doe",
		"warehouse_id":1 
	}
	`)
	messageExpected := "{\"code\":\"422\",\"error\":\"el campo CardNumberID es requerido o su tipo de dato no es el correcto alphanum\"}"
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	assert.Equal(t, messageExpected, rr.Body.String())
}

//Method check if id not exists
func Test_delete_non_existent_employee(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodDelete, "/api/v1/employees/999", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

// Method check if delete workly correctly
func Test_delete_ok_employee(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodDelete, "/api/v1/employees/1234", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNoContent, rr.Code)
}

//method check if method GetInboundOrders workly correctly
func Test_get_inbound_orders(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/employees/reportInboundOrders/", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

//method check if method GetInboundOrdersByEmployee workly correctly
func Test_get_inbound_orders_by_employee(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/employees/reportInboundOrders/1", "")

	r.ServeHTTP(rr, req)
	fmt.Println(rr.Body)
	assert.Equal(t, http.StatusOK, rr.Code)
}

//method check if method GetInboundOrdersByEmployee workly correctly
func Test_get_inbound_orders_by_employee_non_exists(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/employees/reportInboundOrders/99", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
