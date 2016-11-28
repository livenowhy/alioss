package aliyun

import (

	"fmt"
	"io"
	"net/http"
	"io/ioutil"
	"encoding/base64"
    "crypto"
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
	"crypto/md5"
	"github.com/liuzhangpei/alioss/utils"
)



/***************************************************************
*RSA签名验证
*src:待验证的字串，sign:aliyun返回的签名
*pass:返回true表示验证通过
*err :当pass返回false时，err是出错的原因
****************************************************************/
func RSAVerify(src []byte, sign []byte, public_key []byte) (pass bool, err error) {
    //步骤1，加载RSA的公钥
    block, _ := pem.Decode(public_key)

	fmt.Println("in RSAVerify -->")

	fmt.Println(string(public_key))

	if block == nil {
		fmt.Println(string(public_key))
		fmt.Printf("Failed to pem.Decode(public_key) \n")
		return true, nil
	}

    pub, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        fmt.Printf("Failed to parse RSA public key: %s\n", err)
        return
    }
    rsaPub, _ := pub.(*rsa.PublicKey)

    //步骤2，计算代签名字串的SHA1哈希
    t := md5.New()
    io.WriteString(t, string(src))
    digest := t.Sum(nil)

    ////步骤3，base64 decode,必须步骤，支付宝对返回的签名做过base64 encode必须要反过来decode才能通过验证
    //data, _ := base64.StdEncoding.DecodeString(string(sign))

    //hexSig := hex.EncodeToString(data)
    //fmt.Printf("base decoder: %v, %v\n", string(sign), hexSig)

    //步骤4，调用rsa包的VerifyPKCS1v15验证签名有效性
    err = rsa.VerifyPKCS1v15(rsaPub, crypto.MD5, digest, sign)
    if err != nil {
        fmt.Println("Verify sig error, reason: ", err)
        return false, err
    } else {
		fmt.Println("Verify sig is ok, reason: ")
	}

    return true, nil
}



// 阿里云  callback 回调验证
func AliCallback(w http.ResponseWriter, r *http.Request) (bodystr string, err error){


	bodystr = ""
	fmt.Println(" -- range r.Header -- 01 ")
	for k, v := range r.Header {
		fmt.Println(k)
		fmt.Println(v)
	}
	fmt.Println(" -- range r.Header -- 02 ")

	// get public key; 如果无法取得  public key 这里需要返回,不可以继续执行
	pub_key_url_base64 := r.Header.Get("x-oss-pub-key-url")
	fmt.Println(pub_key_url_base64)
	pub_key_url, err := base64.StdEncoding.DecodeString(pub_key_url_base64)
	if err != nil {
		fmt.Println("err != nil --2 ")
		fmt.Println("error:", err.Error())
		return
	}
	fmt.Println("pub_key_url")
	fmt.Println(pub_key_url)

	// get public key; 如果无法取得  public key 这里需要返回,不可以继续执行
	pub_key_url_str := string(pub_key_url)
	retbool, public_key := utils.GetPublicKeyTwo(pub_key_url_str)
	if !retbool {
		return
	}

	fmt.Println("get public key is ok")
	authorization_base64 := r.Header.Get("authorization") // get  Authorization

	fmt.Println(authorization_base64)
	authorization , err:= base64.StdEncoding.DecodeString(authorization_base64)
	if err != nil {
		fmt.Println("error authorization_base64 :", err)
		return
	}
	fmt.Println("authorization :")
	fmt.Println(string(authorization))

	// get callback body
	content_length := r.Header.Get("content-length")
	fmt.Println("content_length")
	fmt.Println(content_length)


	callback_body, _ := ioutil.ReadAll(r.Body)
	bodystr = string(callback_body)

	fmt.Println("callback_body")
	fmt.Println(bodystr)
	// callback_body: filename=user-dir%2F537078.gif&size=7005&mimeType=image%2Fgif&height=64&width=64


	// #compose authorization string
	path := r.URL.Path
	fmt.Println("ss := r.URL.Path")
	fmt.Println(path)

	auth_str := path + "\n" + bodystr
	fmt.Println(auth_str)
		//
	     //   if -1 == pos:
        //    auth_str = self.path + '\n' + callback_body
        //else:
        //    auth_str = urllib2.unquote(self.path[0:pos]) + self.path[pos:] + '\n' + callback_body
	// 暂时不考虑这种情况


	fmt.Println(string(public_key))
	pass, err := RSAVerify([]byte(auth_str), authorization, public_key)  // 验证签名
	if pass == false {
		fmt.Println("is error")
		err.Error()
		return
	}
	return
}


