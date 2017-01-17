package utils


import (
	"fmt"
	"github.com/golang/glog"
)

var ActionTypeDict map[string]interface{}


func (cat *CallbackActionType)UserAvatars(dataSourceName, HostOuter string) (retbool bool, err error) {
	// 用户头像

	//filename := HostOuter + "/" + cat.Filename
	filename := cat.Filename  // 不添加前缀


	glog.V(2).Infof("LoadConfig: %s", filename)

	err = UpdateUserAvatars(cat.Uuid, filename, dataSourceName)
	if err != nil {
		glog.V(2).Infof("UserAvatars: err!= nil : %s", err.Error())
		return false, err
	}

	return true, err
}



func (cat *CallbackActionType)MirrorIcon(dataSourceName, HostOuter string) (retbool bool, err error) {
	// 镜像图标
	filename := cat.Filename

	fmt.Println(filename)
	err = UpdateMirrorIcon(cat.Uuid, cat.ActionResourceId, filename, dataSourceName)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, err
}
