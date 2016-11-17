package main

import (
	"github.com/liuzhangpei/alioss/server"
	"github.com/golang/glog"
	"fmt"
	"net/http"

)


func main() {
	//c, err := server.LoadConfig("/Users/lzp/Desktop/WorkGo/src/github.com/liuzhangpei/alioss/conf/key.yml")
	//
	//if err != nil {
	//	glog.Exitf("Failed to load config: %s", err)
	//}
		// 106.38.76.170
	//fmt.Println(c.AliyunKey.AccessKeySecret)

	http.HandleFunc("/", server.PolicyCallback)
	http.HandleFunc("/callback", server.Callback)
	http.ListenAndServe(":8765", nil)


}
