package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, _ := http.Get("http://www.baidu.com")

	a, _ := io.Copy(os.Stdout, resp.Body)
	fmt.Println(a)

}
