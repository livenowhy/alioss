package server

import (
	"net/http"
	"fmt"
	"io"
	"encoding/json"
	"io/ioutil"
	"github.com/liuzhangpei/alioss/utils"
)


// 获取
func (cg *Config)PolicyCallback(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的

	w.Header().Set("Access-Control-Allow-Headers", "token")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Println("---->")
	headtoken := r.Header.Get("token")
	fmt.Println(headtoken)

	//结构已知，解析到结构体
	result, err:= ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll is error")
	}

	fmt.Println(string(result))
    var actionType utils.CallbackActionType;
	err = json.Unmarshal([]byte(result), &actionType)

	//err = json.NewDecoder(r.Body).Decode(&actionType)
	if err != nil {

		fmt.Println(err.Error())

		fmt.Println("json.Unmarshal is error")
	}

	fmt.Println("sidsdsdsdsds")

	fmt.Println(actionType.ActionResourceId)
	fmt.Println(actionType.ActionType)
	fmt.Println(actionType.Uuid)

	if headtoken == "" {
		fmt.Println("actionT.ActionIcon() is error")
		utils.ResponseError(w, "ERROR", "token is nill")
		return

	}
	_ , err, vt := cg.MysqlConf.CheckToken(headtoken)
	if err != nil {
		fmt.Println("actionT.ActionIcon() is error")
		utils.ResponseError(w, "ERROR", "token is error")
		return
	}

	actionType.Uuid = vt.User_id

	fmt.Println(" toke is ok")
	response := cg.AliyunKey.GetPolicyToken("user-dir/", &actionType)
	fmt.Println("response end")



	io.WriteString(w, response)
}



