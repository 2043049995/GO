package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, _ := http.Get("http://www.baidu.com")
	fmt.Println(io.Copy(os.Stdout, resp.Body))
}
