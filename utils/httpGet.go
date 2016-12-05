package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetPublicKey(pub_key_url string) (retbool bool, public_key []byte) {
	var client = &http.Client{}

	fmt.Println("begin get")
	fmt.Println(pub_key_url)
	request, err := http.NewRequest("GET", pub_key_url, nil)
	if err != nil {
		fmt.Println(" http.NewRequest err != nil")
		return false, nil
	}
	fmt.Println("begin 02 get")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(" client.Do(request) err != nil")
		fmt.Println(err.Error())
		return false, nil
	}

	fmt.Println("begin 03 get")
	defer response.Body.Close()

	if err != nil {
		fmt.Println("err != nil")
		return false, nil
	}

	if response.StatusCode == 200 {
		fmt.Println("response.StatusCode")
		public_key, _ = ioutil.ReadAll(response.Body)
		public_key_str := string(public_key)
		fmt.Println(public_key_str)
		return true, public_key
	} else {
		return false, nil
	}
}



func GetPublicKeyTwo(pub_key_url string) (retbool bool, public_key []byte) {


	fmt.Println("url.Parse(pub_key_url)")
	fmt.Println(pub_key_url)

	fmt.Println("---GetPublicKeyTwo--- 01 0")
	res, err := http.Get(pub_key_url)

	fmt.Println("---GetPublicKeyTwo--- 01 2")
	if err != nil {
		fmt.Println(err.Error())
		return false, nil
	}

	result, err := ioutil.ReadAll(res.Body)

	fmt.Println("-- ioutil.ReadAll --")
	if err != nil {
		fmt.Println(err.Error())
		return false, nil
	}
	fmt.Println("-- ioutil.ReadAll -- 02")
	res.Body.Close()
	return true, result
}
