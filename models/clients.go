package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "github.com/khoroshun/juliette/utils"
	"io/ioutil"
	"net/http"
	"strings"
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

func (Client *Client) Create() map[string] interface{} {

	if resp, ok := Client.Validate(); !ok {
		return resp
	}

	GetDB().Create(Client)

	resp := u.Message(true, "success")
	resp["Client"] = Client


	url := "http://turbosms.in.ua/api/soap.html"

	payload := strings.NewReader("<soapenv:Envelope xmlns:soapenv=\"http://schemas.xmlsoap.org/soap/envelope/\" xmlns:tur=\"http://turbosms.in.ua/api/Turbo\">\n   <soapenv:Header/>\n   <soapenv:Body>\n      <tur:Auth>\n         <!--Optional:-->\n         <tur:login>JulietteBrand</tur:login>\n         <!--Optional:-->\n         <tur:password>0997740160jb</tur:password>\n      </tur:Auth>\n   </soapenv:Body>\n</soapenv:Envelope>")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/xml")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "510cac2e-3980-4c9a-9a60-1c9dfe2afe53")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))


	payload2 := strings.NewReader("<soapenv:Envelope xmlns:soapenv=\"http://schemas.xmlsoap.org/soap/envelope/\" xmlns:tur=\"http://turbosms.in.ua/api/Turbo\">\n   <soapenv:Header/>\n   <soapenv:Body>\n      <tur:SendSMS>\n         <!--Optional:-->\n         <tur:sender>Juliette</tur:sender>\n         <!--Optional:-->\n         <tur:destination>380967154107</tur:destination>\n         <!--Optional:-->\n         <tur:text>text text</tur:text>\n         <!--Optional:-->\n         <tur:wappush>?</tur:wappush>\n      </tur:SendSMS>\n   </soapenv:Body>\n</soapenv:Envelope>")

	req2, _ := http.NewRequest("POST", url, payload2)

	req.Header.Add("Content-Type", "application/xml")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "5b8209f8-0bd9-48e7-80bd-c8792feb280a")

	res2, _ := http.DefaultClient.Do(req2)

	defer res.Body.Close()
	body2, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res2)
	fmt.Println(string(body2))



	return resp
}

func (Client *Client) Update() (map[string] interface{}) {

	if resp, ok := Client.Validate(); !ok {
		return resp
	}

	GetDB().Model(&Client).Where("id = ?",Client.ID).Updates(Client)

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

