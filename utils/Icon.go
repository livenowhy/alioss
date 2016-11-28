package utils

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/liuzhangpei/alioss/models"
)

var ActionTypeDict map[string]interface{}


func (cat *CallbackActionType)ActionIcon() (retbool bool, err error) {
	db_session, err := gorm.Open("mysql", "root:root123admin@tcp(192.168.1.6:3306)/registry?charset=utf8")

	if err != nil {
		fmt.Println("lzp --> CreateEngine : %s", err.Error())
		return false, err
	}
	defer db_session.Close()

	var userBase models.UserBase
	userBase.User_id = cat.Uuid

	db := db_session.Model(&userBase).Update("logo", cat.Filename)


	if db.Error != nil {
		fmt.Println("---lzp db.Error != nil   ")
		fmt.Println(db.Error.Error())
		return false, db.Error
	}
	fmt.Println("db.RowsAffected")
	fmt.Println(db.RowsAffected)
	if db.RowsAffected >= 1 {
		return true, nil
	} else {
		return false, nil
	}
}

