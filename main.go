package main

import (
	"github.com/liuzhangpei/alioss/server"
	"net/http"
)



func main() {
	http.HandleFunc("/policy", server.PolicyCallback)
	http.HandleFunc("/callback", server.Callback)
	http.ListenAndServe(":8765", nil)
}
