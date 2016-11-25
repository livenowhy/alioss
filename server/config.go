
package server

import (
	"fmt"
	"io/ioutil"
	yaml "gopkg.in/yaml.v2"
	"github.com/golang/glog"
	"github.com/liuzhangpei/alioss/aliyun"
	"github.com/liuzhangpei/alioss/token"
)


type Config struct {
	AliyunKey aliyun.AliYunAccessKey   `yaml:"aliyunkey,omitempty"`
	AliyunOss aliyun.AliYunOssConf   `yaml:"oss,omitempty"`
	MysqlConf  token.MysqlConfig   `yaml:"mysqldbconf,omitempty"`

}



// 加载配置文件数据,
func LoadConfig(fileName string) (*Config, error) {
	glog.V(2).Infof("LoadConfig: %s", fileName)

	fmt.Println(fileName)
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("could not read %s: %s", fileName, err)
	}

	fmt.Println("ssss")
	c := &Config{}
	glog.V(2).Infof("LoadConfig: %s", contents)
	if err = yaml.Unmarshal(contents, c); err != nil {
		return nil, fmt.Errorf("could not parse config: %s", err)
	}

	fmt.Println(c.AliyunKey.AccessKeyID)
	fmt.Println(c.AliyunKey.AccessKeySecret)

	fmt.Println(c.AliyunOss.HostOuter)
	fmt.Println(c.AliyunOss.HostIn)
	fmt.Println(c.AliyunOss.CallbackUrl)
	fmt.Println(c.AliyunOss.ExpireTime)
	fmt.Println(c.AliyunOss.UploadDir)
	fmt.Println(c.MysqlConf.Charset)


	return c, nil
}



// server 包下的全局配置变量

var CONF = &Config{}

func init() {
	c, err := LoadConfig("./conf/key.yml")
	if err != nil {
		glog.Exitf("init to load config: %s", err)
	}
	fmt.Println("init_config")
	fmt.Println(c.AliyunKey.AccessKeySecret)

	CONF = c
}