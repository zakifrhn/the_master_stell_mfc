package routers

import (
	"test-hiring/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Routers(db *sqlx.DB) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(config.CorsConfig))

	users(router, db)
	auth(router, db)
	employee(router, db)
	return router
}
