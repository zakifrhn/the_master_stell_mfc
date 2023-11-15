package config

import "github.com/gin-contrib/cors"

type Metas struct {
	Next       interface{} `json:"next"`
	Prev       interface{} `json:"prev"`
	Last_page  interface{} `json:"last_page"`
	Total_data interface{} `json:"total_data"`
}

type Result struct {
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta"`
	Message interface{} `json:"message"`
	Token   interface{} `jsno:"token"`
}

var CorsConfig = cors.Config{
	AllowOrigins:     []string{"*"},
	AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "HEAD", "OPTIONS"},
	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	AllowCredentials: true,
}
