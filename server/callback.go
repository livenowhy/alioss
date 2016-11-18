package server

import (
	//"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"io/ioutil"
	"encoding/base64"
	//"crypto/md5"
    "crypto"
    "crypto/rsa"
    "crypto/sha1"
    "crypto/x509"
    "encoding/hex"
    "encoding/pem"
)



/***************************************************************
*RSA签名验证
*src:待验证的字串，sign:支付宝返回的签名
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
    t := sha1.New()
    io.WriteString(t, string(src))
    digest := t.Sum(nil)

    ////步骤3，base64 decode,必须步骤，支付宝对返回的签名做过base64 encode必须要反过来decode才能通过验证
    //data, _ := base64.StdEncoding.DecodeString(string(sign))

    //hexSig := hex.EncodeToString(data)
    //fmt.Printf("base decoder: %v, %v\n", string(sign), hexSig)

    //步骤4，调用rsa包的VerifyPKCS1v15验证签名有效性
    err = rsa.VerifyPKCS1v15(rsaPub, crypto.SHA1, digest, sign)
    if err != nil {
        fmt.Println("Verify sig error, reason: ", err)
        return false, err
    }

    return true, nil
}



func Callback(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		return
	}

	// get public key; 如果无法取得  public key 这里需要返回,不可以继续执行
	pub_key_url_base64 := r.Header.Get("x-oss-pub-key-url")
	fmt.Println(pub_key_url_base64)
	// aHR0cHM6Ly9nb3NzcHVibGljLmFsaWNkbi5jb20vY2FsbGJhY2tfcHViX2tleV92MS5wZW0=
	// https://gosspublic.alicdn.com/callback_pub_key_v1.pem
	pub_key_url, err := base64.StdEncoding.DecodeString(pub_key_url_base64)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(pub_key_url)

	var client = &http.Client{}
	pub_key_url_str := string(pub_key_url)

	fmt.Println(pub_key_url_str)
	request, _ := http.NewRequest("GET", pub_key_url_str, nil)
	response, err := client.Do(request)
	defer response.Body.Close()

	if err != nil {
		fmt.Println("err != nil")
		return
	}

	var public_key []byte
	if response.StatusCode == 200 {
		fmt.Println("response.StatusCode")
		public_key, _ = ioutil.ReadAll(response.Body)
		public_key_str := string(public_key)
		fmt.Println(public_key_str)
	} else {
		return
	}

	fmt.Println("get public key is ok")
	// get public key; 如果无法取得  public key 这里需要返回,不可以继续执行


	// get authorization
	authorization_base64 := r.Header.Get("authorization") // Authorization

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
	bodystr := string(callback_body)

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
		//


	// 验证签名

	fmt.Println("验证签名")

	fmt.Println(string(public_key))
	pass, err := RSAVerify([]byte(auth_str), authorization, public_key)
	if pass == false {
		fmt.Println("is error")
		err.Error()
	}

	// RSAVerify(src []byte, sign []byte, public_key string) (pass bool, err error)


      //  auth_md5 = md5.new(auth_str).digest()
      //  bio = BIO.MemoryBuffer(pub_key)
      //  rsa_pub = RSA.load_pub_key_bio(bio)
      //  try:
      //      result = rsa_pub.verify(auth_md5, authorization, 'md5')
      //  except e:
      //      result = False
	 //
      //  if not result:
      //      print 'Authorization verify failed!'
      //      print 'Public key : %s' % (pub_key)
      //      print 'Auth string : %s' % (auth_str)
      //      self.send_response(400)
      //      self.end_headers()
      //      return
	 //
      //  #do something accoding to callback_body
	 //

      //  self.send_response(200)

      //  self.send_header('', str(len(resp_body)))

      //  self.wfile.write(resp_body)


	// response to OSS
	var ResponseOss ResponseOss
	ResponseOss.Status = "OK"

	response_oss, err := json.Marshal(ResponseOss)
	if err != nil {
		fmt.Println("json err:", err)
	}
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Content-Length", )
	io.WriteString(w, string(response_oss))
}


type ResponseOss struct {
	Status string `json:"Status"`
}