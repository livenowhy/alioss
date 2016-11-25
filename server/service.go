package server

import (
	"net/http"
	"fmt"
	"io"
	"encoding/json"
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

	w.Header().Set("Access-Control-Allow-Headers", "token")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Println("---->")
	headtoken := r.Header.Get("token")
	fmt.Println(headtoken)

	if headtoken == "" {
		response := error_response(2, "token is nil")
		io.WriteString(w, response)
		return
	}
	_, err := cg.MysqlConf.CheckToken(headtoken)
	if err != nil {
		response := error_response(2, "token is error")
		io.WriteString(w, response)
		return
	}

	fmt.Println(" toke is ok")


	response := cg.AliyunKey.GetPolicyToken("user-dir/")

	fmt.Println("response end")

	//response := "sds"



	io.WriteString(w, response)
}



