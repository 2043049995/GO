package main

import (
	"fmt"
)

func main() {
	//result := 0
	//for i := 1; i <= 100; i++ {
	//	result = result + i
	//}
	//i := 1
	//for i <= 100 {
	//	result = result + i
	//	i++
	//}
	//fmt.Println(result)
	//for range 遍历字符串
	var name string = "go李朝阳"
	for k, v := range name {
		fmt.Println(k, string(v))
	}

}
