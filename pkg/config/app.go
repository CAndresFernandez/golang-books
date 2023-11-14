package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// database points to ORM
var (
	db * gorm.DB
)

// database connection
func Connect(){
	d, err := gorm.Open("mysql", "books:books@(127.0.0.1:3306)/gorest?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

// function to recuperate the db elsewhere in the app
func GetDB() *gorm.DB{
	return db
}