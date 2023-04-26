package models

import(
	"fmt"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"time"
)

type Customer struct {
	Customer_Number int64 `gorm:"primaryKey" json:"customer_number"`
	Customer_Type string `gorm:"check:Customer_Type in ('Individual','Corporate','Premier')" json:"customer_type"`
	Customer_Name string `json:"customer_name"`
	Customer_Country_Code string `json:"customer_country_type"`
	Customer_State string `json:"customer_state"`
	Customer_Postal_Code string `json:"customer_postal_code"`
	Customer_City string `json:"customer_city"`
	Customer_Address string `json:"customer_address"`
	Customer_Phone string `json:"customer_phone"`
	Account []Account `gorm:"foreignKey:Customer_Number;references:Customer_Number;constraint:OnUpdate:CASCADE,onDelete:CASCADE;" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`
	DeletedAt soft_delete.DeletedAt `gorm:"index" json:"-"`
}

func (c *Customer) CreateCustomer() (*Customer, error){
	// db.NewRecord(c)
	result := db.Create(&c)
	return c, result.Error
}

func GetAllCustomers() []Customer{
	fmt.Println("GetAllCustomers called")
	var Customers []Customer
	db.Find(&Customers)
	return Customers
}

func GetCustomerById(Id int64) (*Customer, *gorm.DB, error){
	var getCustomer Customer
	db := db.Where("Customer_Number=?", Id).Find(&getCustomer)
	var err error = nil
	if getCustomer.Customer_Number == 0 {
		err = fmt.Errorf("Customer cannot be found with Customer_Number = %v", Id)
	}
	return &getCustomer, db, err
}

func DeleteCustomer(Id int64) (*Customer, error){
	customer, _, err := GetCustomerById(Id)
	if err == nil {
		db.Delete(&customer)
	}
	return customer, err
}