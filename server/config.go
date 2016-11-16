/*
   Copyright 2015 Cesanta Software Ltd.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       https://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package server

import (
	"fmt"
	"io/ioutil"
	yaml "gopkg.in/yaml.v2"
	"github.com/golang/glog"
	"github.com/liuzhangpei/alioss/aliyun"
)


type Config struct {
	AliyunKey aliyun.AliYunAccessKey   `yaml:"aliyunkey,omitempty"`
	AliyunOss aliyun.AliYunOssConf   `yaml:"oss,omitempty"`
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

	return c, nil
}
