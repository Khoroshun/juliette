package controllers

import (
	"fmt"
	"github.com/khoroshun/juliette/models"
	"net/http"
)


var GetBonusAccount = func(w http.ResponseWriter, r *http.Request) {

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
