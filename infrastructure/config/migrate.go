package config


import (
	"github.com/jinzhu/gorm"
    . "nedatoska-api/models"
)


func Migrate(idb *gorm.DB) {

	idb.Debug().AutoMigrate(
		&User{}, 
		&Product{},
	)

}