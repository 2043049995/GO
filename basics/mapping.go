package main

import "fmt"

func main() {
	//映射是存储一系列无序的key/value对，通过key对value进行访问,map为空的时候是不能进行任何操作的,必须将值初始化后才能进行操作
	var score map[string]int
	fmt.Printf("%T", score)
	fmt.Println(score == nil)

}
