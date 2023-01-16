package main

import (
	"fmt"
	"net/http"
)

func handler(web http.ResponseWriter, request *http.Reques){
	fmt.Fprintf(web, "konichiwa !!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
