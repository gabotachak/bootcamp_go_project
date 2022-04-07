package handler

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_store_inbound_order(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/InboundOrders/", `
	{
		"order_date": "20200712",
		"order_number": "123",
		"employee_id": 1,
		"product_batch_id": 1,
		"warehouse_id": 1
	}
	`)
	r.ServeHTTP(rr, req)
	fmt.Println(req.Body)
	assert.Equal(t, http.StatusCreated, rr.Code)
}

func Test_store_inbound_order_conflict(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/InboundOrders/", `
	{
		"order_date": "20200712",
		"order_number": "123",
		"employee_id": 1,
		"product_batch_id": 1,
		"warehouse_id": 5
	}
	`)
	r.ServeHTTP(rr, req)
	fmt.Println(req.Body)
	assert.Equal(t, http.StatusConflict, rr.Code)
}

func Test_store_inbound_order_json_conflict(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/InboundOrders/", `
	{
		"order_date": "20200712",
		"order_number": "123",
		"employee_id": 1,
		"product_batch_id": 1,
		"warehouse_id": "aa"
	}
	`)
	r.ServeHTTP(rr, req)
	fmt.Println(req.Body)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
}

func Test_get_all_inbound_orders(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/InboundOrders/", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
func Test_get_inbound_order_by_id_existent(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/InboundOrders/1", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func Test_get_inbound_order_by_id_non_existent(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/InboundOrders/99999", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}
