package controllers

import (
	"encoding/json"
	"github.com/khoroshun/juliette/models"
	u "github.com/khoroshun/juliette/utils"
	"net/http"
)

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
	var request  map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	getOrder := models.GetOrderByNum(request["order_num"].(string))
	if getOrder != nil {
		resp := u.Message(false, "order already exists")
		resp["Order"] = getOrder
		u.Respond(w, resp)
		return
	}

	client := models.GetClientByPhone(request["phone"].(string))
	if client == nil {
		newClient := &models.Client{}
		newClient.Phone = request["phone"].(string)
		newClient.Name = "No name"
		CreateClient(*newClient) // создаем покупателя и его бонусный счет
	}
	client = models.GetClientByPhone(request["phone"].(string))

	order.Client = client.ID
	order.OrderNum = request["order_num"].(string)
	order.Bonus = request["bonus"].(int)
	order.Source = source

	resp := CreateOrder(*order) // создаем заказ и транзакцию по заказу
	u.Respond(w, resp)
}

var UpdateOrderHandler = func(w http.ResponseWriter, r *http.Request) {

	source  := r.Context().Value("user") . (uint)
	var request  map[string]interface{}
	order   := &models.Order{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	client  := models.GetClientByPhone(request["phone"].(string))

	order.OrderNum = request["order_num"].(string)
	order.Client = client.ID
	order.Bonus = request["bonus"].(int)
	order.Source = source

	resp := order.Update()
	u.Respond(w, resp)
}

var GetOrderHandler = func(w http.ResponseWriter, r *http.Request) {

	var request map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	if request["phone"] != nil{
		client := models.GetClientByPhone(request["phone"].(string))
		delete(request,"phone")
		request["client"] = client.ID
	}


	resp := u.Message(true, "success")
	resp["Orders"] = GetOrder(request)

	u.Respond(w, resp)
}

var GetOrder = func(Request map[string] interface{}) [] models.Order {

	return models.GetOrder(Request)

}
