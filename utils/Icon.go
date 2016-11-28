package utils


import (
	"fmt"
)

var ActionTypeDict map[string]interface{}


func (cat *CallbackActionType)ActionIcon(dataSourceName, HostOuter string) (retbool bool, err error) {

	filename := HostOuter + "/" + cat.Filename
	err = UpdateLogo(cat.Uuid, filename, dataSourceName)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, err
}

