package src

import (
	"fmt"
	"runtime"
	"sync"
)

// 一般例程函数会写成匿名的
// 如果使用闭包的话遇到变化的值会影响函数的执行要通过参数的方式将值传递进去
func Goroutine2() {
	Group1 := &sync.WaitGroup{}
	Group1.Add(2) //决定了里面有几个信号量，意思是这里面的两个信号量没有出现的时候wait函数一直是等待状态，当出现后wait就结束等待。此处只是定义出了信号量的数目
	//并行：同时运行。
	//并发，同时开始运行，提高系统运行效率
	//go语言里面启动一个并发编程就是启动一个例程
	//main函数叫做主例程，其它函数叫做工作例程
	//var i = 'A'用单引号返回的是码点，当成字符处理，双引号会被当成字符串处理。
	//var j = "A"
	//在主函数中写的为主例程，用go启动的函数为工作例程
	go func() {
		for i := 'A'; i <= 'Z'; i++ {
			fmt.Println("1", string(i))
			runtime.Gosched()
		}
		Group1.Done()
	}()
	go func() {
		for i := 'A'; i <= 'Z'; i++ {
			fmt.Println("2", string(i))
			runtime.Gosched()
		}
		Group1.Done()
	}()
	Group1.Wait()
}
