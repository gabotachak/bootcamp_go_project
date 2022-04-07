package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/locality"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/pkg/web"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

type Locality struct {
	localityService locality.Service
}

func NewLocality(s locality.Service) *Locality {
	return &Locality{
		localityService: s,
	}
}

// List Locality godoc
// @Summary Get Seller Report By Id
// @Tags Localities
// @Description Get Seller Report by Id and locality
// @Produce json
// @Param id path int true "id locality"
// @Success 200 {object} web.Response
// @Router /api/v1/localities/reportSellers{id} [get]
func (s *Locality) GetReport() gin.HandlerFunc {

	return func(c *gin.Context) {
		cid, _ := strconv.Atoi(c.Param("id"))
		sel, err := s.localityService.GetReport(c, cid)

		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "error: id no existe"))
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, sel, ""))
	}
}

// List Locality godoc
// @Summary Get Sellers Report
// @Tags Localities
// @Description Get Seller Report
// @Produce json
// @Success 200 {object} web.Response
// @Router /api/v1/localities/reportSellers [get]
func (s *Locality) GetGeneralReport() gin.HandlerFunc {

	return func(c *gin.Context) {
		sel, err := s.localityService.GetGeneralReport(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewError(http.StatusBadRequest, err.Error()))
			return
		}
		if len(sel) == 0 {
			c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "Reporte vacio"))
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, sel, ""))
	}
}

// List Localities godoc
// @Summary Store Locality
// @Tags Localities
// @Description Store locality
// @Accept json
// @Produce json
// @Success 201 {object} web.Response
// @Router /api/v1/localities [post]
func (s *Locality) Store() gin.HandlerFunc {
	type request struct {
		LocalityName string `json:"locality_name" binding:"required" validate:"ascii"`
		ProvinceName string `json:"province_name" validate:"ascii"`
		CountryName  string `json:"country_name" binding:"required" validate:"ascii"`
	}

	return func(c *gin.Context) {
		var locality domain.Locality
		var localityReq request
		var errorsBind = []string{}
		v := validator.New()
		err3 := c.ShouldBindJSON(&localityReq)
		if err3 != nil {
			c.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, "error: hacen falta datos / tipo de dato"))
			return
		}
		errBind := v.Struct(localityReq)

		if errBind != nil {
			for _, err := range errBind.(validator.ValidationErrors) {
				mensaje := fmt.Sprintf("error: el campo %s es requerido o su tipo de dato no es el correcto (%s).", err.Field(), err.Tag())
				errorsBind = append(errorsBind, mensaje)
			}
			c.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, strings.Join(errorsBind, ",")))
			return
		}

		locality.LocalityName = localityReq.LocalityName
		locality.ProvinceName = localityReq.ProvinceName
		locality.CountryName = localityReq.CountryName

		loc, err1 := s.localityService.Save(c, locality)

		if err1 != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "error al almacenar"))
			return
		}
		c.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, loc, ""))
	}
}
