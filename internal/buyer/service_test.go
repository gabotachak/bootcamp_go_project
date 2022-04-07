package buyer

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type repository_test struct {
	db   *sql.DB
	flag bool
}

func (r *repository_test) GetAll(ctx context.Context) ([]domain.Buyer, error) {
	r.flag = true
	buyer := []domain.Buyer{
		{
			ID:           1,
			CardNumberID: "12345",
			FirstName:    "Julian",
			LastName:     "Velandia",
		},

		{
			ID:           2,
			CardNumberID: "12345",
			FirstName:    "Julian",
			LastName:     "Velandia",
		},
	}

	return buyer, nil
}

func (r *repository_test) Get(ctx context.Context, cardId string) (domain.Buyer, error) {
	r.flag = true

	if cardId == "1" {
		return domain.Buyer{
			ID:           1,
			CardNumberID: "12345",
			FirstName:    "Julian",
			LastName:     "Velandia"}, nil

	} else {

		return domain.Buyer{}, errors.New("el id no existe en la base de datos")
	}
}

func (r *repository_test) Exists(ctx context.Context, cardId string) bool {
	r.flag = true
	return cardId == "1"
}

func (r *repository_test) Save(ctx context.Context, b domain.Buyer) (int, error) {
	r.flag = true

	if b.FirstName == "" || b.LastName == "" {
		return -1, errors.New("falta enviar un campo o algun tipo de dato no es el correcto")
	}

	if b.CardNumberID == "-1" {
		return -100, errors.New("ya existe un elemento con ese CardNumberID")
	}

	return 1, nil
}

func (r *repository_test) Update(ctx context.Context, b domain.Buyer) error {
	r.flag = true
	if b.CardNumberID == "" || b.FirstName == "" || b.LastName == "" {
		return errors.New("falta enviar un campo o algun tipo de dato no es el correcto")
	}
	if b.CardNumberID == "-1" {
		return errors.New("el id no existe en la base de datos")
	}
	return nil
}

func (r *repository_test) Delete(ctx context.Context, cardId string) error {
	r.flag = true

	if cardId != "1" {
		return errors.New("el id es inexistente")
	}

	return nil
}

func (r *repository_test) GetPurchaseOrders(ctx context.Context) ([]BuyerPurchaseOrders, error) {
	return []BuyerPurchaseOrders{{
		ID:           1,
		CardNumberID: "12345",
		FirstName:    "Julian",
		LastName:     "Velandia"}}, nil
}
func (r *repository_test) GetPurchaseOrdersByBuyer(ctx context.Context, id int) (BuyerPurchaseOrders, error) {
	return BuyerPurchaseOrders{
		ID:           1,
		CardNumberID: "12345",
		FirstName:    "Julian",
		LastName:     "Velandia"}, nil
}

type storeMock struct {
	db_test *sql.DB
}

func (sm *storeMock) openBdTest() {
	var err error
	sm.db_test, err = sql.Open("sqlite3", "../.././meli_test.db")
	if err != nil {
		fmt.Println("error inicializacion base de datos", err)
	}
}

func TestSaveOkService(t *testing.T) {

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	store := storeMock{}
	store.openBdTest()

	var repository = &repository_test{store.db_test, false}
	service := NewService(repository)

	newBuyer := domain.Buyer{
		ID:           1,
		CardNumberID: "12345",
		FirstName:    "Julian",
		LastName:     "Velandia"}

	res, _ := service.Save(c, newBuyer)

	assert.NotEmpty(t, res)
	assert.True(t, repository.flag)
}

func TestSaveExistentService(t *testing.T) {
	errorEsperado := errors.New("ya existe un elemento con ese CardNumberID")

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	store := storeMock{}
	store.openBdTest()

	var repository = &repository_test{store.db_test, false}
	service := NewService(repository)

	newBuyer := domain.Buyer{
		ID:           1,
		CardNumberID: "-1",
		FirstName:    "Julian",
		LastName:     "Velandia"}

	res, err := service.Save(c, newBuyer)

	assert.Equal(t, errorEsperado, err, "deben ser iguales")
	assert.Equal(t, -100, res, "deben ser iguales")
	assert.True(t, repository.flag)
}

func TestGetAllService(t *testing.T) {
	//ID, CardNumberID, FirstName, LastName
	// buyerEsperado := []domain.Buyer{
	// 	{
	// 		ID:           1,
	// 		CardNumberID: "12345",
	// 		FirstName:    "Julian",
	// 		LastName:     "Velandia",
	// 	},

	// 	{
	// 		ID:           2,
	// 		CardNumberID: "12345",
	// 		FirstName:    "Julian",
	// 		LastName:     "Velandia",
	// 	},
	// }

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	store := storeMock{}
	store.openBdTest()

	var repository = &repository_test{store.db_test, false}
	service := NewService(repository)

	res, err := service.GetAll(c)

	assert.NotEmpty(t, res, "deben ser iguales")
	assert.True(t, repository.flag)
	assert.Nil(t, err)

}

