package main

import "fmt"

func main() {
	var slice1 = []int{1,2,3}
	a, b ,c := slicess(slice1...)
	fmt.Println(a, b)
}
func work1() (int, int) {
	return 1, 2

}
func slicess(a []int) (int,int,int){
	return (a...)

}