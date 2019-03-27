package models

import (
	u "github.com/khoroshun/juliette/utils"
	"github.com/jinzhu/gorm"
	"fmt"
)

type DiscountAccount struct {
	gorm.Model
	Client uint `json:"client"`
	Percent uint `json:"percent"`
	Status string `json:"status"`
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (DiscountAccount *DiscountAccount) Validate() (map[string] interface{}, bool) {



	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (discountAccount *DiscountAccount) Create() (map[string] interface{}) {

	if resp, ok := discountAccount.Validate(); !ok {
		return resp
	}

	GetDB().Create(discountAccount)

	resp := u.Message(true, "success")
	resp["discountAccount"] = discountAccount
	return resp
}


func GetDiscountAccount(user uint) ([]*DiscountAccount) {

	discountAccount := make([]*DiscountAccount, 0)
	err := GetDB().Table("discountAccount").Where("user_id = ?", user).Find(&discountAccount).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return discountAccount
}

func GetDiscountAccountByClientID(clientID uint) (*DiscountAccount) {

	discountAccount := &DiscountAccount{}
	err := GetDB().Table("discount_accounts").Where("client = ?", clientID).First(discountAccount).Error
	if err != nil {
		return nil
	}
	return discountAccount
}