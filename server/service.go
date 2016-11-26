package server

import (
	"net/http"
	"fmt"
	"io"
	"encoding/json"
	"io/ioutil"
	"github.com/liuzhangpei/alioss/utils"
)


type StatusMsg struct {
	StatusCode int     `json:"statuscode"`  // add lzp 添加验证token操作
	ErrMsg string     `json:"errmsg"`  // add lzp 添加验证token操作
}



func error_response(statuscode int, msg string) string{
	var statusmsg StatusMsg
	statusmsg.StatusCode = statuscode
	statusmsg.ErrMsg = msg
	response, err := json.Marshal(statusmsg)
	if err != nil {
		fmt.Println("json err:", err)
	}
	return string(response)
}

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



	if headtoken == "" {
		response := error_response(2, "token is nil")
		io.WriteString(w, response)
		return
	}
	_ , err = cg.MysqlConf.CheckToken(headtoken)
	if err != nil {
		response := error_response(2, "token is error")
		io.WriteString(w, response)
		return
	}

	fmt.Println(" toke is ok")

	response := cg.AliyunKey.GetPolicyToken("user-dir/", &actionType)

	fmt.Println("response end")

	//response := "sds"



	io.WriteString(w, response)
}



