package config


import (
	"github.com/jinzhu/gorm"
    . "go-base-cleancode/models"
)


func Migrate(idb *gorm.DB) {

	idb.Debug().AutoMigrate(
		&User{}, 
		&Product{},
	)

}