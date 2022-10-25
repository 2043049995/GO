package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		//io.Copy(os.Stdout, req.Body)
		decoder := json.NewDecoder(req.Body)
		var info map[string]interface{}
		decoder.Decode(&info)
		fmt.Println(info)
	})

	http.ListenAndServe(":8889", nil)
}
