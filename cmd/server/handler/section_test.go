package handler

import (
	"net/http"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func Test_create_ok_section(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/sections/", `
	{
		"section_number" : "2",
		"current_temperature": 22,
		"minimum_temperature": 22,
		"current_capacity": 22,
		"minimum_capacity": 22,
		"maximum_capacity": 22,
		"warehouse_id": 1,
		"product_type_id": 1
	}
	`)
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)
}

func Test_create_failed_section(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/sections/", `
	{
		"current_temperature": 2,
		"minimum_temperature": 2,
		"current_capacity": 2,
		"minimum_capacity": 2,
		"maximum_capacity": 2,
		"warehouse_id": 2,
		"product_type_id": 2
	}
	`)
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
}

func Test_create_conflict_section(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/sections/", `
	{
		"section_number" : "1",
		"current_temperature": 11,
		"minimum_temperature": 11,
		"current_capacity": 11,
		"minimum_capacity": 11,
		"maximum_capacity": 11,
		"warehouse_id": 1,
		"product_type_id": 1
	}
	`)
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusConflict, rr.Code)
}

func Test_find_all_section(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/sections/", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func Test_find_by_id_non_existent_section(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/sections/9", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func Test_find_by_id_existent_section(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/sections/1", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func Test_update_ok_section(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodPatch, "/api/v1/sections/2", `
	{
		"section_number" : "2",
		"current_temperature": 222,
		"minimum_temperature": 222,
		"current_capacity": 222,
		"minimum_capacity": 222,
		"maximum_capacity": 222,
		"warehouse_id": 1,
		"product_type_id": 1
	}
	`)
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func Test_update_non_existent_section(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodPatch, "/api/v1/sections/9", `
	{
		"section_number": "9",
		"current_temperature": 9,
		"minimum_temperature": 9,
		"current_capacity": 9,
		"minimum_capacity": 9,
		"maximum_capacity": 9,
		"warehouse_id": 9,
		"product_type_id": 9
	}
	`)
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func Test_update_url_invalid_section(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodPatch, "/api/v1/sections/abc", `
	{
		"section_number" : "2",
		"current_temperature": 22,
		"minimum_temperature": 22,
		"current_capacity": 22,
		"minimum_capacity": 22,
		"maximum_capacity": 22,
		"warehouse_id": 1,
		"product_type_id": 1
	}
	`)
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func Test_update_json_invalid_section(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodPatch, "/api/v1/sections/2", `
	{
		"section_number: 9,
		"current_temperature: 9,
		"minimum_temperature: 9,
		"current_capacity": 9,
		"minimum_capacity": 9,
		"maximum_capacity": 9,
		"warehouse_id": 9,
		"product_type_id": 9
	}
	`)
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
}

func Test_delete_non_existent_section(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodDelete, "/api/v1/sections/9", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func Test_delete_ok_section(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodDelete, "/api/v1/sections/2", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNoContent, rr.Code)
}

func Test_delete_url_invalid_section(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodDelete, "/api/v1/sections/abc", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
