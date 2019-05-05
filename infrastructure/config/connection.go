package config

import (
	"fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

func Init() *gorm.DB {
	
	viper.SetConfigFile(`go-base-cleancode/config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}

	dbhost     := viper.GetString(`database.host`)
	port       := viper.GetString(`database.port`)
	username   := viper.GetString(`database.user`)
	password   := viper.GetString(`database.pass`)
	dbname 	   := viper.GetString(`database.name`)

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s  password=%s", dbhost, port, username, dbname, password)
	conn, err := gorm.Open("postgres", dbUri)
	// defer conn.Close()
	if err != nil {
		fmt.Print(err)
	}
	
	Migrate(conn)
	
	return conn
}
