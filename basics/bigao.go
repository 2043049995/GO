package main

import "fmt"

func main() {

	name := "李朝阳"

	func() {
		fmt.Println("名字是：", name)
		func() {
			fmt.Println("内层名字是：", name)
		}()
	}()

	add2 := func(a int) int {
		return a + 2
	}
	fmt.Println(add2(2))
}
