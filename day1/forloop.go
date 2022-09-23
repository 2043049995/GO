package main

import "fmt"

func main() {
	result := 0
	for i := 1; i <= 100; i++ {
		result = result + i
	}
	fmt.Println(result)
}
