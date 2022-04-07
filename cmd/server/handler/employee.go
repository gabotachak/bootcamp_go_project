package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/domain"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/employee"
	web "github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/pkg/web"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Employee struct {
	employeeService employee.Service
}

func NewEmployee(e employee.Service) *Employee {
	return &Employee{
		employeeService: e,
	}
}

// List Employees godoc
// @Summary Return a Employee
// @Tags Employees
// @Description Return a Employee
// @Produce json
// @Param id path int true "Employee cardNumberID"
// @Success 200 {object} web.Response
// @Router /api/v1/employees/{id} [get]
func (e *Employee) Get() gin.HandlerFunc {

	return func(c *gin.Context) {

		employee, errGet := e.employeeService.Get(c, c.Param("id"))
		if errGet != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "El empleado no existe"))
			return
		} else {
			c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, employee, ""))
		}
	}
}

// List Employees godoc
// @Summary List Employees
// @Tags Employees
// @Description Get all employees
// @Produce json
// @Success 200 {object} web.Response
// @Router /api/v1/employees [get]
func (e *Employee) GetAll() gin.HandlerFunc {

	return func(c *gin.Context) {

		employees, err := e.employeeService.GetAll(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		} else {
			c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, employees, ""))
			return
		}

	}
}

// List Employees godoc
// @Summary Store employee
// @Tags Employees
// @Description Post employee
// @Accept json
// @Produce json
// @Success 201 {object} web.Response
// @Router /api/v1/employees [post]
func (e *Employee) Store() gin.HandlerFunc {
	type request struct {
		CardNumberID string `json:"card_number_id" binding:"required" validate:"alphanum"`
		FirstName    string `json:"first_name" binding:"required" validate:"ascii"`
		LastName     string `json:"last_name" binding:"required" validate:"ascii"`
		WarehouseID  int    `json:"warehouse_id" binding:"required" validate:"numeric"`
	}

	return func(c *gin.Context) {

		var employee request
		var employeeStore domain.Employee
		var errorsBind = []string{}
		v := validator.New()
		errJSON := c.ShouldBindJSON(&employee)
		if errJSON != nil {
			c.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, "Algun dato en el body esta mal ingresado"))
			return
		}
		errBind := v.Struct(employee)
		if errBind != nil {
			for _, err := range errBind.(validator.ValidationErrors) {
				mensaje := fmt.Sprintf("el campo %s es requerido o su tipo de dato no es el correcto %s", err.Field(), err.Tag())
				errorsBind = append(errorsBind, mensaje)
			}
			c.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, strings.Join(errorsBind, ",")))
			return
		}
		employeeStore.CardNumberID = employee.CardNumberID
		employeeStore.FirstName = employee.FirstName
		employeeStore.LastName = employee.LastName
		employeeStore.WarehouseID = employee.WarehouseID
		id, errEmployee := e.employeeService.Save(c, employeeStore)
		if errEmployee != nil {
			c.JSON(http.StatusConflict, web.NewResponse(http.StatusConflict, nil, "El Empleado no pudo ser creado porque el card_number_id ya existe o el warehouse_id no existe"))
			return
		} else {
			c.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, fmt.Sprintf("El empleado ha sido creado correctamente con id: %d", id), ""))
			return
		}

	}
}

// List Employees godoc
// @Summary Update a employee
// @Tags Employees
// @Description Update employee
// @Accept json
// @Produce json
// @Success 200 {object} web.Response
// @Param id path int true "CardNumberID"
// @Router /api/v1/employees/{id} [patch]
func (e *Employee) Update() gin.HandlerFunc {
	type request struct {
		FirstName   string `json:"first_name" binding:"required" validate:"ascii"`
		LastName    string `json:"last_name" binding:"required" validate:"ascii"`
		WarehouseID int    `json:"warehouse_id" binding:"required" validate:"numeric"`
	}
	return func(c *gin.Context) {
		var employee request
		var employeeStore domain.Employee
		var errorsBind = []string{}
		idEmployee := c.Param("id")

		v := validator.New()
		errJSON := c.ShouldBindJSON(&employee)
		if errJSON != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Algun dato en el body esta mal ingresado"))
			return
		}
		errBind := v.Struct(employee)
		if errBind != nil {
			for _, err := range errBind.(validator.ValidationErrors) {
				mensaje := fmt.Sprintf("el campo %s es requerido o su tipo de dato no es el correcto %s", err.ActualTag(), err.Tag())
				errorsBind = append(errorsBind, mensaje)
			}
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, strings.Join(errorsBind, ",")))
			return
		}
		employeeStore.CardNumberID = idEmployee
		employeeStore.FirstName = employee.FirstName
		employeeStore.LastName = employee.LastName
		employeeStore.WarehouseID = employee.WarehouseID
		errUpdate := e.employeeService.Update(c, employeeStore)
		if errUpdate != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "El card_number_id no ha sido encontrado o el warehouse_id no existe"))
			return
		} else {
			c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, fmt.Sprintf("El empleado ha sido modificado correctamente con id: %s", idEmployee), ""))
			return
		}
	}
}

// List Employees godoc
// @Summary Delete a employee
// @Tags Employees
// @Description Delete employee
// @Accept json
// @Param id path int true "CardNumberID"
// @Produce json
// @Success 204 {object} web.Response
// @Router /api/v1/employees/{id} [delete]
func (e *Employee) Delete() gin.HandlerFunc {

	return func(c *gin.Context) {
		errDelete := e.employeeService.Delete(c, c.Param("id"))
		if errDelete != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "El empleado no existe"))
			return
		} else {
			c.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, nil, "El empleado ha sido eiminado exitosamente"))
			return
		}
	}
}

// List Employees godoc
// @Summary List Inbound Orders by employee
// @Tags Employees
// @Description Get all employees with inboundOrders count
// @Produce json
// @Success 200 {object} web.Response
// @Router /api/v1/employees/reportInboundOrders [get]
func (e *Employee) GetInboundOrders() gin.HandlerFunc {

	return func(c *gin.Context) {

		employees, err := e.employeeService.GetInboundOrders(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		} else {
			c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, employees, ""))
			return
		}

	}
}

// List Employees godoc
// @Summary List Inbound Orders by employee
// @Tags Employees
// @Description Get employee with inboundOrders count by id
// @Produce json
// @Success 200 {object} web.Response
// @Router /api/v1/employees/reportInboundOrders/{id} [get]
func (e *Employee) GetInboundOrdersByEmployee() gin.HandlerFunc {

	return func(c *gin.Context) {
		id := c.Param("id")
		idEmployee, _ := strconv.Atoi(id)
		employees, err := e.employeeService.GetInboundOrdersByEmployee(c, idEmployee)
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		} else {
			c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, employees, ""))
			return
		}

	}
}
