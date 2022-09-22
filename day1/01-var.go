package main

import "fmt"

var version string = "v1.0"
var (
	height int = 10
	weight int = 20
)

// 函数外定义的变量可以不使用，函数内部的变量必须被使用
// 如果函数定义后未赋值会默认赋零值
func main() {
	var 李朝阳 string = "名字"
	var li = "学号"
	a := 10 //简短声明只能在函数内声明
	var name, user string = "李朝阳", "li"
	fmt.Println(李朝阳)
	fmt.Println(li)
	fmt.Println(a)
	fmt.Println(weight, height)
	fmt.Println(version)
	fmt.Println("nameIs:"+name, "userIs:", user)
}
