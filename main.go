package main

import (
	"github.com/liuzhangpei/alioss/server"
	"github.com/golang/glog"
	"fmt"
)


func main() {
	c, err := server.LoadConfig("/Users/lzp/Desktop/WorkGo/src/github.com/liuzhangpei/alioss/conf/key.yml")

	if err != nil {
		glog.Exitf("Failed to load config: %s", err)
	}

	fmt.Println(c.AliyunKey.AccessKeySecret)
}
