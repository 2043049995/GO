package main

import (
	"fmt"
	"net/http"
	"strings"
)

func CookieParse(cookie string) map[string]string {
	cookieMap := make(map[string]string)
	values := strings.Split(cookie, ";")
	for _, value := range values {
		es := strings.Split(value, "=")
		cookieMap[strings.TrimSpace(es[0])] = strings.TrimSpace(es[1])
	}
	return cookieMap

}
func main() {

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Println(CookieParse(req.Header.Get("Cookie")))
		counterCookie := &http.Cookie{
			Name:     "counter",
			Value:    "1",
			HttpOnly: true,
		}
		http.SetCookie(resp, counterCookie)
	})
	http.ListenAndServe(":8889", nil)
}
