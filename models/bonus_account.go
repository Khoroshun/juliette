package models

import (
	u "github.com/khoroshun/juliette/utils"
	"github.com/jinzhu/gorm"
	"fmt"
)

type BonusAccount struct {
	gorm.Model
	Client uint `json:"client"`
	Summ int `json:"summ"`
	Status string `json:"status"`
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (bonusAccount *BonusAccount) Validate() (map[string] interface{}, bool) {

	if bonusAccount.Client  <= 0  {
		return u.Message(false, "BonusAccount name should be on the payload"), false
	}

	if bonusAccount.Status == ""  {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (bonusAccount *BonusAccount) Create() (map[string] interface{}) {

	if resp, ok := bonusAccount.Validate(); !ok {
		return resp
	}

	GetDB().Create(bonusAccount)

	resp := u.Message(true, "success")
	resp["bonusAccount"] = bonusAccount
	return resp
}

func GetBonusAccount(id uint) (*BonusAccount) {

	bonusAccount := &BonusAccount{}
	err := GetDB().Table("bonus_accounts").Where("id = ?", id).First(bonusAccount).Error
	if err != nil {
		return nil
	}
	return bonusAccount
}

func GetBonusAccounts(user uint) ([]*BonusAccount) {

	bonusAccounts := make([]*BonusAccount, 0)
	err := GetDB().Table("bonusAccounts").Where("user_id = ?", user).Find(&bonusAccounts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return bonusAccounts
}

func GetBonusAccountByClientID(clientID uint) (*BonusAccount) {

	bonusAccount := &BonusAccount{}
	err := GetDB().Table("bonus_accounts").Where("client = ?", clientID).First(bonusAccount).Error
	if err != nil {
		return nil
	}
	return bonusAccount
}