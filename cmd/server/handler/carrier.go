package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/carrier"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Carrier struct {
	carrierService carrier.Service
}

// Carriers contructor
func NewCarrier(c carrier.Service) *Carrier {
	return &Carrier{
		carrierService: c,
	}
}

// @Summary Return carriers by location with details of each one
// @Tags Carriers by location with details of each one
// @Description Return Carriers by location with details of each one
// @Produce json
// @Success 200 {object} web.Response
// @Router /api/v1/localities/reportCarriers{id}/details [get]
func (c *Carrier) GetReportDetails() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		cId, _ := strconv.Atoi(id)

		carriers, err := c.carrierService.GetReportDetails(ctx, cId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "no existe una localidad con ese id"))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, carriers, ""))
	}
}

// @Summary I get a report of the number of carriers by location
// @Tags number of carriers by location
// @Description I get a report of the number of carriers by location
// @Produce json
// @Success 200 {object} web.Response
// @Router /api/v1/localities/reportCarriers/ [get]
func (c *Carrier) GetAllReport() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var ls []carrier.CarriersByLocality
		ls, err := c.carrierService.GetAllReport(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		if len(ls) == 0 {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "no hay transportistas almacenadas"))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, ls, ""))
	}
}

// @Summary Get a report on the number of carriers in a location
// @Tags Number of carriers
// @Description Get a report on the number of carriers in a location
// @Param id path int true "locality id"
// @Produce json
// @Success 200 {object} web.Response
// @Router /api/v1/localities/reportCarriers{id} [get]
func (c *Carrier) GetReport() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		cId, _ := strconv.Atoi(id)

		carriers, err := c.carrierService.GetReport(ctx, cId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "no existe una localidad con ese id"))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, carriers, ""))
	}
}

// @Summary Create carrier
// @Tags carriers
// @Description Create carrier
// @Accept json
// @Produce json
// @Success 201 {object} web.Response
// @Router /api/v1/carriers [post]
func (c *Carrier) Store() gin.HandlerFunc {
	type request struct {
		CID         string `json:"cid" binding:"required" validate:"ascii"`
		CompanyName string `json:"company_name" binding:"required" validate:"ascii"`
		Address     string `json:"address" binding:"required" validate:"ascii"`
		Telephone   string `json:"telephone" binding:"required" validate:"ascii"`
		LocalityId  int    `json:"locality_id" binding:"required" validate:"numeric"`
	}
	return func(ctx *gin.Context) {
		var dc domain.Carrier
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

		dc.CID = req.CID
		dc.CompanyName = req.CompanyName
		dc.Address = req.Address
		dc.Telephone = req.Telephone
		dc.LocalityId = req.LocalityId

		carr, err := c.carrierService.Store(ctx, dc)

		if err != nil {
			ctx.JSON(http.StatusConflict, web.NewResponse(http.StatusConflict, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, carr, ""))
	}
}
