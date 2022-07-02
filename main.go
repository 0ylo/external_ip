package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", ip)
	http.ListenAndServe(":8080", nil)
}

func ip(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(req.RemoteAddr))
	return
}
