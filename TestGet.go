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

	fmt.Println("HttpGet 00")
	res, err := http.Get("https://gosspublic.alicdn.com/callback_pub_key_v1.pem")
	fmt.Println("HttpGet 001")



	if err != nil {
		fmt.Println("HttpGet 02")
		log.Fatal(err)
		return
	}

	defer res.Body.Close()
	fmt.Println("HttpGet 01")

	fmt.Println("HttpGet 03")

	result, err := ioutil.ReadAll(res.Body)

	fmt.Println("HttpGet 04")

	if err != nil {
		fmt.Println("HttpGet 05")
		log.Fatal(err)
		return
	}
	fmt.Println("%v", string(result))
}

func main() {

	for i := 0; i < 30; i ++ {
		go HttpGet()
	}

   ch1 := make(chan int)

    <- ch1

}
