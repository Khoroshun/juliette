package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "github.com/khoroshun/juliette/utils"
	sms "github.com/wildsurfer/turbosms-go"
)

type BonusTransaction struct {
	gorm.Model
	Account uint   `json:"account"`
	Summ    int    `json:"summ"`
	Reason  string `json:"reason"`
	Date    string `json:"date"`
	Source  uint   `json:"source"`
	Num     string `json:"num"`
	ErpUid  string `json:"erpuid"`
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
	count := 0

	GetDB().Where("num = ?",  bonusTransaction.Num).Find(bonusTransaction).Count(&count)
	client_id := 0
	if count == 0 {

		GetDB().Create(bonusTransaction)
		bonusAccount := GetBonusAccount(bonusTransaction.Account)
		GetDB().Model(bonusAccount).Update("Summ",bonusAccount.Summ + bonusTransaction.Summ)
		resp["bonusTransaction"] = bonusTransaction
		client_id = int(bonusAccount.Client)
	}else{

		resp = u.Message(false, "failure - transaction with this number already exists")
	}

	Client := &Client{}
	GetDB().Model(Client).Where("id = ?",client_id)

	c := sms.NewClient("JulietteBrand", "0997740160jb")
	c.SendSMS("Juliette", Client.Phone, "Программа лояльности JULIETTE - начисление бонусов! https://juliette-sun.com.ua/check_bonus.php", "")

	return resp
}

func (bonusTransaction *BonusTransaction) Update() map[string] interface{} {

	if resp, ok := bonusTransaction.Validate(); !ok {
		return resp
	}

	fmt.Print(bonusTransaction.Summ)

	GetDB().Model(&bonusTransaction).Where("num = ?",bonusTransaction.Num).Updates(bonusTransaction)
	// костыль, разобраться-переделать!
	GetDB().Model(&bonusTransaction).Where("num = ?",bonusTransaction.Num).Update("summ",bonusTransaction.Summ)

	resp := u.Message(true, "success")
	resp["bonusTransaction"] = bonusTransaction
	return resp
}


func GetBonusTransaction(request map[string] interface{}) [] BonusTransaction {

	var bonusTransaction []BonusTransaction
	err := GetDB().Table("bonus_transactions").Where(request).Find(&bonusTransaction).Error
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
