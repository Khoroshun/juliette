package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/khoroshun/juliette/models"
	u "github.com/khoroshun/juliette/utils"
	sms "github.com/wildsurfer/turbosms-go"
	"net/http"
	"strings"
	"time"
	//sms "github.com/wildsurfer/turbosms-go"
)

type CreateBonusTransactionRequest struct {
	Phone string `json:"phone"`
	Summ int `json:"summ"`
	Reason string `json:"reason"`
	Num string `json:"num"`
	ErpUid string `json:"erpuid"`
}


var CreateBonusTransaction = func(Order *models.Order) *models.BonusTransaction {

	bonusTransaction := &models.BonusTransaction{}

	bonusAccount := models.GetBonusAccountByClientID(Order.Client)

	bonusTransaction.Active 	= false
	bonusTransaction.Account 	= bonusAccount.ID
	bonusTransaction.Summ 		= Order.Bonus
	bonusTransaction.Reason 	= strings.Join([]string{"За заказ #", Order.OrderNum }, "")
	bonusTransaction.Date 		= time.Now().String()
	bonusTransaction.Source 	= uint(0) // TODO: прокинуть источник
	bonusTransaction.Num		= strings.Join([]string{"За заказ #", Order.OrderNum }, "")

	bonusTransaction.Create()

	return bonusTransaction
}

var CreateBonusTransactionHandler = func(w http.ResponseWriter, r *http.Request) {


	bonusTransaction := &models.BonusTransaction{}

	createBonusTransactionRequest := CreateBonusTransactionRequest{}
	err := json.NewDecoder(r.Body).Decode(&createBonusTransactionRequest)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	client := models.GetClientByPhone(createBonusTransactionRequest.Phone)
	if client == nil {
		newClient := &models.Client{}
		newClient.Phone = createBonusTransactionRequest.Phone
		newClient.Name = "No name"
		CreateClient(*newClient)

	}
	client = models.GetClientByPhone(createBonusTransactionRequest.Phone)

	bonusAccount := models.GetBonusAccountByClientID(client.ID)

	bonusTransaction.Account 	= bonusAccount.ID
	bonusTransaction.Summ 		= createBonusTransactionRequest.Summ
	bonusTransaction.Reason 	= createBonusTransactionRequest.Reason
	bonusTransaction.Date 		= time.Now().String()
	bonusTransaction.Num		= createBonusTransactionRequest.Num
	bonusTransaction.ErpUid		= createBonusTransactionRequest.ErpUid

	c := sms.NewClient("JulietteBrand", "0997740160jb")
	sms_text := ""
	if bonusTransaction.Summ > 0 {
		sms_text = fmt.Sprintf("%s%d%s", "Благодарим за покупку! Вам начислено ", bonusTransaction.Summ, " бонусов. 1 бонус = 1 грн. Активация через 10 дней. Подробнее по ссылке https://juliette-sun.com.ua/check_bonus.php")
	}else{
		sms_text = fmt.Sprintf("%s%d%s", "Программа лояльности JULIETTE - списано бонусов ", bonusTransaction.Summ, "грн. Подробнее https://juliette-sun.com.ua/check_bonus.php")
	}
	c.SendSMS("Juliette", client.Phone, sms_text, "")


	resp := bonusTransaction.Create()
	u.Respond(w, resp)
}


var GetBonusTransactions = func(request map[string] interface{}) [] models.BonusTransaction {


	return models.GetBonusTransaction(request)

}

var GetBonusTransactionsHandler = func(w http.ResponseWriter, r *http.Request) {

	request := make(map[string]interface{}, 10)
	values := r.URL.Query()

	for i, v := range values {
		if i == "phone"{// для телефона удаляем пробелы из строки и добавляем в начале +
			v[0] = fmt.Sprintf("%s%s", "+", strings.Replace(v[0]," ","",-1))
		}
		request[i] = v[0]
	}

	if request["phone"] != nil{
		client := models.GetClientByPhone(request["phone"].(string))
		if client != nil{
			delete(request,"phone")
			bonusAccount := models.GetBonusAccountByClientID(client.ID)
			request["account"] = bonusAccount.ID
		}
	}

	resp := u.Message(true, "success")
	resp["BonusTransactions"] = GetBonusTransactions(request)

	u.Respond(w, resp)
}

var UpdateBonusTransactionHandler = func(w http.ResponseWriter, r *http.Request) {

	bonusTransactionsRequest := CreateBonusTransactionRequest{}

	bonusTransaction   := &models.BonusTransaction{}

	err := json.NewDecoder(r.Body).Decode(&bonusTransactionsRequest)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}


	bonusTransaction.Summ = bonusTransactionsRequest.Summ
	bonusTransaction.Reason = bonusTransactionsRequest.Reason
	bonusTransaction.Num = bonusTransactionsRequest.Num
	bonusTransaction.ErpUid = bonusTransactionsRequest.ErpUid

	resp := bonusTransaction.Update()
	u.Respond(w, resp)



}
