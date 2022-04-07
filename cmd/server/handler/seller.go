package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/seller"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/pkg/web"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Seller struct {
	sellerService seller.Service
}

func NewSeller(s seller.Service) *Seller {
	return &Seller{
		sellerService: s,
	}
}

// ListSellers godoc
// @Summary List sellers
// @Tags Sellers
// @Description Get all sellers
// @Produce json
// @Success 200 {object} web.Response
// @Router /api/v1/sellers [get]
func (s *Seller) GetAll() gin.HandlerFunc {

	type response struct {
		Data []domain.Seller `json:"data"`
	}

	return func(c *gin.Context) {
		sel, err := s.sellerService.GetAll(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewError(http.StatusBadRequest, err.Error()))
			return
		}
		if len(sel) == 0 {
			c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "No hay vendedores almacenados"))
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, sel, ""))
	}
}

// List Seller godoc
// @Summary Return a seller
// @Tags Sellers
// @Description Return a seller
// @Produce json
// @Param id path int true "id seller"
// @Success 200 {object} web.Response
// @Router /api/v1/sellers/{id} [get]
func (s *Seller) Get() gin.HandlerFunc {

	return func(c *gin.Context) {
		cid, _ := strconv.Atoi(c.Param("id"))
		sel, err := s.sellerService.Get(c, cid)

		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "error: id no existe"))
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, sel, ""))
	}
}

// List Sellers godoc
// @Summary Store Seller
// @Tags Sellers
// @Description Store seller
// @Accept json
// @Produce json
// @Success 201 {object} web.Response
// @Router /api/v1/sellers [post]
func (s *Seller) Store() gin.HandlerFunc {
	type request struct {
		CID         int    `json:"cid" binding:"required" validate:"numeric"`
		CompanyName string `json:"company_name" binding:"required" validate:"ascii"`
		Address     string `json:"address" binding:"required" validate:"ascii"`
		Telephone   string `json:"telephone" binding:"required" validate:"alphanum"`
		LocalityID  int    `json:"locality_id" binding:"required" validate:"numeric"`
	}

	return func(c *gin.Context) {
		var seller domain.Seller
		var sellerReq request
		var errorsBind = []string{}
		v := validator.New()
		err3 := c.ShouldBindJSON(&sellerReq)
		if err3 != nil {
			c.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, "error: hacen falta datos / tipo de dato"))
			return
		}
		errBind := v.Struct(sellerReq)

		if errBind != nil {
			for _, err := range errBind.(validator.ValidationErrors) {
				mensaje := fmt.Sprintf("error: el campo %s es requerido o su tipo de dato no es el correcto (%s).", err.Field(), err.Tag())
				errorsBind = append(errorsBind, mensaje)
			}
			c.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, strings.Join(errorsBind, ",")))
			return
		}

		seller.CID = sellerReq.CID
		seller.CompanyName = sellerReq.CompanyName
		seller.Address = sellerReq.Address
		seller.Telephone = sellerReq.Telephone
		seller.LocalityID = sellerReq.LocalityID

		id, err1 := s.sellerService.Save(c, seller)
		fmt.Println(id)
		fmt.Println(err1)
		if id == -1 {
			c.JSON(http.StatusConflict, web.NewResponse(http.StatusConflict, nil, err1.Error()))
			return
		}
		if err1 != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "error al almacenar"))
			return
		}
		c.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, fmt.Sprintf("Seller creado con id: %d", id), ""))
	}
}

// List Sellers godoc
// @Summary Update a seller
// @Tags Sellers
// @Description Update seller
// @Accept json
// @Produce json
// @Param id path int true "id seller"
// @Success 200 {object} web.Response
// @Router /api/v1/sellers/{id} [patch]
func (s *Seller) Update() gin.HandlerFunc {

	type request struct {
		CID         int    `json:"cid" binding:"required" validate:"numeric"`
		CompanyName string `json:"company_name" binding:"required" validate:"ascii"`
		Address     string `json:"address" binding:"required" validate:"ascii"`
		Telephone   string `json:"telephone" binding:"required" validate:"alphanum"`
		LocalityID  int    `json:"locality_id" binding:"required" validate:"numeric"`
	}

	return func(c *gin.Context) {
		var seller domain.Seller
		var sellerReq request
		var errorsBind = []string{}

		v := validator.New()
		err3 := c.ShouldBindJSON(&sellerReq)

		if err3 != nil {
			c.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, "error: hacen falta datos / tipo de dato"))
			return
		}
		errBind := v.Struct(sellerReq)

		if errBind != nil {
			for _, err := range errBind.(validator.ValidationErrors) {
				mensaje := fmt.Sprintf("error: el campo %s es requerido o su tipo de dato no es el correcto (%s).", err.Field(), err.Tag())
				errorsBind = append(errorsBind, mensaje)
			}
			c.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, strings.Join(errorsBind, ",")))
			return
		}

		seller.CID = sellerReq.CID
		seller.CompanyName = sellerReq.CompanyName
		seller.Address = sellerReq.Address
		seller.Telephone = sellerReq.Telephone
		seller.LocalityID = sellerReq.LocalityID

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Invalid ID"))
			return
		}

		seller.ID = int(id)
		err1 := s.sellerService.Update(c, seller)
		if err1 != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err1.Error()))
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, err1, ""))
	}

}

// List Warehouse godoc
// @Summary Delete a seller
// @Tags Sellers
// @Description Delete seller
// @Accept json
// @Produce json
// @Param id path int true "id seller"
// @Success 204 {object} web.Response
// @Router /api/v1/sellers/{id} [delete]
func (s *Seller) Delete() gin.HandlerFunc {

	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "Invalid ID"))
			return
		}
		err1 := s.sellerService.Delete(c, id)
		if err1 != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err1.Error()))
			return
		}
		c.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, err1, "seller deleted"))
	}
}
