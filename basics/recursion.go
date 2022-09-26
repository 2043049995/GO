package main

import "fmt"

func main() {
	fmt.Println(Factorial(5))
}
func Factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func Addn(n int) int {
	if n == 1 {
		return 1
	}
	return n + Addn(n-1)
}
func sum() {
	var sum int = 0
	for i := 0; i <= 100; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
}
