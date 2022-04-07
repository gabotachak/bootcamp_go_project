package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/section"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/pkg/web"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Section struct {
	sectionService section.Service
}

func NewSection(s section.Service) *Section {
	return &Section{
		sectionService: s,
	}
}

// List Sections godoc
// @Summary Return a section
// @Tags Sections
// @Description Return a section
// @Produce json
// @Param id path int true "id section"
// @Success 200 {object} web.Response
// @Router /api/v1/sections/{id} [get]
func (s *Section) Get() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		sId, _ := strconv.Atoi(id)

		section, err := s.sectionService.Get(ctx, sId)

		if err != nil {

			// Error from internal
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		// Success response
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, section, ""))

	}
}

// List Sections godoc
// @Summary List sections
// @Tags Sections
// @Description Get all sections
// @Produce json
// @Success 200 {object} web.Response
// @Router /api/v1/sections [get]
func (s *Section) GetAll() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		sections, err := s.sectionService.GetAll(ctx)
		if err != nil {

			// Error from internal
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		if len(sections) == 0 {

			// No content
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "No hay Sections en la Base de Datos"))
		}

		// Success response
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, sections, ""))
	}
}

// List Sections godoc
// @Summary Store section
// @Tags Sections
// @Description Post sections
// @Accept json
// @Produce json
// @Success 201 {object} web.Response
// @Router /api/v1/sections [post]
func (s *Section) Store() gin.HandlerFunc {

	type request struct {
		SectionNumber      string  `json:"section_number" binding:"required" validate:"alphanum"`
		CurrentTemperature float64 `json:"current_temperature" binding:"required" validate:"numeric"`
		MinTemperature     float64 `json:"minimum_temperature" binding:"required" validate:"numeric"`
		CurrentCapacity    int     `json:"current_capacity" binding:"required" validate:"numeric"`
		MinCapacity        int     `json:"minimum_capacity" binding:"required" validate:"numeric"`
		MaxCapacity        int     `json:"maximum_capacity" binding:"required" validate:"numeric"`
		WarehouseID        int     `json:"warehouse_id" binding:"required" validate:"numeric"`
		ProductTypeID      int     `json:"product_type_id" binding:"required" validate:"numeric"`
	}

	return func(ctx *gin.Context) {

		var sectionStore domain.Section
		var req request
		var errorsBind = []string{}

		if err := ctx.ShouldBindJSON(&req); err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				msg := fmt.Sprintf("%s camp is required", err.Field())
				errorsBind = append(errorsBind, msg)
			}

			// Error by blank fields
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, strings.Join(errorsBind, ","), ""))
			return
		}

		v := validator.New()
		errBind := v.Struct(req)

		if errBind != nil {
			for _, err := range errBind.(validator.ValidationErrors) {
				msg := fmt.Sprintf("%s camp is required", err.Field())
				errorsBind = append(errorsBind, msg)
			}

			// Error by data type
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, strings.Join(errorsBind, ","), ""))
			return
		}

		sectionStore.SectionNumber = req.SectionNumber
		sectionStore.CurrentTemperature = req.CurrentTemperature
		sectionStore.MinimumTemperature = req.MinTemperature
		sectionStore.CurrentCapacity = req.CurrentCapacity
		sectionStore.MinimumCapacity = req.MinCapacity
		sectionStore.MaximumCapacity = req.MaxCapacity
		sectionStore.WarehouseID = req.WarehouseID
		sectionStore.ProductTypeID = req.ProductTypeID

		id, err := s.sectionService.Save(ctx, sectionStore)

		if id == -1 {

			// Conflict with Section Number
			ctx.JSON(http.StatusConflict, web.NewResponse(http.StatusConflict, nil, err.Error()))
			return
		}

		if err != nil {

			// Error from internal
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		// Success response
		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, id, "Exitoso"))
	}
}

// List Sections godoc
// @Summary Update a section
// @Tags Sections
// @Description Update section
// @Accept json
// @Produce json
// @Param id path int true "Section id"
// @Success 200 {object} web.Response
// @Router /api/v1/sections/{id} [patch]
func (s *Section) Update() gin.HandlerFunc {

	type request struct {
		SectionNumber      string  `json:"section_number" validate:"alphanum"`
		CurrentTemperature float64 `json:"current_temperature" validate:"numeric"`
		MinTemperature     float64 `json:"minimum_temperature" validate:"numeric"`
		CurrentCapacity    int     `json:"current_capacity" validate:"numeric"`
		MinCapacity        int     `json:"minimum_capacity" validate:"numeric"`
		MaxCapacity        int     `json:"maximum_capacity" validate:"numeric"`
		WarehouseID        int     `json:"warehouse_id" validate:"numeric"`
		ProductTypeID      int     `json:"product_type_id" validate:"numeric"`
	}

	return func(ctx *gin.Context) {
		var sectionStore domain.Section
		var req request
		var errorsBind = []string{}

		if err := ctx.ShouldBindJSON(&req); err != nil {

			// Error by blank fields
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, strings.Join(errorsBind, ","), ""))
			return
		}

		v := validator.New()
		errBind := v.Struct(req)

		if errBind != nil {
			for _, err := range errBind.(validator.ValidationErrors) {
				msg := fmt.Sprintf("el campo %s es requerido o su tipo de dato es incorrecto %s", err.Field(), err.Tag())
				errorsBind = append(errorsBind, msg)
			}

			// Error from internal
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, strings.Join(errorsBind, ","), ""))
			return
		}

		sectionStore.SectionNumber = req.SectionNumber
		sectionStore.CurrentTemperature = req.CurrentTemperature
		sectionStore.MinimumTemperature = req.MinTemperature
		sectionStore.CurrentCapacity = req.CurrentCapacity
		sectionStore.MinimumCapacity = req.MinCapacity
		sectionStore.MaximumCapacity = req.MaxCapacity
		sectionStore.WarehouseID = req.WarehouseID
		sectionStore.ProductTypeID = req.ProductTypeID

		fmt.Println("req:", req)

		id, e := strconv.Atoi(ctx.Param("id"))

		if e != nil {

			// id from url is invalid
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Invalid ID"))
			return
		}

		sectionStore.ID = id

		err := s.sectionService.Update(ctx, sectionStore)

		if err != nil {
			// Error from internal
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		// Update sucess
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, id, "Exitoso"))
	}
}

// List Sections godoc
// @Summary Delete a section
// @Tags Sections
// @Description Delete section
// @Accept json
// @Param id path int true "Section id"
// @Produce json
// @Success 204 {object} web.Response
// @Router /api/v1/sections/{id} [delete]
func (s *Section) Delete() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {

			// Invalid id from url
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "ID inválido"))
			return
		}
		err = s.sectionService.Delete(ctx, id)
		if err != nil {

			// Section does not exist
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "Error, ID no encontrado"))
			return
		}

		// Sucess response
		ctx.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, fmt.Sprintf("La sección %d ha sido eliminada correctamente", id), ""))
	}
}
