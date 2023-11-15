package handlers

import (
	"test-hiring/config"
	"test-hiring/internal/models"
	"test-hiring/internal/repositories"
	"test-hiring/pkg"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type Handler_Auth struct {
	repositories.Repo_AuthIF
}

func New_Auth(r repositories.Repo_AuthIF) *Handler_Auth {
	return &Handler_Auth{r}
}

func (h *Handler_Auth) Login(ctx *gin.Context) {
	var user models.Auth
	var err_val error

	if err := ctx.ShouldBind(&user); err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	_, err_val = govalidator.ValidateStruct(&user)
	if err_val != nil {
		pkg.Responses(400, &config.Result{Message: err_val.Error()}).Send(ctx)
		return
	}

	response, err := h.GetUser(&user)
	if err != nil {
		pkg.Responses(401, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	if err := pkg.VerifyPassword(response.Pass_user, user.Pass_user); err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	jwtt := pkg.NewToken(response.Id_user, response.Role, response.Email_user)
	tokens, err := jwtt.Generate()
	if err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	pkg.Responses(200, &config.Result{Token: tokens}).Send(ctx)
}
