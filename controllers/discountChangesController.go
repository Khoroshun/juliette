package controllers

import (
	"encoding/json"
	"github.com/khoroshun/juliette/models"
	u "github.com/khoroshun/juliette/utils"
	"net/http"
	"time"
)

type CreateDiscountChangesRequest struct {
	Phone string `json:"phone"`
	Percent uint `json:"percent"`
	Reason string `json:"reason"`
}

type GetDiscountChangesRequest struct {
	Phone string `json:"phone"`
}

var CreateDiscountChanges = func(w http.ResponseWriter, r *http.Request) {

	discountChanges := &models.DiscountChanges{}

	createDiscountChangesRequest := CreateDiscountChangesRequest{}
	err := json.NewDecoder(r.Body).Decode(&createDiscountChangesRequest)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	client := models.GetClientByPhone(createDiscountChangesRequest.Phone)
	if client == nil {
		u.Respond(w, u.Message(false, "Error client not found"))
		return
	}

	discountChanges.Account 	= CreateDiscountAccount(client.ID).ID
	discountChanges.Percent		= createDiscountChangesRequest.Percent
	discountChanges.Reason 		= createDiscountChangesRequest.Reason
	discountChanges.Date 		= time.Now().String()

	resp := discountChanges.Create()
	u.Respond(w, resp)

}

var GetDiscountChanges = func(w http.ResponseWriter, r *http.Request) {

	discountChangesRequest := GetDiscountChangesRequest{}

	err := json.NewDecoder(r.Body).Decode(&discountChangesRequest)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	client := models.GetClientByPhone(discountChangesRequest.Phone)
	discountAccount := models.GetDiscountAccountByClientID(client.ID)
	discountTransactions := models.GetDiscountChanges(discountAccount.ID)

	resp := u.Message(true, "success")
	resp["discountTransactions"] = discountTransactions

	u.Respond(w, resp)
}
