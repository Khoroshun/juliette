package controllers

import (
	"encoding/json"
	"github.com/khoroshun/juliette/models"
	u "github.com/khoroshun/juliette/utils"
	"net/http"
)

type response struct {
	OrderNum string `json:"order_num"`
	Phone string `json:"phone"`
	Bonus uint `json:"bonus"`
}

var CreateOrder = func(w http.ResponseWriter, r *http.Request) {

	//source := r.Context().Value("user") . (uint) //Grab the id of the user that send the request

	order  				:= &models.Order{}
	res 				:= response{}

	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	client := CreateClient(res);

	bonusAccount	:= CreateBonusAccount(client.ID);
	CreateBonusTransaction(bonusAccount.ID,res)

	order.Client = client.ID
	order.OrderNum = res.OrderNum
	order.Bonus = res.Bonus

	//fmt.Println(source)
	//order.Source = source

	resp := order.Create()
	u.Respond(w, resp)
}
