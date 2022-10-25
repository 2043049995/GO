package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		io.WriteString(resp, req.Method)
		req.ParseMultipartForm(1024 * 1024)
		file, _ := req.MultipartForm.File["a"][0].Open()
		a, _ := io.Copy(os.Stdout, file)
		fmt.Println(a)
	})
	http.ListenAndServe(":8889", nil)

}
