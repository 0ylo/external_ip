package main

import (
	"fmt"
	"net/http"
)

func main() {
	handleRequest()
}

func handleRequest() {
	http.HandleFunc("/", ip)
	http.ListenAndServe(":8080", nil)
}

func ip(w http.ResponseWriter, req *http.Request) {
	fmt.Println("IP adress:\n", req.RemoteAddr)

}
