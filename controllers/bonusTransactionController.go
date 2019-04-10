package controllers

import (
	"encoding/json"
	"github.com/khoroshun/juliette/models"
	u "github.com/khoroshun/juliette/utils"
	"net/http"
	"strings"
	"time"
)

type CreateBonusTransactionRequest struct {
	Phone string `json:"phone"`
	Summ uint `json:"summ"`
	Reason string `json:"reason"`
}



var CreateBonusTransaction = func(Order *models.Order) *models.BonusTransaction {

	bonusTransaction := &models.BonusTransaction{}

	bonusAccount := models.GetBonusAccountByClientID(Order.Client)

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

	//user := r.Context().Value("user") . (uint) //Grab the id of the user that send the request

	bonusTransaction := &models.BonusTransaction{}

	createBonusTransactionRequest := CreateBonusTransactionRequest{}
	err := json.NewDecoder(r.Body).Decode(&createBonusTransactionRequest)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	client := models.GetClientByPhone(createBonusTransactionRequest.Phone)
	if client == nil {
		u.Respond(w, u.Message(false, "Error client not found 1"))
		return
	}
	bonusAccount := models.GetBonusAccountByClientID(client.ID)

	bonusTransaction.Account 	= bonusAccount.ID
	bonusTransaction.Summ 		= createBonusTransactionRequest.Summ
	bonusTransaction.Reason 	= createBonusTransactionRequest.Reason
	bonusTransaction.Date 		= time.Now().String()

	resp := bonusTransaction.Create()
	u.Respond(w, resp)
}

var GetBonusTransactionsHandler = func(w http.ResponseWriter, r *http.Request) {

	bonusTransactionsRequest := CreateBonusTransactionRequest{}

	err := json.NewDecoder(r.Body).Decode(&bonusTransactionsRequest)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	client := models.GetClientByPhone(bonusTransactionsRequest.Phone)
	bonusAccount := models.GetBonusAccountByClientID(client.ID)
	bonusTransactions := models.GetBonusTransactions(bonusAccount.ID)

	resp := u.Message(true, "success")
	resp["bonusTransactions"] = bonusTransactions

	u.Respond(w, resp)
}

var UpdateBonusTransactionHandler = func(w http.ResponseWriter, r *http.Request) {

	bonusTransactionsRequest := CreateBonusTransactionRequest{}

	err := json.NewDecoder(r.Body).Decode(&bonusTransactionsRequest)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	client := models.GetClientByPhone(bonusTransactionsRequest.Phone)
	bonusAccount := models.GetBonusAccountByClientID(client.ID)
	bonusTransactions := models.GetBonusTransactions(bonusAccount.ID)

	resp := u.Message(true, "success")
	resp["bonusTransactions"] = bonusTransactions

	u.Respond(w, resp)
}
