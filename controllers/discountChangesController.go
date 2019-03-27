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

	//user := r.Context().Value("user") . (uint) //Grab the id of the user that send the request
	//order := &models.Order{}
	//
	//err := json.NewDecoder(r.Body).Decode(order)
	//if err != nil {
	//	u.Respond(w, u.Message(false, "Error while decoding request body"))
	//	return
	//}
	//
	////order.UserId = user
	//resp := order.Create()
	//u.Respond(w, resp)
}
