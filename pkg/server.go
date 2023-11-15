package pkg

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Server(router *gin.Engine) *http.Server {
	var addr string = "0.0.0.0:8080"

	if port := viper.GetString("port"); port != "" {
		addr = ":" + port
	}

	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 15,
		Handler:      router,
	}

	return srv
}
