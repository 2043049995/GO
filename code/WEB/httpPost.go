package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		io.WriteString(resp, req.Method)
		fmt.Println(req.PostFormValue("a"))
	})
	http.ListenAndServe(":8889", nil)

}
