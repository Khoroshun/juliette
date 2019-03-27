package controllers

import (
	"fmt"
	"github.com/khoroshun/juliette/models"
	"net/http"
)


var GetDiscountAccount = func(w http.ResponseWriter, r *http.Request) {

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
