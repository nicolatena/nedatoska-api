package http

import (
	// "net/http"
	// "time"
	// "fmt"
	
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"



	"go-base-cleancode/entities/product/repository"

)



func NewProductHttpHandler(route *gin.Engine, db *gorm.DB){

	handler := &repository.InDB{DB: db}
	
	v1 := route.Group("/v1")
	{
		api := v1.Group("/api")
		{
			users := api.Group("/product")
			// users.Use(AuthRequired())
			{
				users.GET("/select", handler.Fetch)
				users.POST("/insert", handler.Store)
				users.PUT("/update/:id", handler.Update)
				users.DELETE("/delete/:id", handler.Delete)
			}
		}
	}
}