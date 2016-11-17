package main

import (
	"github.com/liuzhangpei/alioss/server"
	"github.com/golang/glog"
	"fmt"
	"net/http"

)


func main() {
	c, err := server.LoadConfig("/Users/lzp/Desktop/WorkGo/src/github.com/liuzhangpei/alioss/conf/key.yml")

	if err != nil {
		glog.Exitf("Failed to load config: %s", err)
	}

	http.HandleFunc("/", server.PolicyCallback)
	http.ListenAndServe(":1234", nil)
	fmt.Println(c.AliyunKey.AccessKeySecret)
}
