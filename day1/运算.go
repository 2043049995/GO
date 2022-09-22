package main

import "fmt"

func main() {
	//关系运算
	a := 1
	b := 2
	fmt.Println(a == b)
	fmt.Println(a != b)
	//位运算
	//十进制转换成二进制
	//&全1为1，| 其中1个为1就是1，异或：相同为0，不同为1.要转换成二进制运算
	fmt.Println(7 & 2)
	//移位运算
	fmt.Println(2<<1, 2>>1)
	var c int = 2
	var d int = 3 //两个数的数据类型必须一致才能进行算术运算
	fmt.Println(c + d)
}
