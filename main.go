package main

import (

    _ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
    "github.com/gin-gonic/gin"

	"nedatoska-api/infrastructure/config"
	"nedatoska-api/infrastructure/middleware"

	_userHttpDeliver "nedatoska-api/entities/user/delivery/http"
	_productHttpDeliver "nedatoska-api/entities/product/delivery/http"
	_loginHttpDeliver "nedatoska-api/entities/login/delivery/http"
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