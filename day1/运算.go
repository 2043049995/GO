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
	//类型转换
	var aa int32 = 555
	bb := int8(aa)
	fmt.Println(aa)
	fmt.Println(bb) //丢失精度
	var cc byte = 'A'
	var dd rune = '我'
	fmt.Println(cc, dd)
	fmt.Printf("%T,%c\n", cc, cc)
	fmt.Printf("%T,%q\n", dd, dd)
	aaa := "a李朝阳"
	fmt.Println(len([]rune(aaa)[:]))
	fmt.Println([]rune(aaa))
	fmt.Println([]byte(aaa))
	bbb := []rune(aaa)[1:2]
	fmt.Println(string(bbb))
	fmt.Printf("%U %b\n", bbb, bbb)
	var aaaa float64 = 1.111
	fmt.Println(aaaa)
	var bbbb float64 = 1.111
	fmt.Println(bbbb)

}
