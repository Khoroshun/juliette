package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/khoroshun/juliette/models"
	u "github.com/khoroshun/juliette/utils"
	"net/http"
)


type GetDiscountAccountRequest struct {
	Phone string `json:"phone"`
}

var GetDiscountAccount = func(w http.ResponseWriter, r *http.Request) {

	discountAccountRequest := GetDiscountAccountRequest{}

	err := json.NewDecoder(r.Body).Decode(&discountAccountRequest)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	client := models.GetClientByPhone(discountAccountRequest.Phone)
	discountAccount := models.GetDiscountAccountByClientID(client.ID)
	discountTransactions := models.GetDiscountChanges(discountAccount.ID)

	resp := u.Message(true, "success")
	resp["discountAccount"] = discountAccount
	resp["client"] = client
	resp["discountTransactions"] = discountTransactions

	u.Respond(w, resp)
}








var CreateDiscountAccount = func(clientID uint) (*models.DiscountAccount) {

	discountAccount := models.GetDiscountAccountByClientID(clientID)
	if discountAccount == nil {
		fmt.Println("create new bAccount")
		discountAccount := &models.DiscountAccount{}
		discountAccount.Client = clientID
		discountAccount.Percent = 0
		discountAccount.Status = "active"
		fmt.Println(discountAccount.Create())
	}
	discountAccount = models.GetDiscountAccountByClientID(clientID)
	return discountAccount
}
