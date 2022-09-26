package main

import "fmt"

func main() {
	//go里面没有提供异常的机制，抛出运行时的错误需要用panic抛出，panic抛出后，go允许在延迟函数里使用recover去终止错误的处理
	//不终止错误的话程序就会退出。recover一般在defer函数里去应用
	defer func() { //一般用来做发生错误的恢复处理
		if err := recover(); err != nil {
			fmt.Println(err)

		}
		//fmt.Println(recover()) //recover一般用来做恢复，主动处理panic抛出的错误.默认会返回panic抛出的错误信息
	}()
	fmt.Println("main satrt")
	//panic("error") //在panic抛出错误后程序会执行结束，剩下的不会再执行
	//recover一般用来做恢复，主动处理panic抛出的错误

	fmt.Println("over")

}
