package main

import (
	"fmt"
	//"net/url"
	"net/http"
	"io/ioutil"
	"log"
	"bytes"
)

func HttpGet()  {
	//u, _ := url.Parse("https://gosspublic.alicdn.com/callback_pub_key_v1.pem")
	//res, err := http.Get("https://gosspublic.alicdn.com/callback_pub_key_v1.pem")
	res, err := http.Get("https://www.baidu.com/")
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

func TestGet()  {

	host := "https://gosspublic.alicdn.com/callback_pub_key_v1.pem"

	var b []byte
	req,err := http.NewRequest("GET",host, bytes.NewBuffer(b))
	if err != nil {
		fmt.Printf("make request error: %v",err)
		} else {
			client := &http.Client{}
			res,err := client.Do(req)
			defer res.Body.Close()
			if err != nil {
				fmt.Printf("send log reponse error: %v",err)
			} else {
				body,err := ioutil.ReadAll(res.Body)
				if err != nil {
					fmt.Printf("send log read response body error: %v",err)
				}
				fmt.Printf("send log response: %s", string(body))
			}
		}
}


func main() {

	for i := 0; i < 30; i ++ {
		go TestGet()
	}

   ch1 := make(chan int)

    <- ch1

}
