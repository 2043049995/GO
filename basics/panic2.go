package main

import "fmt"

func main() {
	//抓取其他函数中的错误进行处理
	if errs := test(); errs == nil {
		fmt.Println("未发生错误")
	} else {
		fmt.Println("发生错误")
	}

}
func test() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	panic("error")
	fmt.Println("666")
	return //由于该函数已声明了要接收一个err名称变量，所以此处不用retrun err
}
