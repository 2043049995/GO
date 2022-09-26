package main

import "fmt"

func main() {

	a := func(name string) {
		fmt.Println("牛逼666", name)
	}
	a("李朝阳")
	//定义匿名函数并且调用，延迟处理会用到
	func(name string) {
		fmt.Println("666", name)
	}("李钊")
	//调用函数的时候定义匿名函数

}
func print01(callback01 func(...string), arge01 ...string) {
	fmt.Println("1")
}
