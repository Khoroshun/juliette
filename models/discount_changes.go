package models

import (
	u "github.com/khoroshun/juliette/utils"
	"github.com/jinzhu/gorm"
	"fmt"
)

type DiscountChanges struct {
	gorm.Model
	Account uint `json:"Account"`
	Percent string `json:"percent"`
	Reason string `json:"reason"`
	Date string `json:"date"`
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (DiscountChanges *DiscountChanges) Validate() (map[string] interface{}, bool) {



	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (discountChanges *DiscountChanges) Create() (map[string] interface{}) {

	if resp, ok := discountChanges.Validate(); !ok {
		return resp
	}

	GetDB().Create(discountChanges)

	resp := u.Message(true, "success")
	resp["discountChanges"] = discountChanges
	return resp
}


func GetDiscountChanges(user uint) ([]*DiscountChanges) {

	discountChanges := make([]*DiscountChanges, 0)
	err := GetDB().Table("discountChanges").Where("user_id = ?", user).Find(&discountChanges).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return discountChanges
}
