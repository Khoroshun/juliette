package models

import (
	"github.com/jinzhu/gorm"
	u "github.com/khoroshun/juliette/utils"
)

type Client struct {
	gorm.Model
	Phone string `json:"phone"`
	Name string `json:"name"`
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (Client *Client) Validate() (map[string] interface{}, bool) {


	if Client.Phone == "" {
		return u.Message(false, "Phone number should be on the payload"), false
	}

	//if Client.Name == ""  {
	//	return u.Message(false, "User is not recognized"), false
	//}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (Client *Client) Create() (map[string] interface{}) {

	if resp, ok := Client.Validate(); !ok {
		return resp
	}

	GetDB().Create(Client)

	resp := u.Message(true, "success")
	resp["Client"] = Client
	return resp
}

func GetClientByPhone(phone string) (*Client) {

	Client := &Client{}
	err := GetDB().Table("clients").Where("phone = ?", phone).First(Client).Error
	if err != nil {
		return nil
	}
	return Client
}

