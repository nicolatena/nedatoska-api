package middleware

import (
	"github.com/itsjamie/gin-cors"
	"time"
    "github.com/gin-gonic/gin"
)

func Cors(route *gin.Engine) {
	route.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
}