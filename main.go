package main

import (
	"github.com/liuzhangpei/alioss/server"
	"net/http"
	"github.com/golang/glog"
	"fmt"
)

func main() {

	c, err := server.LoadConfig("./conf/key.yml")


	if err != nil {
		glog.Exitf("init to load config: %s", err)
	}
	fmt.Println("init_config")
	fmt.Println(c.AliyunKey.AccessKeySecret)



	http.HandleFunc("/policy", c.PolicyCallback)
	http.HandleFunc("/callback", c.Callback)
	http.ListenAndServe(":8765", nil)
}
