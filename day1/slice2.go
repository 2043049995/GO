package main

import "fmt"

func main() {
	var slice1 = []int{}
	var s2 []int
	fmt.Printf("%T", slice1)
	fmt.Printf("%T", s2)
	slice2 := []int{1, 2, 3}
	slice4 := slice2
	slice4[0] = 100
	slice5 := append(slice4, 60)

	fmt.Println(slice2)
	fmt.Println(slice5)
	slices := []int{}
	arr:=[2]int{1,2]}
	fmt.Println(slices == nil)
	fmt.Println(arr[0] == nil)

}
