package main

import (
	"github.com/liuzhangpei/alioss/server"
	"net/http"
	"github.com/golang/glog"
)

func main() {

	c, err := server.LoadConfig("./conf/key.yml")


	if err != nil {
		glog.Exitf("init to load config: %s", err)
	}

	glog.V(2).Infof("init_config")
	glog.V(2).Infof("c.AliyunKey.AccessKeySecret: %", c.AliyunKey.AccessKeySecret)
	glog.V(2).Infof("c.AliyunKey.AccessKeyID: %", c.AliyunKey.AccessKeyID)

	glog.V(2).Infof("c.AliyunKey.CallbackUrl: %", c.AliyunKey.CallbackUrl)

	glog.V(2).Infof("c.AliyunKey.HostOuter: %", c.AliyunKey.HostOuter)

	glog.V(2).Infof("c.AliyunKey.HostIn: %", c.AliyunKey.HostIn)

	// http://img.boxlinker.com/test
	http.HandleFunc("/test", c.ServerTest)
	http.HandleFunc("/testu", c.ServerTest)
	http.HandleFunc("/policy", c.PolicyCallback)
	http.HandleFunc("/callback", c.Callback)
	http.ListenAndServe(":8765", nil)
}
