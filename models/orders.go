package models

import (
	u "github.com/khoroshun/juliette/utils"
	"github.com/jinzhu/gorm"
	"fmt"
)

type Order struct {
	gorm.Model
	OrderNum string `json:"order_num"`
	Summ uint `json:"summ"`
	Source uint `json:"source"`
	Client uint `json:"client"`
	Bonus uint `json:"bonus"`

}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (Order *Order) Validate() (map[string] interface{}, bool) {

	if Order.Client  <= 0  {
		return u.Message(false, "Client should be on the payload"), false
	}

	//if Order.Summ  <= 0  {
	//	return u.Message(false, "Phone number should be on the payload"), false
	//}

	//if Order.Source <= 0  {
	//	return u.Message(false, "User is not recognized"), false
	//}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (Order *Order) Create() (map[string] interface{}) {

	if resp, ok := Order.Validate(); !ok {
		return resp
	}

	GetDB().Create(Order)

	resp := u.Message(true, "success")
	resp["Order"] = Order
	return resp
}

func GetOrder(id uint) (*Order) {

	Order := &Order{}
	err := GetDB().Table("Orders").Where("id = ?", id).First(Order).Error
	if err != nil {
		return nil
	}
	return Order
}

func GetOrders(user uint) ([]*Order) {

	Orders := make([]*Order, 0)
	err := GetDB().Table("Orders").Where("user_id = ?", user).Find(&Orders).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return Orders
}
