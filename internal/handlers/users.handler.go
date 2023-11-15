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

type HandlerUser struct {
	*repositories.RepoUser
}

func New_User(r *repositories.RepoUser) *HandlerUser {
	return &HandlerUser{r}
}

func (h *HandlerUser) PostData(ctx *gin.Context) {
	user := models.User{
		Role: "user",
	}
	var err_val error

	if err := ctx.ShouldBind(&user); err != nil {
		fmt.Println(err)
		pkg.Responses(400, &config.Result{Message: err_val.Error()}).Send(ctx)
		return
	}

	_, err_val = govalidator.ValidateStruct(&user)
	if err_val != nil {
		pkg.Responses(401, &config.Result{
			Message: err_val.Error()}).Send(ctx)
		return
	}

	cek_email := h.GetUserByEmail(user.Email_user)
	if cek_email > 0 {
		pkg.Responses(400, &config.Result{Message: "Email has been registered"}).Send(ctx)
	}

	hash_pass, err_hash := pkg.HashPassword(user.Pass_user)
	if err_hash != nil {
		pkg.Responses(400, &config.Result{Message: err_hash.Error()}).Send(ctx)
		return
	}
	user.Pass_user = hash_pass

	respone, err := h.CreateUser(&user)
	if err != nil {
		fmt.Println(err)
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	pkg.Responses(201, &config.Result{Message: respone}).Send(ctx)

}

func (h *HandlerUser) UpdateData(ctx *gin.Context) {
	var user models.User
	user.Id_user = ctx.Param("id")
	if err := ctx.ShouldBind(&user); err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	cek_user := h.GetUserById(user.Id_user)
	if cek_user == 0 {
		pkg.Responses(400, &config.Result{Message: "data not found."}).Send(ctx)

		return
	}

	cek_email := h.GetUserByEmail(user.Email_user)
	if cek_email > 0 {
		pkg.Responses(400, &config.Result{Message: "email has been used by another user."}).Send(ctx)
		return
	}

	var err_val error
	_, err_val = govalidator.ValidateStruct(&user)
	if err_val != nil {
		pkg.Responses(400, &config.Result{Message: err_val.Error()}).Send(ctx)
		return
	}

	hash_pass, err_has := pkg.HashPassword(user.Pass_user)
	if err_has != nil {
		pkg.Responses(400, &config.Result{Message: err_has.Error()}).Send(ctx)
		return
	}
	user.Pass_user = hash_pass

	response, err := h.UpdateUser(&user)
	if err != nil {
		pkg.Responses(304, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.Responses(203, &config.Result{Message: response}).Send(ctx)
}

func (h *HandlerUser) DeleteData(ctx *gin.Context) {
	var user models.User
	user.Id_user = ctx.Param("id")

	if err := ctx.ShouldBind(&user); err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	count_by_id := h.GetUserById(user.Id_user)
	if count_by_id == 0 {
		pkg.Responses(400, &config.Result{Message: "data not found."}).Send(ctx)
		return
	}

	response1, err1 := h.DeleteUser(&user)
	if err1 != nil {
		pkg.Responses(304, &config.Result{Message: err1.Error()}).Send(ctx)
		return
	}

	pkg.Responses(203, &config.Result{Message: response1}).Send(ctx)
}

func (h *HandlerUser) GetDataById(ctx *gin.Context) {
	var user models.User
	user.Id_user = ctx.Param("id")
	if err := ctx.ShouldBind(&user); err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.GetInfoUserById(&user)
	if err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.Responses(200, response).Send(ctx)
}
