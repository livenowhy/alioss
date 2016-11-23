package main

import (
	"fmt"
	//"net/url"
	"net/http"
	"io/ioutil"
	"log"
)

func HttpGet()  {
	//u, _ := url.Parse("https://gosspublic.alicdn.com/callback_pub_key_v1.pem")
	res, err := http.Get("https://gosspublic.alicdn.com/callback_pub_key_v1.pem")

	if err != nil {
		log.Fatal(err)
		return
	}

	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("%v", string(result))
}

func main() {

	for i := 0; i < 30; i ++ {
		HttpGet()
	}

}
