package main

import (
    "encoding/base64"
    "fmt"
	"strings"
	//"net/http/httputil"
	"golang.org/x/net/dict"
	"html"
)


func main() {

	//
	//	// aHR0cHM6Ly9nb3NzcHVibGljLmFsaWNkbi5jb20vY2FsbGJhY2tfcHViX2tleV92MS5wZW0=

	dst := "https://gosspublic.alicdn.com/callback_pub_key_v1.pem"
	dst_b := []byte(dst)
	base_str := base64.StdEncoding.EncodeToString(dst_b)
    fmt.Println(string(base_str))



	str := "aHR0cHM6Ly9nb3NzcHVibGljLmFsaWNkbi5jb20vY2FsbGJhY2tfcHViX2tleV92MS5wZW0="
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	path := "/swd=golang%20sds?sss"
	callback_body := "callback_body"


    pos := strings.Index(path, "?")

	if -1 == pos {
		auth_str := path + "\n" + callback_body
		fmt.Println(auth_str)
	} else {

		fmt.Println(html.UnescapeString(path[0:pos]))
		auth_str := dict.UnescapeString(path[0:pos]) + path[pos:] + "\n" + callback_body
        fmt.Println(auth_str)
		//auth_str = urllib2.unquote(path[0:pos]) + path[pos:] + '\n' + callback_body
	}


	fmt.Println(pos)
    //callback_body := "callback_body"


	fmt.Println(string(data))
}


