package main

import "fmt"

func main() {

	//panic与延迟执行
	defer func() { //一般使用匿名函数调用，当其所在的函数退出的时候执行
		fmt.Println("defer")
	}()
	defer func() {
		fmt.Println("defer02") //会先执行，符合先入后出的栈结构
	}()
	fmt.Println("main over") //main会先于defer打印出来

}
