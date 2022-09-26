package main

import (
	"errors"
	"fmt"
)

func main() {

	//错误处理，go语言中没有异常。显式的通过返回值告知调用者是否发生了错误
	if k, err := chu(10, 2); err == nil {
		fmt.Println(k)
	} else {
		fmt.Println(err)

	}

}

// 如何定义错误类型，如何创建错误对应的值
// 定义error，无错误的话返回nil,有错误返回已有的错误类型值，通过errors.new包创建

func chu(a, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("除数不能为0") //fmt.Errorf("ERROR:%s", "除数不能为0")
	}
	return a / b, nil
}
