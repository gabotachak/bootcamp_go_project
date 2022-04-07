package handler

import (
	"fmt"
	"net/http"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/stretchr/testify/assert"
)

func TestCreateOkSeller(t *testing.T) {
	r := CreateServer()
	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/sellers/",
		`{
		    "cid": 13100,
		    "company_name": "Digital House",
		    "address": "Buenos Aires, Argentina, 54121",
		    "telephone": "1234567890",
		    "locality_id": 1
		  }`)
	r.ServeHTTP(rr, req)
	assert.Equal(t, 201, rr.Code)
}

func TestCreateFailSeller(t *testing.T) {
	r := CreateServer()
	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/sellers/",
		`{"cid": 6,
		"address": "Buenos Aires, Argentina, 54121",
		"telephone": "1234567890 ",
		"locality_id": 1,}`)
	reqValidator, errValidator := CreateRequestTest(http.MethodPost, "/api/v1/sellers/",
		`{"cid": 111,
		"company_name": "ACTUALIZADA",
		"address": "ACTUALIZADA",
		"telephone": "1234567890 -",
		"locality_id": 1}`)
	r.ServeHTTP(rr, req)
	assert.Equal(t, 422, rr.Code)

	r.ServeHTTP(errValidator, reqValidator)
	assert.Equal(t, 422, errValidator.Code)

}

func TestCreateConflictSeller(t *testing.T) {
	r := CreateServer()
	req, rr := CreateRequestTest(http.MethodPost, "/api/v1/sellers/",
		`{"cid": 5,
		"company_name": "Digital House",
		"address": "Buenos Aires, Argentina, 54121",
		"telephone": "1234567890",
		"locality_id": 1}`)
	r.ServeHTTP(rr, req)
	fmt.Println(rr.Body)
	assert.Equal(t, 409, rr.Code)
}

func TestGetAllOkSeller(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/sellers/", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}

func TestGetOneOkSeller(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/sellers/8", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}

func TestGetOneInexistentSeller(t *testing.T) {
	r := CreateServer()

	req, rr := CreateRequestTest(http.MethodGet, "/api/v1/sellers/101", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 404, rr.Code)
}

func TestUpdateOkSeller(t *testing.T) {
	r := CreateServer()
	req, rr := CreateRequestTest(http.MethodPatch, "/api/v1/sellers/9",
		`{"cid": 1111,
		"company_name": "ACTUALIZADA",
		"address": "ACTUALIZADA",
		"telephone": "1234567890",
		"locality_id": 1}`)
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}

func TestUpdateNonExistentSeller(t *testing.T) {
	r := CreateServer()
	reqNonExist, errNonExist := CreateRequestTest(http.MethodPatch, "/api/v1/sellers/101",
		`{"cid": 111,
		"company_name": "ACTUALIZADA",
		"address": "ACTUALIZADA",
		"telephone": "1234567890",
		"locality_id": 1}`)
	reqMissingData, errMissingData := CreateRequestTest(http.MethodPatch, "/api/v1/sellers/101",
		`{"cid": 111,
		"address": "ACTUALIZADA",
		"telephone": "1234567890",
		"locality_id": 1}`)
	reqBadRequest, errBadRequest := CreateRequestTest(http.MethodPatch, "/api/v1/sellers/hola",
		`{"cid": 111,
		"company_name": "ACTUALIZADA",
		"address": "ACTUALIZADA",
		"telephone": "1234567890",
		"locality_id": 1}`)
	reqValidator, errValidator := CreateRequestTest(http.MethodPatch, "/api/v1/sellers/hola",
		`{"cid": 111,
		"company_name": "ACTUALIZADA",
		"address": "ACTUALIZADA",
		"telephone": "1234567890 -",
		"locality_id": 1}`)
	r.ServeHTTP(errNonExist, reqNonExist)
	assert.Equal(t, 404, errNonExist.Code)

	r.ServeHTTP(errMissingData, reqMissingData)
	assert.Equal(t, 422, errMissingData.Code)

	r.ServeHTTP(errBadRequest, reqBadRequest)
	assert.Equal(t, 400, errBadRequest.Code)

	r.ServeHTTP(errValidator, reqValidator)
	assert.Equal(t, 422, errValidator.Code)
}

func TestDeleteOkSeller(t *testing.T) {
	r := CreateServer()
	req, rr := CreateRequestTest(http.MethodDelete, "/api/v1/sellers/11", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 204, rr.Code)
}

func TestDeleteInexistentSeller(t *testing.T) {
	r := CreateServer()
	req, rr := CreateRequestTest(http.MethodDelete, "/api/v1/sellers/101", "")
	reqInvalid, errInvalid := CreateRequestTest(http.MethodDelete, "/api/v1/sellers/hola", "")

	r.ServeHTTP(rr, req)
	assert.Equal(t, 404, rr.Code)

	r.ServeHTTP(errInvalid, reqInvalid)
	assert.Equal(t, 404, errInvalid.Code)
}
