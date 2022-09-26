package main

import (
	"fmt"
)

func main() {

	var maps = map[string]int{"李朝阳": 666, "李钊": 777}
	for k, _ := range maps {
		fmt.Println(k)

	}
}
