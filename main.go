package main

import (
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/", ip)
	http.ListenAndServe(":8080", nil)
}

func ip(w http.ResponseWriter, req *http.Request) {
	host, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return
	}
	w.Write([]byte(host))
	return
}
