package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"io/ioutil"
)




func Callback(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		return
	}

	// get public key; 如果无法取得  public key 这里需要返回,不可以继续执行
	pub_key_url_base64 := r.Header["x-oss-pub-key-url"]
	pub_key_url := base64.StdEncoding.EncodeToString(pub_key_url_base64)

	var client = &http.Client{}

	request, _ := http.NewRequest("GET", pub_key_url, nil)
	response, _ := client.Do(request)
	defer response.Body.Close()

	if response.StatusCode == 200 {
		str, _ := ioutil.ReadAll(response.Body)
		bodystr := string(str)
		fmt.Println(bodystr)
	}
	// get public key; 如果无法取得  public key 这里需要返回,不可以继续执行


	// get authorization
	authorization_base64 := r.Header["authorization"] // Authorization
	authorization := base64.StdEncoding.EncodeToString(authorization_base64)

	fmt.Printf("authorization : %s", authorization)

	// get callback body
	content_length := r.Header["content-length"]
	callback_body := r.Body.Read()
	// callback_body = self.rfile.read(int(content_length))


	// compose authorization string
	pos := r.URL





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
	w.Header().Set()



	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", )
	io.WriteString(w, string(response_oss))

}


type ResponseOss struct {
	Status string `json:"Status"`
}