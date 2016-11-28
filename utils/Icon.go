package utils


import (
	"fmt"
)

var ActionTypeDict map[string]interface{}


func (cat *CallbackActionType)ActionIcon() (retbool bool, err error) {
	const dataSourceName  = "root:root123admin@tcp(192.168.1.6:3306)/registry?autocommit=true"
	err = UpdateLogo(cat.Uuid, cat.Filename, dataSourceName)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	fmt.Println("shhh")

	return true, err
}

