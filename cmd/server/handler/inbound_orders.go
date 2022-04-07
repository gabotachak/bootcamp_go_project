package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	inboundOrders "github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/inbound_orders"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type InboundOrders struct {
	inboundOrdersService inboundOrders.Service
}

func NewInboundOrders(io inboundOrders.Service) *InboundOrders {
	return &InboundOrders{
		inboundOrdersService: io,
	}
}

// List Inbound orders godoc
// @Summary Return a Inbound Order
// @Tags InboundOrders
// @Description Return a Inbound orders
// @Produce json
// @Param id path int true "id Inbound order"
// @Success 200 {object} web.Response
// @Router /api/v1/InboundOrders/{id} [get]
func (io *InboundOrders) Get() gin.HandlerFunc {

	return func(c *gin.Context) {
		id := c.Param("id")
		idio, _ := strconv.Atoi(id)
		InboundOrders, errGet := io.inboundOrdersService.Get(c, idio)
		if errGet != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "inbound order no encontrada"))
			return
		} else {
			c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, InboundOrders, ""))
		}
	}
}

// List InboundOrderss godoc
// @Summary List InboundOrderss
// @Tags InboundOrderss
// @Description Get all InboundOrderss
// @Produce json
// @Success 200 {object} web.Response
// @Router /api/v1/InboundOrders [get]
func (e *InboundOrders) GetAll() gin.HandlerFunc {

	return func(c *gin.Context) {

		InboundOrderss, err := e.inboundOrdersService.GetAll(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		} else {
			c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, InboundOrderss, ""))
			return
		}

	}
}

// List InboundOrderss godoc
// @Summary Store InboundOrders
// @Tags InboundOrderss
// @Description Post InboundOrders
// @Accept json
// @Produce json
// @Success 201 {object} web.Response
// @Router /api/v1/InboundOrders [post]
func (io *InboundOrders) Store() gin.HandlerFunc {
	type request struct {
		OrderDate      string `json:"order_date" binding:"required" validate:"ascii"`
		OrderNumber    string `json:"order_number" binding:"required" validate:"ascii"`
		EmployeeID     int    `json:"employee_id" binding:"required" validate:"numeric"`
		ProductBatchID int    `json:"product_batch_id" binding:"required" validate:"numeric"`
		WarehouseID    int    `json:"warehouse_id" binding:"required" validate:"numeric"`
	}

	return func(c *gin.Context) {

		var InboundOrders request
		var InboundOrdersStore domain.Inbound_Orders
		var errorsBind = []string{}
		v := validator.New()
		errJSON := c.ShouldBindJSON(&InboundOrders)
		if errJSON != nil {
			c.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, "Algun dato en el body esta mal ingresado"))
			return
		}
		errBind := v.Struct(InboundOrders)
		if errBind != nil {
			for _, err := range errBind.(validator.ValidationErrors) {
				mensaje := fmt.Sprintf("el campo %s es requerido o su tipo de dato no es el correcto %s", err.Field(), err.Tag())
				errorsBind = append(errorsBind, mensaje)
			}
			c.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, strings.Join(errorsBind, ",")))
			return
		}
		InboundOrdersStore.OrderDate = InboundOrders.OrderDate
		InboundOrdersStore.OrderNumber = InboundOrders.OrderNumber
		InboundOrdersStore.EmployeeID = InboundOrders.EmployeeID
		InboundOrdersStore.WarehouseID = InboundOrders.WarehouseID
		InboundOrdersStore.ProductBatchID = InboundOrders.ProductBatchID
		order, errInboundOrders := io.inboundOrdersService.Save(c, InboundOrdersStore)
		if errInboundOrders != nil {
			c.JSON(http.StatusConflict, web.NewResponse(http.StatusConflict, nil, errInboundOrders.Error()))
			return
		} else {
			c.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, order, ""))
			return
		}

	}
}
