package main

import (
	"fmt"
	"net/http"
)

var a = "测试"

func main() {
	b := fmt.Sprintf("这回一个%s", a)
	fmt.Printf("%T", b)
	http.Client{}

}
