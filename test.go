package main

import (
    "encoding/base64"
    "fmt"
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


	fmt.Println(string(data))
}


