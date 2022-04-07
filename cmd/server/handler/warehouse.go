package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/warehouse"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Warehouse struct {
	warehouseService warehouse.Service
}

func NewWarehouse(w warehouse.Service) *Warehouse {
	return &Warehouse{
		warehouseService: w,
	}
}

// List Warehouse godoc
// @Summary Return a warehouse
// @Tags Warehouses
// @Description Return a warehouse
// @Param id path int true "warehouse id"
// @Produce json
// @Success 200 {object} web.Response
// @Router /api/v1/warehouses/{id} [get]
func (w *Warehouse) Get() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		wId, _ := strconv.Atoi(id)

		warehouse, err := w.warehouseService.Get(ctx, wId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "el elemento es inexistente"))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, warehouse, ""))

	}
}

// ListWarehouses godoc
// @Summary List warehouses
// @Tags Warehouses
// @Description Get all warehouses
// @Produce json
// @Success 200 {object} web.Response
// @Router /api/v1/warehouses [get]
func (w *Warehouse) GetAll() gin.HandlerFunc {

	return func(c *gin.Context) {
		var ws []domain.Warehouse
		ws, err := w.warehouseService.GetAll(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		if len(ws) == 0 {
			c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "No hay warehouses almacenadas"))
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, ws, ""))
	}
}

// List Warehouse godoc
// @Summary Create warehouse
// @Tags Warehouses
// @Description Create warehouse
// @Accept json
// @Produce json
// @Success 201 {object} web.Response
// @Router /api/v1/warehouses [post]
func (w *Warehouse) Store() gin.HandlerFunc {
	type request struct {
		Address       string `json:"address" binding:"required" validate:"ascii"`
		Telephone     string `json:"telephone" binding:"required" validate:"alphanum"`
		WarehouseCode string `json:"warehouse_code" binding:"required" validate:"alphanum"`
		LocalityId    int    `json:"locality_id" binding:"required" validate:"numeric"`
	}
	return func(ctx *gin.Context) {
		var dw domain.Warehouse
		var req request
		var errorsBind = []string{}

		e := ctx.ShouldBindJSON(&req)
		if e != nil {
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

		dw.Address = req.Address
		dw.Telephone = req.Telephone
		dw.LocalityId = req.LocalityId
		dw.WarehouseCode = req.WarehouseCode

		ware, err := w.warehouseService.Save(ctx, dw)
		if ware == -1 {
			ctx.JSON(http.StatusConflict, web.NewResponse(http.StatusConflict, nil, err.Error()))
			return
		}
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, ware, ""))
	}

}

// List Warehouse godoc
// @Summary Update a warehouse
// @Tags Warehouses
// @Description Update warehouse
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {object} web.Response
// @Router /api/v1/warehouses/{id} [patch]
func (w *Warehouse) Update() gin.HandlerFunc {
	type request struct {
		Address       string `json:"address" binding:"required" validate:"ascii"`
		Telephone     string `json:"telephone" binding:"required" validate:"alphanum"`
		WarehouseCode string `json:"warehouse_code" binding:"required" validate:"alphanum"`
		LocalityId    int    `json:"locality_id" binding:"required" validate:"numeric"`
	}
	return func(ctx *gin.Context) {

		var dw domain.Warehouse
		var req request
		var errorsBind = []string{}

		e := ctx.ShouldBindJSON(&req)
		if e != nil {
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

		dw.Address = req.Address
		dw.Telephone = req.Telephone
		dw.LocalityId = req.LocalityId
		dw.WarehouseCode = req.WarehouseCode

		id, e := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if e != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Invalid ID"))
			return
		}

		dw.ID = int(id)
		err := w.warehouseService.Update(ctx, dw)

		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, err, ""))
	}
}

// List Warehouse godoc
// @Summary Delete a warehouse
// @Tags Warehouses
// @Description Delete warehouse
// @Accept json
// @Param id path int true "Id"
// @Produce json
// @Success 204 {object} web.Response
// @Router /api/v1/warehouses/{id} [delete]
func (w *Warehouse) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "invalid ID"))
			return
		}
		err = w.warehouseService.Delete(c, int(id))
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "el elemento es inexistente"))
			return
		}
		c.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, nil, "la tienda %d ha sido eliminada"))

	}
}
