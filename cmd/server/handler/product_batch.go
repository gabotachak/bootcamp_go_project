package handler

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/productBatch"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const (
	LAYOUT_ISO  = "2006-01-02"
	LAYOUT_HOUR = "2006-01-02 03:04"
)

type ProductBatch struct {
	productBatchService productBatch.Service
}

func NewProductBatch(pb productBatch.Service) *ProductBatch {
	return &ProductBatch{
		productBatchService: pb,
	}
}

// List Product Batches godoc
// @Summary Store section
// @Tags Product Batches
// @Description Post product batches
// @Accept json
// @Produce json
// @Success 201 {object} web.Response
// @Router /api/v1/productBatches [post]
func (s *ProductBatch) Store() gin.HandlerFunc {

	type request struct {
		BatchNumber        string  `json:"batch_number" binding:"required" validate:"alphanum"`
		CurrentQuantity    int     `json:"current_quantity" binding:"required" validate:"numeric"`
		CurrentTemperature float64 `json:"current_temperature" binding:"required" validate:"numeric"`
		DueDate            string  `json:"due_date" binding:"required" validate:"ascii"`
		InitialQuantity    int     `json:"initial_quantity" binding:"required" validate:"numeric"`
		ManufacturingDate  string  `json:"manufacturing_date" binding:"required" validate:"ascii"`
		ManufacturingHour  string  `json:"manufacturing_hour" binding:"required" validate:"ascii"`
		MinimumTemperature float64 `json:"minimum_temperature" binding:"required" validate:"numeric"`
		ProductId          int     `json:"product_id" binding:"required" validate:"numeric"`
		SectionId          int     `json:"section_id" binding:"required" validate:"numeric"`
	}

	return func(ctx *gin.Context) {

		var productBatchStore domain.ProductBatch
		var req request
		var errorsBind = []string{}

		if err := ctx.ShouldBindJSON(&req); err != nil {
			fmt.Print(err)
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

		productBatchStore.BatchNumber = req.BatchNumber
		productBatchStore.CurrentQuantity = req.CurrentQuantity
		productBatchStore.CurrentTemperature = req.CurrentTemperature
		productBatchStore.DueDate, _ = time.Parse(LAYOUT_ISO, req.DueDate)
		productBatchStore.InitialQuantity = req.InitialQuantity
		productBatchStore.ManufacturingDate, _ = time.Parse(LAYOUT_ISO, req.ManufacturingDate)
		productBatchStore.ManufacturingHour, _ = time.Parse(LAYOUT_HOUR, req.ManufacturingHour)
		productBatchStore.MinimumTemperature = req.MinimumTemperature
		productBatchStore.ProductId = req.ProductId
		productBatchStore.SectionId = req.SectionId

		id, err := s.productBatchService.Save(ctx, productBatchStore)
		if err != nil {
			// Error from internal
			ctx.JSON(http.StatusConflict, web.NewResponse(http.StatusConflict, nil, err.Error()))
			return
		}

		// Success response
		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, id, "Success"))
	}
}

// List Product Batches godoc
// @Summary List all product batches reports
// @Tags Product Batches
// @Description Get all product batches reports
// @Produce json
// @Success 200 {object} web.Response
// @Router /api/v1/reportProducts [get]
func (s *ProductBatch) ReportAll() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		reports, err := s.productBatchService.ReportAll(ctx)
		if err != nil {

			// Error from internal
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		if len(reports) == 0 {

			// No content
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "There is no reports on database"))
		}

		// Success response
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, reports, ""))
	}
}

func (s *ProductBatch) ReportBySection() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var filter domain.ReportProductBatch
		if ctx.Bind(&filter) != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "bad request"))
			return
		}
		sId := filter.SectionId
		if sId < 0 {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "bad request, only positive integers"))
			return
		}
		report, err := s.productBatchService.ReportBySection(ctx, sId)

		if err != nil {

			// Error from internal
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		// Success response
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, report, ""))
	}
}
