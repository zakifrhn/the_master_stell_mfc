package routers

import (
	"test-hiring/internal/handlers"
	"test-hiring/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func users(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/user")

	repo := repositories.NewUser(d)
	handler := handlers.New_User(repo)

	route.GET("/:id", handler.GetDataById)
	route.POST("/", handler.PostData)
	route.PUT("/:id", handler.UpdateData)
	route.DELETE("/:id", handler.DeleteData)
}
