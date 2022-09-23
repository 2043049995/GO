package main

import "fmt"

func main() {
	var a = 1
	var b = &a

	fmt.Println(b)
	fmt.Println(*b)
}
