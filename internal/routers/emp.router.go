package routers

import (
	"test-hiring/internal/handlers"
	"test-hiring/internal/middleware"
	"test-hiring/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func employee(g *gin.Engine, db *sqlx.DB) {
	route := g.Group("/emp")

	repo := repositories.NewEmployee(db)
	handler := handlers.New_Employee(repo)

	route.GET("/", handler.GetAllDataEmployee)
	route.GET("/:id", middleware.AuthJwt("admin"), handler.DetailDataEmployee)
	route.POST("/", middleware.AuthJwt("admin"), handler.PostDataEmployee)
	route.PUT("/:id", middleware.AuthJwt("admin"), handler.UpdateDataEmployee)
	route.DELETE("/:id", middleware.AuthJwt("admin"), handler.DeleteDataEmployee)
}
