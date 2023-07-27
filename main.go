package main

import (
	"Golang/channel"
	"net/http"
)

func main() {
	http.HandleFunc("/", channel.OverTimeControl)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
