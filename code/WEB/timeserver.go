package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	timefunc := func(response http.ResponseWriter, request *http.Request) {
		fmt.Println(request)
		io.WriteString(response, time.Now().Format("2006-01-02 15:04:05"))
	}

	http.HandleFunc("/path", timefunc)
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {

		file, _ := os.Open("./index.html")
		defer file.Close()
		reader := bufio.NewReader(file)
		text, _ := reader.ReadString('\n')
		io.WriteString(response, text)
	})
	http.ListenAndServe("0.0.0.0:8889", nil)
}
