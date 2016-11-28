/**
 * 1. 用户向应用服务器取到上传policy和回调设置
 * 2. 应用服务器返回上传policy和回调
 */

package aliyun

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash"
	"io"
	"time"
	"github.com/liuzhangpei/alioss/utils"

)


const (
	base64Table = "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
)

var coder = base64.NewEncoding(base64Table)

func base64Encode(src []byte) []byte {
	return []byte(coder.EncodeToString(src))
}

func get_gmt_iso8601(expire_end int64) string {
	var tokenExpire = time.Unix(expire_end, 0).Format("2006-01-02T15:04:05Z")
	return tokenExpire
}

type ConfigStruct struct {
	Expiration string     `json:"expiration"`
	Conditions [][]string `json:"conditions"`
}

type PolicyToken struct {
	StatusCode int     `json:"statuscode"`  // add lzp 添加验证token操作
	ErrMsg string     `json:"errmsg"`  // add lzp 添加验证token操作
	AccessKeyId string `json:"accessid"`
	Host        string `json:"host"`
	Expire      int64  `json:"expire"`
	Signature   string `json:"signature"`
	Policy      string `json:"policy"`
	Directory   string `json:"dir"`
	Callback    string `json:"callback"`
}

type CallbackParam struct {
	CallbackAction   *utils.CallbackActionType   `json:"callbackAction"`  // add lzp 上传动作(用户头像,镜像头像)
	CallbackUrl      string `json:"callbackUrl"`
	CallbackBody     string `json:"callbackBody"`
	CallbackBodyType string `json:"callbackBodyType"`
}


// 阿里云访问秘钥
type AliYunAccessKey struct {
	AccessKeyID string `yaml:"AccessKeyID,omitempty"`
	AccessKeySecret string `yaml:"AccessKeySecret,omitempty"`
	HostOuter string `yaml:"HostOuter,omitempty"`   // 外网访问地址
	HostIn string `yaml:"HostIn,omitempty"`   // 内网访问地址
	CallbackUrl string `yaml:"CallbackUrl,omitempty"`   // oss 回调
	ExpireTime int64 `yaml:"ExpireTime,omitempty"`
}

type AliYunOssConf struct {

	UploadDir string `yaml:"UploadDir,omitempty"`
}


// 阿里云获取授权码和回调地址
func (aly *AliYunAccessKey)GetPolicyToken(dir string, action *utils.CallbackActionType) string {
	now := time.Now().Unix()

	fmt.Println("ONF.AliyunOss.ExpireTime")
	fmt.Println(aly.ExpireTime)
	expire_end := now + aly.ExpireTime
	var tokenExpire = get_gmt_iso8601(expire_end)

	//create post policy json
	var config ConfigStruct
	config.Expiration = tokenExpire
	var condition []string
	condition = append(condition, "starts-with")
	condition = append(condition, "$key")
	condition = append(condition, dir)
	config.Conditions = append(config.Conditions, condition)

	//calucate signature
	result, err := json.Marshal(config)
	debyte := base64.StdEncoding.EncodeToString(result)
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(aly.AccessKeySecret))
	io.WriteString(h, debyte)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	var callbackParam CallbackParam
	callbackParam.CallbackUrl = aly.CallbackUrl
	// action 标明  操作的对象与资源; 拼接一个action
	actionStr := "/?&actionType=" + action.ActionType + "&actionResourceId=" + action.ActionResourceId + "&uuid=" + action.Uuid


	fmt.Printf(actionStr)

	callbackParam.CallbackBody = actionStr + "&filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}"
	callbackParam.CallbackBodyType = "application/x-www-form-urlencoded"
	callback_str, err := json.Marshal(callbackParam)
	if err != nil {
		fmt.Println("callback json err:", err)
	}
	callbackBase64 := base64.StdEncoding.EncodeToString(callback_str)

	var policyToken PolicyToken
	policyToken.StatusCode = 0
	policyToken.AccessKeyId = aly.AccessKeyID
	policyToken.Host = aly.HostOuter
	policyToken.Expire = expire_end
	policyToken.Signature = string(signedStr)
	policyToken.Directory = dir
	policyToken.Policy = string(debyte)
	policyToken.Callback = string(callbackBase64)
	response, err := json.Marshal(policyToken)
	if err != nil {
		fmt.Println("json err:", err)
	}
	return string(response)
}



