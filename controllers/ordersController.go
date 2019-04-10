package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/khoroshun/juliette/models"
	u "github.com/khoroshun/juliette/utils"
	"net/http"
)

type request struct {
	OrderNum string `json:"order_num"`
	Phone string `json:"phone"`
	Bonus uint `json:"bonus"`
}

type GetGetOrder struct {
	Phone string `json:"phone"`
}

var CreateOrder = func(Order models.Order) map[string] interface{} {

	resCreate := Order.Create()

	if resCreate["status"] == false {
		return resCreate
	}

	bonusTransaction := CreateBonusTransaction(&Order)

	resp := u.Message(true, "success")
	resp["Order"] = Order
	resp["bonusTransaction"] = bonusTransaction

	return resp
}


var CreateOrderHandler = func(w http.ResponseWriter, r *http.Request) {

	source := r.Context().Value("user") . (uint) //Grab the id of the user that send the request
	order := &models.Order{}
	request := &request{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	fmt.Print(request.OrderNum)
	getOrder := models.GetOrderByNum(request.OrderNum)
	if getOrder != nil {
		resp := u.Message(false, "order already exists")
		resp["Order"] = getOrder
		u.Respond(w, resp)
		return
	}

	client := models.GetClientByPhone(request.Phone)
	if client == nil {
		newClient := &models.Client{}
		newClient.Phone = request.Phone
		newClient.Name = "No name"
		CreateClient(*newClient) // создаем покупателя и его бонусный счет
	}
	client = models.GetClientByPhone(request.Phone)

	order.Client = client.ID
	order.OrderNum = request.OrderNum
	order.Bonus = request.Bonus
	order.Source = source

	resp := CreateOrder(*order) // созздаем заказ и транзакцию по заказу
	u.Respond(w, resp)
}

var UpdateOrderHandler = func(w http.ResponseWriter, r *http.Request) {

	source  := r.Context().Value("user") . (uint)
	request := request{}
	order   := &models.Order{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	client  := models.GetClientByPhone(request.Phone)

	order.OrderNum = request.OrderNum
	order.Client = client.ID
	order.Bonus = request.Bonus
	order.Source = source

	resp := order.Update()
	u.Respond(w, resp)
}

var GetOrderHandler = func(w http.ResponseWriter, r *http.Request) {

	source := r.Context().Value("user") . (uint) //Grab the id of the user that send the request
	source = source

	res 				:= request{}

	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}



	//resp := GetOrder
	//u.Respond(w, resp)
}
