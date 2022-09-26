package main

import (
	"fmt"
	"sort"
)

func main() {
	slice1 := []int{1, 2, 3, 7, 8, 3, 10, 5, 200}
	a := sort.SearchInts(slice1, 2)
	fmt.Println(a)
}
