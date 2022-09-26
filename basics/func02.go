package main

import (
	"fmt"
)

func main() {
	//函数的参数、参数的数量、对应的类型、返回值统称为函数的签名
	fmt.Printf("%T", add01) //返回类型：func(int, int) int
	f := add01              //等同于如下
	//var f func(int, int) int
	fmt.Printf("%T\n", f) //返回类型：func(int, int) int
	//某个函数可以赋值给函数签名相同的变量
	//var f func(int, int) int = add
	fmt.Println(f(1, 5), "\n")

	sayHello := func(name string) {
		fmt.Println("hello", name)
	}

	sayHello("lizhaoyang")
	print(list, "A", "B", "C")
}

//匿名函数

// callback格式化，将传递的数据按照每行打印还是按照一行按|分割打印
func print(callback func(...string), args ...string) {
	fmt.Println("print函数的输出：")
	callback(args...)
}
func list(args ...string) {
	for i, v := range args {
		fmt.Println(i, ":", v)
	}
}

func add01(a, b int) int {
	return a + b
}
