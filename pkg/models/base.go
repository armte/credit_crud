package models

import(
	"gorm.io/gorm"
	"github.com/armte/credit_crud/pkg/config"
)

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Customer{}, &Account{}, &Card{})
}

