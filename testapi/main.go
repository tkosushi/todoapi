package main

import (
	"fmt"
	"net/http"
)

func handler(web http.ResponseWriter, request *http.Request){
	header := request
	fmt.Fprintln(web, header)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
