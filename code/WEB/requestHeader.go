package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		io.WriteString(resp, req.Method)
		//req.ParseForm()
		//fmt.Println(req.Form)
		//fmt.Println(req.Form["a"])
		fmt.Println(req.FormValue("a"))
	})
	http.ListenAndServe(":8889", nil)

}
