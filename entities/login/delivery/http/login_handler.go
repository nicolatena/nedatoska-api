package http

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"



	"go-base-cleancode/entities/login/repository"
)



func NewLoginHttpHandler(route *gin.Engine, db *gorm.DB){

	handler := &repository.InDB{DB: db}
	
	v1 := route.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", handler.LoginHandler)
			auth.POST("/logout", handler.LogoutHandler)
		}
	}
}