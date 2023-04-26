package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var (
	db *gorm.DB
)

func Connect(){
	dsn := "toarms:smraot@tcp(localhost:33061)/credit_crud?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}