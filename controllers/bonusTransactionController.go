package controllers

import (
	"github.com/khoroshun/juliette/models"
	"net/http"
	"strings"
	"time"
)

func CreateBonusTransaction (accountID uint,res response) {

	bonusTransaction := &models.BonusTransaction{}

	bonusTransaction.Account = accountID
	bonusTransaction.Summ = res.Bonus
	bonusTransaction.Reason =  strings.Join([]string{"Order #", res.OrderNum}, "")
	bonusTransaction.Date =  time.Now().String()
	//bonusTransaction.Source = source
	bonusTransaction.Create()
}

var GetBonusTransaction = func(w http.ResponseWriter, r *http.Request) {

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
