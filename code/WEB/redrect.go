package main

import (
	"fmt"
	"net/http"
)

func main() {
	//重定向，让浏览器发起新的请求到新的地址
	http.HandleFunc("/home", func(resp http.ResponseWriter, req *http.Request) {
		if req.Header.Values("Cookie") == nil {
			http.Redirect(resp, req, "/login", 302) //设置没有登录就重定向到登录页
		} else {
			fmt.Fprintf(resp, "首页")
		}

	})
	http.HandleFunc("/login", func(resp http.ResponseWriter, req *http.Request) {
		cookies := &http.Cookie{
			Name:  "counter",
			Value: "1",
		}
		http.SetCookie(resp, cookies)
		fmt.Fprintf(resp, "登录页面")
	})
	http.ListenAndServe(":8889", nil)

}
