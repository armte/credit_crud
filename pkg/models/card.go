package models

import(
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"time"
)

type Card struct {
	Payment_Card_Number uuid.UUID `gorm:"primaryKey" json:"payment_card_number"`
	Credit_Limit int64 `gorm:"check: Credit_Limit <= 10000" json:"credit_limit"`
	Account_Number int64 `gorm:"not null" json:"account_number"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`
	DeletedAt soft_delete.DeletedAt `gorm:"index" json:"-"`
}

func (c *Card) CreateCard() (*Card, error){
	result := db.Create(&c)
	return c, result.Error
}

func GetAllCards() []Card{
	fmt.Println("GetAllCards called")
	var Cards []Card
	db.Find(&Cards)
	return Cards
}

func GetCardById(Id uuid.UUID) (*Card, *gorm.DB, error){
	var getCard Card
	db := db.Where("Card_Number=?", Id).Find(&getCard)
	var err error = nil
	if getCard.Payment_Card_Number == uuid.Nil {
		err = fmt.Errorf("Card cannot be found with Card_Number = %v", Id)
	}
	return &getCard, db, err
}

func DeleteCard(Id uuid.UUID) (*Card, error){
	card, _, err := GetCardById(Id)
	if err == nil {
		db.Delete(&card)
	}
	return card, err
}