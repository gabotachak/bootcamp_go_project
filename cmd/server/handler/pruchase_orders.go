package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/purchaseOrders"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/pkg/web"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PurchaseOrder struct {
	purchaseOrderService purchaseOrders.Service
}

func NewPurchaseOrder(c purchaseOrders.Service) *PurchaseOrder {
	return &PurchaseOrder{
		purchaseOrderService: c,
	}
}

// @Summary Get a purchaseOrders
// @Tags purchaseOrders
// @Description Get a report on the purchaseOrders of buyer
// @Param id path int true "Buyer_id"
// @Produce json
// @Success 200 {object} web.Response
// @Router /api/v1/PurchaseOrders/{id} [get]
func (p *PurchaseOrder) Get() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		order, _ := strconv.Atoi(id)

		purchaseOrders, err := p.purchaseOrderService.Get(ctx, order)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "no existen ordenes de compra con esa identificación"))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, purchaseOrders, ""))
	}
}

// List PurchaseOrders godoc
// @Summary List purchaseOrders
// @Tags PurchaseOrders
// @Description Get all PurchaseOrders
// @Produce json
// @Success 200 {object} web.Response
// @Router /api/v1/PurchaseOrders [get]
func (p *PurchaseOrder) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		PurchaseOrders, err := p.purchaseOrderService.GetAll(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		} else {
			ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, PurchaseOrders, ""))
			return
		}
	}
}

// List purchaseOrderService godoc
// @Summary Create purchaseOrderService
// @Tags purchaseOrderServices
// @Description Create purchaseOrderService
// @Accept json
// @Produce json
// @Success 201 {object} web.Response
// @Router /api/v1/purchaseOrder [post]
func (p *PurchaseOrder) Store() gin.HandlerFunc {
	type request struct {
		OrderNumber   string `json:"order_number" binding:"required" validate:"ascii"`
		OrderDate     string `json:"order_date" binding:"required" validate:"ascii"`
		TrackingCode  string `json:"tracking_code" binding:"required" validate:"ascii"`
		BuyerId       int    `json:"buyer_id" binding:"required" validate:"numeric"`
		OrderStatusId int    `json:"order_status_id" binding:"required" validate:"numeric"`
		CarrierId     int    `json:"carrier_id" validate:"numeric"`
		WarehouseId   int    `json:"warehouse_id" validate:"numeric"`
	}
	return func(ctx *gin.Context) {
		var purchaseOrdersStore domain.PurchaseOrder
		var req request
		var errorsBind = []string{}

		ess := ctx.ShouldBindJSON(&req)
		if ess != nil {
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, "algun elemento del body no existe o fue mal ingresado"))
			return
		}

		v := validator.New()
		errBind := v.Struct(req)
		if errBind != nil {
			for _, err := range errBind.(validator.ValidationErrors) {
				msg := fmt.Sprintf("el campo %s es requerido o su tipo de dato no es el correcto %s", err.Field(), err.Tag())
				errorsBind = append(errorsBind, msg)
			}
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, strings.Join(errorsBind, ",")))
			return
		}

		purchaseOrdersStore.OrderNumber = req.OrderNumber
		purchaseOrdersStore.OrderDate = req.OrderDate
		purchaseOrdersStore.TrackingCode = req.TrackingCode
		purchaseOrdersStore.BuyerId = req.BuyerId
		purchaseOrdersStore.OrderStatusId = req.OrderStatusId
		purchaseOrdersStore.CarrierId = req.CarrierId
		purchaseOrdersStore.WarehouseId = req.WarehouseId

		order, err := p.purchaseOrderService.Store(ctx, purchaseOrdersStore)

		if err != nil {
			ctx.JSON(http.StatusConflict, web.NewResponse(http.StatusConflict, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, order, ""))
	}
}
