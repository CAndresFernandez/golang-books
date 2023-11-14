package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// database points to ORM
var (
	db * gorm.DB
)

// database connection
func Connect(){
	dsn := "books:books@(127.0.0.1:3306)/gorest?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

// function to recuperate the db elsewhere in the app
func GetDB() *gorm.DB{
	return db
}