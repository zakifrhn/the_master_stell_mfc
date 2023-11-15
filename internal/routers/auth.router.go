package routers

import (
	"test-hiring/internal/handlers"
	"test-hiring/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func auth(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/")

	repo := repositories.New_Auth(d)
	handler := handlers.New_Auth(repo)

	route.POST("/login", handler.Login)

}
