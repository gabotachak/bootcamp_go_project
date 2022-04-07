package handler

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorePurchaseOrder(t *testing.T) {
	// crear el Server y definir las Rutas
	r := CreateServer()
	// crear Request del tipo GET y Response para obtener el resultado
	reqStorePurchaseOrder, rrStorePurchaseOrder := CreateRequestTest(http.MethodPost, "/api/v1/PurchaseOrders/",
		`{"order_number": "order#3",
		"order_date": "2021-04-04 00:00:00",
		"tracking_code": "trackingCode3",
		"buyer_id": 1,
		"order_status_id": 1,
		"carrier_id": 1,
		"warehouse_id": 1}`)

	// indicar al servidor que pueda atender la solicitud

	fmt.Println(rrStorePurchaseOrder.Code)
	r.ServeHTTP(rrStorePurchaseOrder, reqStorePurchaseOrder)
	assert.Equal(t, http.StatusCreated, rrStorePurchaseOrder.Code)
}

func TestStorePurchaseOrderConflict(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/PurchaseOrders/", `
	{
		"order_number": "",
		"order_date": "2021-04-04 00:00:00",
		"tracking_code": "trackingCode3",
		"buyer_id": 1,
		"order_status_id": 1,
		"carrier_id": 1,
		"warehouse_id": 1
	}
	`)
	r.ServeHTTP(rr, req)
	fmt.Println(req.Body)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
}

func TestStorePurchaseOrderFail(t *testing.T) {
	// crear el Server y definir las Rutas
	r := CreateServer()
	// crear Request del tipo GET y Response para obtener el resultado
	reqMissingData, errMissingData := CreateRequestTest(http.MethodPost, "/api/v1/PurchaseOrders/",
		`{"order_number": "",
	"order_date": "",
	"tracking_code": "",
	"buyer_id": 1,
	"order_status_id": 1,
	"carrier_id": 1,
	"warehouse_id": 1}`)

	// indicar al servidor que pueda atender la solicitud
	//Solicitud Datos faltantes
	r.ServeHTTP(errMissingData, reqMissingData)
	assert.Equal(t, 422, errMissingData.Code)
}

func TestStorePurchaseOrderJsonConflict(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/PurchaseOrders/", `
	{
		"order_number": "order#3",
		"order_date": "2021-04-04 00:00:00",
		"tracking_code": "trackingCode3",
		"buyer_id": 1,
		"order_status_id": 123,
		"carrier_id": 1,
		"warehouse_id": 1
	}
	`)
	r.ServeHTTP(rr, req)
	fmt.Println(req.Body)
	assert.Equal(t, http.StatusConflict, rr.Code)
}

func TestGetAllPurchaseOrders(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/PurchaseOrders/", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
func TestGetPurchaseOrderByIdExistent(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/PurchaseOrders/1", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetPurchaseOrderByIdNonExistent(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/PurchaseOrders/99999", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}
