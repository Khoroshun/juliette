package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "github.com/khoroshun/juliette/utils"
)

type BonusTransaction struct {
	gorm.Model
	Account uint `json:"account"`
	Summ uint `json:"summ"`
	Reason string `json:"reason"`
	Date string `json:"date"`
	Source uint `json:"source"`
	Num string `json:"num"`
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (bonusTransaction *BonusTransaction) Validate() (map[string] interface{}, bool) {

	//if BonusTransaction.Client  <= 0  {
	//	return u.Message(false, "BonusTransaction name should be on the payload"), false
	//}
	//
	//if BonusTransaction.Summ == "" {
	//	return u.Message(false, "Phone number should be on the payload"), false
	//}
	//
	//if BonusTransaction.Status == ""  {
	//	return u.Message(false, "User is not recognized"), false
	//}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (bonusTransaction *BonusTransaction) Create() map[string] interface{} {

	if resp, ok := bonusTransaction.Validate(); !ok {
		return resp
	}

	resp := u.Message(true, "success")

	if GetDB().Where("num = ?", bonusTransaction.Num) == nil {

		GetDB().Create(bonusTransaction)
		bonusAccount := GetBonusAccount(bonusTransaction.Account)
		GetDB().Model(bonusAccount).Update("Summ",bonusAccount.Summ + bonusTransaction.Summ)
		resp["bonusTransaction"] = bonusTransaction

	}else{
		resp = u.Message(false, "failure")
	}

	return resp
}


func GetBonusTransaction(id uint) (*BonusTransaction) {

	bonusTransaction := &BonusTransaction{}
	err := GetDB().Table("bonusTransactions").Where("id = ?", id).First(bonusTransaction).Error
	if err != nil {
		return nil
	}
	return bonusTransaction
}

func GetBonusTransactions(account uint) ([]*BonusTransaction) {

	bonusTransactions := make([]*BonusTransaction, 0)
	err := GetDB().Table("bonus_transactions").Where("account = ?", account).Find(&bonusTransactions).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return bonusTransactions
}
