package models

import(
	"fmt"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"time"
)

type Account struct {
	Account_Number int64 `gorm:"primaryKey" json:"account_number"`
	Account_Type string `gorm:"not null;check:Account_Type in ('Checking','Savings','Loan')" json:"account_type"`
	Account_State string `gorm:"not null" json:"account_state"`
	Account_Postal_Code string `gorm:"not null" json:"account_postal_code"`
	Account_Country_Code string `gorm:"not null" json:"account_country_code"`
	Customer_Number int64 `gorm:"not null" json:"customer_number"`
	Card []*Card `gorm:"foreignKey:Account_Number;references:Account_Number;constraint:OnUpdate:CASCADE,onDelete:CASCADE;" json:"cards"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`
	DeletedAt soft_delete.DeletedAt `gorm:"index" json:"-"`
}

func (a *Account) CreateAccount() (*Account, error){
	_, _, errGet := GetCustomerById(a.Customer_Number)
	if errGet != nil {
		return a, errGet
	}
	result := db.Create(&a)
	return a, result.Error
}

func GetAllAccounts() []Account{
	fmt.Println("GetAllAccounts called")
	var Accounts []Account
	db.Preload("Card").Find(&Accounts)
	return Accounts
}

func GetAccountById(Id int64) (*Account, *gorm.DB, error){
	var getAccount Account
	db := db.Preload("Card").Where("Account_Number=?", Id).Find(&getAccount)
	var err error = nil
	if getAccount.Account_Number == 0 {
		err = fmt.Errorf("Account cannot be found with Account_Number = %v", Id)
	}
	return &getAccount, db, err
}

func DeleteAccount(Id int64) (*Account, error){
	account, _, err := GetAccountById(Id)
	if err == nil {
		db.Select("Card").Preload("Card").Delete(&account)
	}
	return account, err
}
