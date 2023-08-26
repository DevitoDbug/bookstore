package config

import "github.com/jinzhu/gorm"

var db *gorm.DB

func Connect() {
	connectionString := "david:david@12@/bookstore?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetBD() *gorm.DB {
	return db
}
