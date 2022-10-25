package src

import (
	"fmt"
	"runtime"
	"sync"
)

func Count1() {
	lock := &sync.Mutex{} //定义一个锁
	var counter int
	waitGroup := &sync.WaitGroup{}

	incr := func() {
		for i := 0; i <= 100; i++ {
			lock.Lock()   //获取锁（互斥锁）
			counter++     //拿count，count+1.写回count
			lock.Unlock() //释放锁
			runtime.Gosched()
		}
		waitGroup.Done()
	}
	decr := func() {
		for i := 0; i <= 100; i++ {
			lock.Lock() //获取不到锁的话后面的代码不会执行
			counter--
			lock.Unlock()
			runtime.Gosched()
		}
		waitGroup.Done()
	}
	for i := 0; i < 10; i++ {
		waitGroup.Add(2)
		go incr()
		go decr()
	}
	waitGroup.Wait()
	fmt.Println(counter) //

}
