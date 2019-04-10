package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/khoroshun/juliette/models"
	u "github.com/khoroshun/juliette/utils"
	"net/http"
	"strconv"
)

type GetClientRequest struct {
	Phone string `json:"phone"`
}

var CreateClient = func(Client models.Client) map[string] interface{} {

	resCreate := Client.Create()

	if resCreate["status"] == false {
		return resCreate
	}

	bonusAccount := CreateBonusAccount(Client.ID)

	resp := u.Message(true, "success")
	resp["Client"] = Client
	resp["BonusAccount"] = bonusAccount

	return resp

}

var CreateClientHandler = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user") . (uint) //Grab the id of the user that send the request
	user = user

	client := models.Client{}
	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	if models.GetClientByPhone(client.Phone) != nil {
		u.Respond(w, u.Message(false, "Client found"))
		return
	}
	resp := CreateClient(client)

	u.Respond(w, resp)
}

var UpdateClientHandler = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user") . (uint) //Grab the id of the user that send the request
	user = user

	vars := mux.Vars(r)

	client := models.Client{}

	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	i, err := strconv.ParseUint(vars["id"],10,32)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	client.ID = uint(i)
	resp := client.Update()
	u.Respond(w, resp)
}

var GetClientHandler = func(w http.ResponseWriter, r *http.Request) {

	clientRequest := GetClientRequest{}

	err := json.NewDecoder(r.Body).Decode(&clientRequest)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	client := models.GetClientByPhone(clientRequest.Phone)
	bonusAccount := models.GetBonusAccountByClientID(client.ID)
	bonusTransactions := models.GetBonusTransactions(bonusAccount.ID)

	resp := u.Message(true, "success")
	resp["client"] = client
	resp["bonusAccount"] = bonusAccount
	resp["bonusTransactions"] = bonusTransactions

	u.Respond(w, resp)
}
