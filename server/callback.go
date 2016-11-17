package server

import (
	//"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"io/ioutil"
	"encoding/base64"
)





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
	response, _ := client.Do(request)
	defer response.Body.Close()

	if response.StatusCode == 200 {
		str, _ := ioutil.ReadAll(response.Body)
		bodystr := string(str)
		fmt.Println(bodystr)
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
	fmt.Printf("authorization : %s", string(authorization))


	// get callback body
	content_length := r.Header.Get("content-length")
	fmt.Println("content_length")
	fmt.Println(content_length)
	//content_length := r.Header["content-length"]
	//callback_body := r.Body.Read()
	//// callback_body = self.rfile.read(int(content_length))
	//
	//
	//// compose authorization string
	//pos := r.URL





      //  #
      //  auth_str = ''
      //  pos = self.path.find('?')
      //  if -1 == pos:
      //      auth_str = self.path + '\n' + callback_body
      //  else:
      //      auth_str = urllib2.unquote(self.path[0:pos]) + self.path[pos:] + '\n' + callback_body
      //  print auth_str
	 //
      //  #verify authorization
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