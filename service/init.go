package service

import (
	"fmt"

	"github.com/cepljxiongjun/hotwind/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // import mysql
	"github.com/spf13/viper"
)

var (
	db *gorm.DB
)

// Init initalizes mysql pkg.
func Init() {
	var err error
	db, err = gorm.Open("mysql", viper.GetString("DB.default"))
	if err != nil {
		panic(fmt.Errorf("failed to connect database %s", err))
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.AutoMigrate(&models.Post{}, &models.Category{}, &models.User{})
}
