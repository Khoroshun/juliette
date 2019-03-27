package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/khoroshun/juliette/models"
	u "github.com/khoroshun/juliette/utils"
	"net/http"
)


type GetBonusAccountRequest struct {
	Phone string `json:"phone"`
}

var GetBonusAccount = func(w http.ResponseWriter, r *http.Request) {

	bonusAccountRequest := GetBonusAccountRequest{}

	err := json.NewDecoder(r.Body).Decode(&bonusAccountRequest)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	client := models.GetClientByPhone(bonusAccountRequest.Phone)
	bonusAccount := models.GetBonusAccountByClientID(client.ID)
	bonusTransactions := models.GetBonusTransactions(bonusAccount.ID)

	resp := u.Message(true, "success")
	resp["bonusAccount"] = bonusAccount
	resp["client"] = client
	resp["bonusTransactions"] = bonusTransactions

	u.Respond(w, resp)
}


var CreateBonusAccount = func(clientID uint) (*models.BonusAccount) {

	bonusAccount := models.GetBonusAccountByClientID(clientID)
	if bonusAccount == nil {
		fmt.Println("create new bAccount")
		bonusAccount := &models.BonusAccount{}
		bonusAccount.Client = clientID
		bonusAccount.Summ = 0
		bonusAccount.Status = "active"
		fmt.Println(bonusAccount.Create())
	}
	bonusAccount = models.GetBonusAccountByClientID(clientID)
	return bonusAccount
}
