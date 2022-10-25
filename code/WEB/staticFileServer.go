package main

import "net/http"

// 静态文件服务器
func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))

	http.ListenAndServe(":8889", nil)
}
