package server


import (
	"net/http"
	"github.com/liuzhangpei/alioss/aliyun"
	"github.com/liuzhangpei/alioss/utils"
	"fmt"
)


func (cg *Config) Callback(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		return
	}

	bodystr, err := aliyun.AliCallback(w, r)

	if err != nil {
		utils.ResponseError(w, "ERROR", err.Error())
		return
	}

	actionT, err := utils.NewCallbackActionType(bodystr)

	if err != nil || actionT.ActionType == "" {
		utils.ResponseError(w, "ERROR", err.Error())
		return
	}

	fmt.Println("actionT.ActionIcon(")

	retbool, err := actionT.ActionIcon(cg.MysqlConf.DataSourceName, cg.AliyunKey.HostOuter)
	if !retbool {
		fmt.Println("actionT.ActionIcon() is error")
		utils.ResponseError(w, "ERROR", err.Error())
		return
	}

	utils.ResponseError(w, "OK", "is ok")
}
