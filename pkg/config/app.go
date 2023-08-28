package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	connectionString := "root:davi@tcp(localhost:3306)/bookstore?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetBD() *gorm.DB {
	return db
}
