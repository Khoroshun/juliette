package controllers

import (
	"github.com/khoroshun/juliette/models"
)


func CreateClient (res response)  (*models.Client){

	client := models.GetClientByPhone(res.Phone)
	if client == nil { // if new client
		client := &models.Client{}
		client.Name = "anonim"
		client.Phone = res.Phone
		client.Create()
	}
	client = models.GetClientByPhone(res.Phone)
	return client
}
