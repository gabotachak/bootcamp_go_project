package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/productRecord"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProductRecord struct {
	productRecordService productRecord.Service
}

func NewProductRecord(s productRecord.Service) *ProductRecord {
	return &ProductRecord{
		productRecordService: s,
	}
}

// List Product Record godoc
// @Summary Return product records of all Products
// @Tags Products
// @Description This function returns all product records for all products
// @Accept json
// @Produce json
// @Success 200 {object} web.Response
// @Router /api/v1/productsRecords [get]
func (p *ProductRecord) GetAll() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var productos []productRecord.ProductRecordByProduct

		productos, err := p.productRecordService.GetAll(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		if len(productos) == 0 {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "No hay product records almacenados"))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, productos, ""))

	}
}

// List Products Record godoc
// @Summary Return information of a product record record
// @Tags Products
// @Description This function returns a product record serched by id
// @Accept json
// @Produce json
// @Param id path int true "id product record"
// @Success 200 {object} web.Response
// @Router /api/v1/productsRecords/{id} [get]
func (p *ProductRecord) Get() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		id := ctx.Param("id")
		prodId, _ := strconv.Atoi(id)

		productRec, err := p.productRecordService.Get(ctx, prodId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return

		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, productRec, ""))
	}
}

// List Product Record godoc
// @Summary Save a product record
// @Tags Products
// @Description This function stores a product record with an existing product id
// @Accept json
// @Produce json
// @Success 201 {object} web.Response
// @Router /api/v1/productsRecords [post]
func (p *ProductRecord) Save() gin.HandlerFunc {
	type request struct {
		ID             int     `json:"id"`
		LastUpdateDate string  `json:"last_update_date"`
		PurchasePrice  float32 `json:"purchase_price"`
		SalePrice      float32 `json:"sale_price"`
		ProductId      int     `json:"product_id"`
	}

	return func(ctx *gin.Context) {
		var req request
		var errorsBind = []string{}
		var productRecSave domain.ProductRecord
		va := validator.New()

		err3 := ctx.ShouldBindJSON(&req)
		if err3 != nil {
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, err3.Error()))
			return
		}
		errBind := va.Struct(req)

		if errBind != nil {
			for _, err := range errBind.(validator.ValidationErrors) {
				mensaje := fmt.Sprintf(" %s es requerido o dato incorrecto %s", err.Field(), err.Tag())
				errorsBind = append(errorsBind, mensaje)
			}
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, strings.Join(errorsBind, ",")))
			return
		}
		fmt.Println(req)
		productRecSave.LastUpdateDate = req.LastUpdateDate
		productRecSave.ProductId = req.ProductId
		productRecSave.PurchasePrice = req.PurchasePrice
		productRecSave.SalePrice = req.SalePrice
		fmt.Println(productRecSave)
		prodRec, err2 := p.productRecordService.Save(ctx, productRecSave)
		if err2 != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err2.Error()))
			return
		} else {
			//ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, fmt.Sprintf("El producto ha sido creado correctamente con id: %d", prodRec), ""))
			ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, prodRec, ""))
			return
		}

	}
}
