package handler

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_create_ok_product_batch(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/productBatches/", `
	{
		"batch_number": "xyz",
		"current_quantity": 1,
		"current_temperature": 1,
		"due_date": "2022-04-04",
		"initial_quantity": 1,
		"manufacturing_date": "2020-04-04",
		"manufacturing_hour": "2020-04-04 10:30",
		"minimum_temperature": 1,
		"product_id": 1,
		"section_id": 1
	}
	`)
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)
}

func Test_create_failed_missing_field_product_batch(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/productBatches/", `
	{
		"current_quantity": 1,
		"current_temperature": 1,
		"due_date": "2022-04-04",
		"initial_quantity": 1,
		"manufacturing_date": "2020-04-04",
		"manufacturing_hour": "2020-04-04 10:30",
		"minimum_temperature": 1,
		"product_id": 1,
		"section_id": 1
	}
	`)
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
}

func Test_create_failed_blank_field_product_batch(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/productBatches/", `
	{
		"batch_number": "",
		"current_quantity": 1,
		"current_temperature": 1,
		"due_date": "2022-04-04",
		"initial_quantity": 1,
		"manufacturing_date": "2020-04-04",
		"manufacturing_hour": "2020-04-04 10:30",
		"minimum_temperature": 1,
		"product_id": 1,
		"section_id": 1
	}
	`)
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
}

func Test_create_conflict_product_batch(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/productBatches/", `
	{
		"batch_number": "abc",
		"current_quantity": 1,
		"current_temperature": 1,
		"due_date": "2022-04-04",
		"initial_quantity": 1,
		"manufacturing_date": "2020-04-04",
		"manufacturing_hour": "2020-04-04 10:30",
		"minimum_temperature": 1,
		"product_id": 1,
		"section_id": 1
	}
	`)
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusConflict, rr.Code)
}

func Test_reports_all_product_batch(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/sections/reportProducts/", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func Test_reports_by_id_non_existent_section(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/sections/reportProducts?id=9", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func Test_reports_by_id_existent_section(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/sections/reportProducts?id=1", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func Test_reports_by_id_bad_query_section(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/sections/reportProducts?id=abc", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func Test_reports_by_id_negative_num_section(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/sections/reportProducts?id=-1", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
