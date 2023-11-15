package handlers

import (
	"fmt"
	"test-hiring/config"
	"test-hiring/internal/models"
	"test-hiring/internal/repositories"
	"test-hiring/pkg"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type HandlerEmployee struct {
	repositories.RepoEmployeeIF
}

func New_Employee(r *repositories.RepoEmployee) *HandlerEmployee {
	return &HandlerEmployee{r}
}

func (h *HandlerEmployee) PostDataEmployee(ctx *gin.Context) {
	var emp models.Employee
	var err_val error

	if err := ctx.ShouldBind(&emp); err != nil {
		pkg.Responses(400, &config.Result{Message: err_val.Error()}).Send(ctx)
		return
	}

	_, err_val = govalidator.ValidateStruct(&emp)
	if err_val != nil {
		pkg.Responses(401, &config.Result{
			Message: err_val.Error()}).Send(ctx)
		return
	}

	respone, err := h.CreateEmployee(&emp)
	if err != nil {
		fmt.Println(err)
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	pkg.Responses(200, &config.Result{Message: respone}).Send(ctx)

}

func (h *HandlerEmployee) UpdateDataEmployee(ctx *gin.Context) {
	var emp models.Employee

	emp.Id_Employee = ctx.Param("id")

	count_by_id := h.GetDataById(emp.Id_Employee)
	if count_by_id == 0 {
		pkg.Responses(400, &config.Result{Message: "data not found."}).Send(ctx)
		return
	}

	if err := ctx.ShouldBind(&emp); err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	var err_val error
	_, err_val = govalidator.ValidateStruct(&emp)
	if err_val != nil {
		pkg.Responses(400, &config.Result{Message: err_val.Error()}).Send(ctx)
		return
	}

	response, err := h.UpdateEmployee(&emp)
	if err != nil {
		pkg.Responses(304, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.Responses(203, &config.Result{Message: response}).Send(ctx)
}

func (h *HandlerEmployee) DeleteDataEmployee(ctx *gin.Context) {
	var emp models.Employee
	emp.Id_Employee = ctx.Param("id")

	count_by_id := h.GetDataById(emp.Id_Employee)
	if count_by_id == 0 {
		pkg.Responses(400, &config.Result{Message: "data not found."}).Send(ctx)
		return
	}

	if err := ctx.ShouldBind(&emp); err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.DeleteEmployee(&emp)
	if err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.Responses(200, &config.Result{Message: response}).Send(ctx)
}

func (h *HandlerEmployee) DetailDataEmployee(ctx *gin.Context) {
	var emp models.Employee
	emp.Id_Employee = ctx.Param("id")
	if err := ctx.ShouldBind(&emp); err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.GetInfoEmployee(&emp)
	if err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.Responses(200, response).Send(ctx)
}

func (h *HandlerEmployee) GetAllDataEmployee(ctx *gin.Context) {
	var user models.Employee
	page := ctx.Query("page")
	limit := ctx.Query("limit")

	if err := ctx.ShouldBind(&user); err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.GetAllEmployee(&user, page, limit)
	if err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.Responses(200, response).Send(ctx)
}
