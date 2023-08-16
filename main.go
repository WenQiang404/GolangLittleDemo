package main

import (
	"Golang/channel"
	"Golang/usual"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", channel.OverTimeControl)
	http.ListenAndServe("127.0.0.1:8080", nil)
	fmt.Print(usual.ConvertBin(8))
}
