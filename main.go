package main

import (

    _ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
    "github.com/gin-gonic/gin"

	"go-base-cleancode/infrastructure/config"
	"go-base-cleancode/infrastructure/middleware"

	_userHttpDeliver "go-base-cleancode/entities/user/delivery/http"
	_productHttpDeliver "go-base-cleancode/entities/product/delivery/http"
	_loginHttpDeliver "go-base-cleancode/entities/login/delivery/http"
)


func main() {
	db := config.Init()

	route := gin.Default()

	middleware.Cors(route)
	middleware.Session(route)

	
	_userHttpDeliver.NewUserHttpHandler(route, db)
	_productHttpDeliver.NewProductHttpHandler(route, db)
	_loginHttpDeliver.NewLoginHttpHandler(route, db)


	route.Run(viper.GetString("server.address"))
}