func TestGetByIdNonExistentService(t *testing.T) {
	errorEsperado := errors.New("el id no existe en la base de datos")
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	store := storeMock{}

	var repository = &repository_test{store.db_test, false}
	service := NewService(repository)

	_, err := service.Get(c, "CardIdNoValido")

	assert.Equal(t, errorEsperado, err, "deben ser iguales")
	assert.True(t, repository.flag)

}

func TestGetByIdExistentService(t *testing.T) {

	buyerEsperado := domain.Buyer{
		ID:           1,
		CardNumberID: "12345",
		FirstName:    "Julian",
		LastName:     "Velandia"}

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	store := storeMock{}
	store.openBdTest()

	var repository = &repository_test{store.db_test, false}
	service := NewService(repository)

	res, err := service.Get(c, "1")

	assert.Equal(t, buyerEsperado, res, "deben ser iguales")
	assert.True(t, repository.flag)
	assert.Nil(t, err)
}

func TestUpdateExistentService(t *testing.T) {

	newBuyer := domain.Buyer{
		ID:           1,
		CardNumberID: "12345",
		FirstName:    "Julian",
		LastName:     "Velandia",
	}

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	store := storeMock{}
	store.openBdTest()

	var repository = &repository_test{store.db_test, false}
	service := NewService(repository)

	newBuyer.CardNumberID = "1"
	er := service.Update(c, newBuyer)

	assert.Equal(t, nil, er, "deben ser iguales")
	assert.True(t, repository.flag)
}

func TestUpdateNonExistentService(t *testing.T) {
	expectedError := errors.New("el id no existe en la base de datos")
	newBuyer := domain.Buyer{
		ID:           -1,
		CardNumberID: "12345",
		FirstName:    "Julian",
		LastName:     "Velandia",
	}

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	store := storeMock{}
	store.openBdTest()

	var repository = &repository_test{store.db_test, false}
	service := NewService(repository)

	newBuyer.CardNumberID = "-1"
	err := service.Update(c, newBuyer)
	assert.Equal(t, expectedError, err, "deben ser iguales")
	assert.True(t, repository.flag)
}

func TestDeleteNotExistentService(t *testing.T) {
	expectedError := errors.New("el id es inexistente")
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	store := storeMock{}
	store.openBdTest()

	var repository = &repository_test{store.db_test, false}
	service := NewService(repository)

	err := service.Delete(c, "CardIdNoValido")

	assert.Equal(t, expectedError, err, "deben ser iguales")
	assert.True(t, repository.flag)

}

func TestDeleteOkService(t *testing.T) {

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	store := storeMock{}
	store.openBdTest()

	var repository = &repository_test{store.db_test, false}
	service := NewService(repository)

	res := service.Delete(c, "1")

	assert.Equal(t, nil, res, "deben ser iguales")
	assert.True(t, repository.flag)
}

func TestExistService(t *testing.T) {

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	store := storeMock{}
	store.openBdTest()

	var repository = &repository_test{store.db_test, false}
	service := NewService(repository)

	res := service.Exists(c, "1")
	assert.True(t, res)
	assert.True(t, repository.flag)
}

func TestGetPurchaseOrdersService(t *testing.T) {

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	store := storeMock{}
	store.openBdTest()

	var repository = &repository_test{store.db_test, false}
	service := NewService(repository)
	expected := []BuyerPurchaseOrders{{
		ID:           1,
		CardNumberID: "12345",
		FirstName:    "Julian",
		LastName:     "Velandia"}}

	res, err := service.GetPurchaseOrders(c)
	assert.Equal(t, expected, res)
	//assert.True(t, repository.flag)
	assert.Nil(t, err)
}

func TestGetPurchaseOrdersByBuyerService(t *testing.T) {

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	store := storeMock{}
	store.openBdTest()

	var repository = &repository_test{store.db_test, false}
	service := NewService(repository)
	expected := BuyerPurchaseOrders{
		ID:           1,
		CardNumberID: "12345",
		FirstName:    "Julian",
		LastName:     "Velandia"}

	res, err := service.GetPurchaseOrdersByBuyer(c, 1)
	assert.Equal(t, expected, res)
	//assert.True(t, repository.flag)
	assert.Nil(t, err)
}
