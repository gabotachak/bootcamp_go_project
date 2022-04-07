package handler

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateOkBuyer(t *testing.T) {
	// crear el Server y definir las Rutas
	r := CreateServer()
	// crear Request del tipo GET y Response para obtener el resultado
	reqCreateOk, errCreateOk := CreateRequestTest(http.MethodPost, "/api/v1/buyers/",
		`{"id":2,
		"card_number_id":"2",
		"first_name":"Nombre",
		"last_name":"Apellido"}`)

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(errCreateOk, reqCreateOk)
	assert.Equal(t, 201, errCreateOk.Code)
}

func TestCreateFailBuyer(t *testing.T) {
	// crear el Server y definir las Rutas
	r := CreateServer()
	// crear Request del tipo GET y Response para obtener el resultado
	reqMissingData, errMissingData := CreateRequestTest(http.MethodPost, "/api/v1/buyers/",
		`{"id":1,
		"card_number_id":"",
		"first_name":"",
		"last_name":""}`)

	reqNotFound, errNotFound := CreateRequestTest(http.MethodPost, "/api/v1/buyers/",
		`{"id":1,
		"card_number_id":"-1",
		"first_name":"Julian",
		"last_name":"Velandia"}`)

	// indicar al servidor que pueda atender la solicitud
	//Solicitud CardNumberId No válido
	r.ServeHTTP(errNotFound, reqNotFound)
	assert.Equal(t, 400, errNotFound.Code)

	//Solicitud Datos faltantes
	r.ServeHTTP(errMissingData, reqMissingData)
	assert.Equal(t, 422, errMissingData.Code)
}

func TestCreateConflictBuyer(t *testing.T) {
	r := CreateServer()
	reqExistingBuyer, errExistingBuyer := CreateRequestTest(http.MethodPost, "/api/v1/buyers/",
		`{"id":1,
		"card_number_id":"1",
		"first_name":"No",
		"last_name":"No"}`)
	r.ServeHTTP(errExistingBuyer, reqExistingBuyer)
	assert.Equal(t, 409, errExistingBuyer.Code)
}

func TestGetOkBuyer(t *testing.T) {
	r := CreateServer()

	reqGetOk, errGetOk := CreateRequestTest(http.MethodGet, "/api/v1/buyers/", "")
	r.ServeHTTP(errGetOk, reqGetOk)
	assert.Equal(t, 200, errGetOk.Code)
}

func TestGetOneOkBuyer(t *testing.T) {
	r := CreateServer()

	reqGetOneOk, errGetOneOk := CreateRequestTest(http.MethodGet, "/api/v1/buyers/1", "")
	r.ServeHTTP(errGetOneOk, reqGetOneOk)
	assert.Equal(t, 200, errGetOneOk.Code)
}

func TestGetOneInexistentBuyer(t *testing.T) {
	r := CreateServer()

	reqGetOneInexistent, errGetOneInexistent := CreateRequestTest(http.MethodGet, "/api/v1/buyers/100", "")
	r.ServeHTTP(errGetOneInexistent, reqGetOneInexistent)
	assert.Equal(t, 404, errGetOneInexistent.Code)
}

func TestUpdateOkBuyer(t *testing.T) {
	r := CreateServer()
	reqUpdateOk, errUpdateOk := CreateRequestTest(http.MethodPatch, "/api/v1/buyers/1",
		`{"id":1,
		"card_number_id":"1",
		"first_name":"Actualizado",
		"last_name":"Actualizado"}`)
	r.ServeHTTP(errUpdateOk, reqUpdateOk)
	assert.Equal(t, 200, errUpdateOk.Code)
}

func TestUpdateNonExistentBuyer(t *testing.T) {
	r := CreateServer()
	reqUpdateNonExistent, errUpdateNonExistent := CreateRequestTest(http.MethodPatch, "/api/v1/buyers/100",
		`{"id":1,
		"card_number_id":"100",
		"first_name":"Julian",
		"last_name":"Velandia"}`)

	reqMissingData, errMissingData := CreateRequestTest(http.MethodPatch, "/api/v1/buyers/100",
		`{"id":1,
		"card_number_id":"",
		"first_name":"",
		"last_name":""}`)

	reqBadRequest, errBadRequest := CreateRequestTest(http.MethodPatch, "/api/v1/buyers/100",
		`{"id":1,
		"card_number_id":"-1",
		"first_name":"Julian",
		"last_name":"Velandia"}`)

	// indicar al servidor que pueda atender la solicitud
	//Solicitud CardNumberId No válido
	r.ServeHTTP(errMissingData, reqMissingData)
	assert.Equal(t, 422, errMissingData.Code)

	//Solicitud Buyer no existe
	r.ServeHTTP(errUpdateNonExistent, reqUpdateNonExistent)
	assert.Equal(t, 404, errUpdateNonExistent.Code)

	//Solicitud Petición erronea
	r.ServeHTTP(errBadRequest, reqBadRequest)
	assert.Equal(t, 400, errBadRequest.Code)
}

func TestDeleteOKBuyer(t *testing.T) {
	r := CreateServer()
	reqDeleteOK, errDeleteOK := CreateRequestTest(http.MethodDelete, "/api/v1/buyers/2", "")
	r.ServeHTTP(errDeleteOK, reqDeleteOK)
	assert.Equal(t, 204, errDeleteOK.Code)
}

func TestDeleteInexistentBuyer(t *testing.T) {
	r := CreateServer()
	reqInexistent, errInexistent := CreateRequestTest(http.MethodDelete, "/api/v1/buyers/100", "")
	r.ServeHTTP(errInexistent, reqInexistent)
	assert.Equal(t, 404, errInexistent.Code)
}

func TestGetPurchaseOrders(t *testing.T) {
	r := CreateServer()
	reqGetPurchaseOrders, errGetPurchaseOrders := CreateRequestTest(http.MethodGet, "/api/v1/buyers/reportPurchaseOrders/", "")

	r.ServeHTTP(errGetPurchaseOrders, reqGetPurchaseOrders)
	assert.Equal(t, 200, errGetPurchaseOrders.Code)
}

func TestGetPurchaseOrdersByBuyerOk(t *testing.T) {
	r := CreateServer()
	reqGetPurchaseOrdersByBuyerOk, errGetPurchaseOrdersByBuyerOk := CreateRequestTest(http.MethodGet, "/api/v1/buyers/reportPurchaseOrders/1", "")

	r.ServeHTTP(errGetPurchaseOrdersByBuyerOk, reqGetPurchaseOrdersByBuyerOk)
	assert.Equal(t, 200, errGetPurchaseOrdersByBuyerOk.Code)
}

func TestGetPurchaseOrdersByBuyerFail(t *testing.T) {
	r := CreateServer()
	reqGetPurchaseOrdersByBuyerFail, errGetPurchaseOrdersByBuyerFail := CreateRequestTest(http.MethodGet, "/api/v1/buyers/reportPurchaseOrders/99999", "")

	r.ServeHTTP(errGetPurchaseOrdersByBuyerFail, reqGetPurchaseOrdersByBuyerFail)
	assert.Equal(t, 400, errGetPurchaseOrdersByBuyerFail.Code)
}
