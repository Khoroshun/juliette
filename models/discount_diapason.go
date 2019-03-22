package models

import (
	u "github.com/khoroshun/juliette/utils"
	"github.com/jinzhu/gorm"
	"fmt"
)

type DiscountDiapason struct {
	gorm.Model
	For uint `json:"for"`
	To uint `json:"to"`
	Percent uint `json:"percent"`
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (DiscountDiapason *DiscountDiapason) Validate() (map[string] interface{}, bool) {



	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (discountDiapason *DiscountDiapason) Create() (map[string] interface{}) {

	if resp, ok := discountDiapason.Validate(); !ok {
		return resp
	}

	GetDB().Create(discountDiapason)

	resp := u.Message(true, "success")
	resp["discountDiapason"] = discountDiapason
	return resp
}


func GetDiscountDiapason(user uint) ([]*DiscountDiapason) {

	discountDiapason := make([]*DiscountDiapason, 0)
	err := GetDB().Table("discountDiapason").Where("user_id = ?", user).Find(&discountDiapason).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return discountDiapason
}
