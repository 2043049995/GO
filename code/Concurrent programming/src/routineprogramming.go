package src

import (
	"fmt"
	"runtime"
	"sync"
)

func Routin() {
	Group := &sync.WaitGroup{}
	Group.Add(2) //决定了里面有几个信号量，意思是这里面的两个信号量没有出现的时候wait函数一直是等待状态，当出现后wait就结束等待。此处只是定义出了信号量的数目
	//并行：同时运行。
	//并发，同时开始运行，提高系统运行效率
	//go语言里面启动一个并发编程就是启动一个例程
	//main函数叫做主例程，其它函数叫做工作例程
	//var i = 'A'用单引号返回的是码点，当成字符处理，双引号会被当成字符串处理。
	//var j = "A"
	//在主函数中写的为主例程，用go启动的函数为工作例程
	go PrintChars("1", Group) //加了go后已经变为了并发运行
	//go里面每个并发执行单元称为Goroutine,启动一个goroutine只需要在需要并发执行的函数前加上go关键词即可
	go PrintChars("2", Group) //前面加go的两个函数实际效果看着像没执行
	PrintChars("3", Group)
	//time.Sleep(time.Second * 3) //添加此项后会发现在3执行完后再执行2和1。加了runtime.GOshaed后删除此项后发现只要3打印完成后就会全部结束，1和2可能还未打印完
	//在go语言里主例程执行完毕后工作例程会自动结束，上面的time就是为了等其他进程一会，但是不知道等多长时间，
	//如何正确等待工作例程结束，1、通过计数信号量。
	Group.Wait() //当达到Add函数里需要的信号量数量时才会执行，其执行完毕后主例程结束，所有工作例程也结束
}

func PrintChars(name string, Group *sync.WaitGroup) {
	for i := 'A'; i <= 'Z'; i++ {
		fmt.Println(name, string(i))
		runtime.Gosched() //加上此行后就会发现此函数的执行顺序已经乱了，，执行完上面的语句后会让出CPU，已经变为了并发运行
		//GOsched的作用就是让出CPU,让出CPU后被哪个进程抢到哪个进程就执行，用sleep也会让出时间片

	}
	Group.Done() //产生一个信号量，在每个例程结束的时候为其生成一个信号量，

}
