package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/buyer"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/pkg/web"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Buyer struct {
	buyerService buyer.Service
}

func NewBuyer(b buyer.Service) *Buyer {
	return &Buyer{
		buyerService: b,
	}
}

// ListBuyers godoc
// @Summary List Buyer
// @Tags Buyer
// @Description Get Buyer
// @Accept  json
// @Produce  json
// @Param id path int true "Buyer cardNumberID"
// @Success 200 {object} web.Response
// @Router /api/v1/buyers/{id} [get]
func (b *Buyer) Get() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		cardId := ctx.Param("id")

		buyer, err := b.buyerService.Get(ctx, cardId)
		if err != nil {

			//Error desde internal
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		//Respuesta exitosa
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, buyer, ""))

	}
}

// ListBuyer godoc
// @Summary List Buyers
// @Tags Buyer
// @Description Get Buyers
// @Accept  json
// @Produce  json
// @Param nombre query string false "Buyer name"
// @Success 200 {object} web.Response
// @Router /api/v1/buyers [get]
func (b *Buyer) GetAll() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		buyer, err := b.buyerService.GetAll(ctx)
		if err != nil {

			//Error en internal
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		if len(buyer) == 0 {

			//No hay contenido
			ctx.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, nil, "No hay warehouses almacenadas"))
			return
		}

		//Datos retornados exitosamente
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, buyer, ""))
	}
}

// StoreBuyer godoc
// @Summary Store Buyer
// @Tags Buyer
// @Description Store Buyer
// @Accept json
// @Produce json
// @Success 201 {object} web.Response
// @Router /api/v1/buyers [post]
func (b *Buyer) Store() gin.HandlerFunc {

	type request struct {
		CardNumberID string `json:"card_number_id" binding:"required" validate:"alphanum"`
		FirstName    string `json:"first_name" binding:"required" validate:"alphanum"`
		LastName     string `json:"last_name" binding:"required" validate:"alphanum"`
	}

	return func(ctx *gin.Context) {

		var buyer domain.Buyer
		var req request
		var errorsBind = []string{}

		if err := ctx.ShouldBindJSON(&req); err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				msg := fmt.Sprintf("el campo %s es requerido", err.Field())
				errorsBind = append(errorsBind, msg)
			}

			//Error por campos vacios
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, strings.Join(errorsBind, ",")))
			return
		}

		v := validator.New()
		errBind := v.Struct(req)

		if errBind != nil {
			for _, err := range errBind.(validator.ValidationErrors) {
				msg := fmt.Sprintf("El tipo de dato no es el correcto %s", err.Field())
				errorsBind = append(errorsBind, msg)
			}
			//Error en el tipo de datos
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, strings.Join(errorsBind, ",")))
			return
		}
		buyer.CardNumberID = req.CardNumberID
		buyer.FirstName = req.FirstName
		buyer.LastName = req.LastName

		buyerId, err := b.buyerService.Save(ctx, buyer)

		if buyerId == -100 {

			//Error por el exist de internal
			ctx.JSON(http.StatusConflict, web.NewResponse(http.StatusConflict, nil, err.Error()))
			return
		}
		/*
			if err != nil {

				//Error de internal
				ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
				return
			}
		*/

		//Buyer guardado
		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, fmt.Sprintf("El buyer %v creado exitosamente", buyerId), ""))

	}
}

// UpdateBuyer godoc
// @Summary Update Buyer
// @Tags Buyer
// @Description Update Buyer
// @Accept json
// @Produce json
// @Param id path int true "Buyer ID"
// @Success 200 {object} web.Response
// @Router /api/v1/buyers/{id} [patch]
func (b *Buyer) Update() gin.HandlerFunc {

	type request struct {
		CardNumberID string `json:"card_number_id" binding:"required" validate:"alphanum"`
		FirstName    string `json:"first_name" binding:"required" validate:"alphanum"`
		LastName     string `json:"last_name" binding:"required" validate:"alphanum"`
	}

	return func(ctx *gin.Context) {
		var buyer domain.Buyer
		var req request
		var errorsBind = []string{}

		if err := ctx.ShouldBindJSON(&req); err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				msg := fmt.Sprintf("el campo %s es requerido", err.Field())
				errorsBind = append(errorsBind, msg)
			}

			//Error de internal
			ctx.JSON(http.StatusUnprocessableEntity, web.NewErrorf(http.StatusUnprocessableEntity, strings.Join(errorsBind, ","), nil))
			return
		}

		v := validator.New()
		errBind := v.Struct(req)

		if errBind != nil {
			for _, err := range errBind.(validator.ValidationErrors) {
				msg := fmt.Sprintf("el campo %s es requerido o su tipo de dato no es el correcto %s", err.Field(), err.Tag())
				errorsBind = append(errorsBind, msg)
			}

			//Error de internal
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, strings.Join(errorsBind, ",")))
			return
		}
		buyer.CardNumberID = req.CardNumberID
		buyer.FirstName = req.FirstName
		buyer.LastName = req.LastName

		err := b.buyerService.Update(ctx, buyer)

		if err != nil {

			//Error de internal
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		//Actualización de datos exitosa
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, "Datos parcheados correctamente", ""))

	}
}

// DeleteBuyer godoc
// @Summary Delete Buyer
// @Tags Buyer
// @Description Delete buyer
// @Accept  json
// @Buyer  json
// @Param id path int true "Buyer ID"
// @Success 204 {object} web.Response
// @Router /api/v1/buyers/{id} [delete]
func (b *Buyer) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		cardId := ctx.Param("id")

		err := b.buyerService.Delete(ctx, cardId)
		if err != nil {

			//El buyer no existe
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		//retorno de no contenido
		ctx.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, fmt.Sprintf("El buyer %v ha sido eliminado", cardId), ""))

	}
}

// List Buyers godoc
// @Summary List Inbound Orders by Buyer
// @Tags Buyers
// @Description Get all Buyers with PurchaseOrders count
// @Produce json
// @Success 200 {object} web.Response
// @Router /api/v1/Buyers/reportPurchaseOrders [get]
func (b *Buyer) GetPurchaseOrders() gin.HandlerFunc {

	return func(c *gin.Context) {

		Buyers, err := b.buyerService.GetPurchaseOrders(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		} else {
			c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, Buyers, ""))
			return
		}

	}
}

// List Buyers godoc
// @Summary List Inbound Orders by Buyer
// @Tags Buyers
// @Description Get Buyer with PurchaseOrders count by id
// @Produce json
// @Success 200 {object} web.Response
// @Router /api/v1/Buyers/reportPurchaseOrders/{id} [get]
func (b *Buyer) GetPurchaseOrdersByBuyer() gin.HandlerFunc {

	return func(c *gin.Context) {
		id := c.Param("id")
		idBuyer, _ := strconv.Atoi(id)
		Buyers, err := b.buyerService.GetPurchaseOrdersByBuyer(c, idBuyer)
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		} else {
			c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, Buyers, ""))
			return
		}

	}
}
