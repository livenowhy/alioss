package main

import (
	"github.com/liuzhangpei/alioss/server"
	"net/http"
	"github.com/golang/glog"
	"fmt"
	//"github.com/liuzhangpei/alioss/aliyun"
	//"github.com/liuzhangpei/alioss/aliyun"
	"github.com/liuzhangpei/alioss/aliyun"
)

func main() {

	c, err := server.LoadConfig("./conf/key.yml")


	if err != nil {
		glog.Exitf("init to load config: %s", err)
	}
	fmt.Println("init_config")
	fmt.Println(c.AliyunKey.AccessKeySecret)



	http.HandleFunc("/policy", c.PolicyCallback)
	http.HandleFunc("/callback", aliyun.Callback)
	http.ListenAndServe(":8765", nil)
}
