package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Println(req.URL)

	})
	http.ListenAndServe(":8889", nil)
}
