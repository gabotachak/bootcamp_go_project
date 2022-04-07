package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/product"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Product struct {
	productService product.Service
}

func NewProduct(s product.Service) *Product {
	return &Product{
		productService: s,
	}
}

//
// List Products godoc
// @Summary Return all products
// @Tags Products
// @Description This function returns all product
// @Accept json
// @Produce json
// @Success 200 {object} web.Response
// @Router /api/v1/products [get]
func (p *Product) GetAll() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var productos []domain.Product

		productos, err := p.productService.GetAll(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		if len(productos) == 0 {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "No hay productos almacenados"))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, productos, ""))

	}
}

// List Products godoc
// @Summary Return a product
// @Tags Products
// @Description This function returns a product serched by id
// @Accept json
// @Produce json
// @Param id path int true "id product"
// @Success 200 {object} web.Response
// @Router /api/v1/products/{id} [get]
func (p *Product) Get() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		id := ctx.Param("id")
		prodId, _ := strconv.Atoi(id)

		product, err := p.productService.Get(ctx, prodId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return

		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, product, ""))
	}
}

// List Products godoc
// @Summary Save a product
// @Tags Products
// @Description This function stores a product with a productCode that doesnt exist
// @Accept json
// @Produce json
// @Success 201 {object} web.Response
// @Router /api/v1/products [post]
func (p *Product) Save() gin.HandlerFunc {
	type request struct {
		Description    string  `json:"description" binding:"required" validate:"alphanum"`
		ExpirationRate int     `json:"expiration_rate" binding:"required" validate:"numeric"`
		FreezingRate   int     `json:"freezing_rate" binding:"required" validate:"numeric"`
		Height         float32 `json:"height" binding:"required" validate:"numeric"`
		Length         float32 `json:"length" binding:"required" validate:"numeric"`
		Netweight      float32 `json:"net_weight" binding:"required" validate:"numeric"`
		ProductCode    string  `json:"product_code" binding:"required" validate:"alphanum"`
		RecomFreezTemp float32 `json:"recommended_freezing_temperature" binding:"required" validate:"numeric"`
		Width          float32 `json:"width" binding:"required" validate:"numeric"`
		ProductTypeID  int     `json:"product_type_id" binding:"required" validate:"numeric"`
		SellerID       int     `json:"seller_id" binding:"required" validate:"numeric"`
	}

	return func(ctx *gin.Context) {
		var req request
		var errorsBind = []string{}
		var productSave domain.Product
		va := validator.New()

		err3 := ctx.ShouldBindJSON(&req)
		if err3 != nil {
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, "Algun dato fue ingresado de forma incorrecta"))
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

		productSave.Description = req.Description
		productSave.ExpirationRate = req.ExpirationRate
		productSave.FreezingRate = req.FreezingRate
		productSave.Height = req.Height
		productSave.Length = req.Length
		productSave.Netweight = req.Netweight
		productSave.ProductCode = req.ProductCode
		productSave.RecomFreezTemp = req.RecomFreezTemp
		productSave.Width = req.Width
		productSave.ProductTypeID = req.ProductTypeID
		productSave.SellerID = req.SellerID

		prodId, err2 := p.productService.Save(ctx, productSave)

		if prodId == -1 {
			ctx.JSON(http.StatusConflict, web.NewResponse(http.StatusConflict, nil, "El código de producto ya existe"))
		}
		if err2 != nil {
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, err2.Error()))
			return
		}
		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, fmt.Sprintf("El producto ha sido creado correctamente con id: %d", prodId), ""))

	}
}

// List Products godoc
// @Summary Update a product
// @Tags Products
// @Description This function updates a product searched by id
// @Accept json
// @Produce json
// @Param id path int true "id product"
// @Success 200 {object} web.Response
// @Router /api/v1/products/{id} [patch]
func (p *Product) Update() gin.HandlerFunc {

	type request struct {
		Description    string  `json:"description" binding:"required" validate:"alphanum"`
		ExpirationRate int     `json:"expiration_rate" binding:"required" validate:"numeric"`
		FreezingRate   int     `json:"freezing_rate" binding:"required" validate:"numeric"`
		Height         float32 `json:"height" binding:"required" validate:"numeric"`
		Length         float32 `json:"length" binding:"required" validate:"numeric"`
		Netweight      float32 `json:"net_weight" binding:"required" validate:"numeric"`
		ProductCode    string  `json:"product_code" binding:"required" validate:"alphanum"`
		RecomFreezTemp float32 `json:"recommended_freezing_temperature" binding:"required" validate:"numeric"`
		Width          float32 `json:"width" binding:"required" validate:"numeric"`
		ProductTypeID  int     `json:"product_type_id" binding:"required" validate:"numeric"`
		SellerID       int     `json:"seller_id" binding:"required" validate:"numeric"`
	}

	return func(ctx *gin.Context) {

		var productSave domain.Product
		var req request
		var errorsBind = []string{}

		e := ctx.ShouldBindJSON(&req)
		if e != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "algun elemento del body fue mal ingresado"))
			return
		}
		v := validator.New()
		errBind := v.Struct(req)
		if errBind != nil {
			for _, err := range errBind.(validator.ValidationErrors) {
				msg := fmt.Sprintf("%s es requerido o  dato incorrecto %s", err.Field(), err.Tag())
				errorsBind = append(errorsBind, msg)
			}
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, strings.Join(errorsBind, ",")))
			return
		}

		productSave.Description = req.Description
		productSave.ExpirationRate = req.ExpirationRate
		productSave.FreezingRate = req.FreezingRate
		productSave.Height = req.Height
		productSave.Length = req.Length
		productSave.Netweight = req.Netweight
		productSave.ProductCode = req.ProductCode
		productSave.RecomFreezTemp = req.RecomFreezTemp
		productSave.Width = req.Width
		productSave.ProductTypeID = req.ProductTypeID
		productSave.SellerID = req.SellerID

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "Invalid ID"))
			return
		}

		productSave.ID = int(id)

		// prod, err := p.productService.Get(ctx, productSave.ID)
		// if err != nil {
		// 	ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusNotFound, nil, "el elemento es inexistente"))
		// 	return
		// }
		// if prod.ProductCode != productSave.ProductCode {
		// 	ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "ya existe un elemento con ese product code"))
		// 	return
		// }

		err1 := p.productService.Update(ctx, productSave)

		if err1 != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err1.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusCreated, err1, ""))
	}
}

// List Products godoc
// @Summary Delete a product
// @Tags Products
// @Description This function returns a product searched by id
// @Accept json
// @Param id path int true "id product"
// @Produce json
// @Success 204 {object} web.Response
// @Router /api/v1/products/{id} [delete]
func (p *Product) Delete() gin.HandlerFunc {

	return func(c *gin.Context) {

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "invalid ID"))
			return
		}
		err = p.productService.Delete(c, int(id))
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		c.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, fmt.Sprintf("El producto %d ha sido eliminado", id), ""))
	}
}
