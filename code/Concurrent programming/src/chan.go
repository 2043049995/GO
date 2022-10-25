package src

import "fmt"

func Chan01() {
	var notice chan string = make(chan string) //管道里面需要放元素，定义一个字符串的管道,管道要用make进行初始化后才能使用,不初始化其地址会为空
	//fmt.Printf("%T  %v", notice, notice)
	var channel chan int = make(chan int, 4) //带缓冲区的管道
	fmt.Println(len(channel))
	close(channel) //管道关闭后只能读数据，不能写数据
	//使用for range 读取管道内数据的话必须有一个例程能将管道关闭，否则会一直读导致死锁
	
	//notice <- "一个字符串" //向管道里写数据
	//fmt.Println(<-notice)  //从管道里读数据，如果管道内没数据还要一直读的话就会出现阻塞，无缓冲区管道，需要在例程里写数据
	go func() {
		fmt.Println("go start")
		notice <- "一个字符串"
		fmt.Println("end start")
	}()
	fmt.Println("start")
	fmt.Println(<-notice)
	fmt.Println("end")
	for i := 0; i < 10; i++ {
		<-notice ///从管道内读10次，如果读不到数据会被阻塞
	}
}
