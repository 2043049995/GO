package main

import "fmt"

func main() {
	//数组里面存放数组类型数据
	var marr [3][5]int
	fmt.Println(marr)
	fmt.Printf("%T", marr)
	//多维数组只有第一维可以推倒长度，数组的类型必须是一致的
}
