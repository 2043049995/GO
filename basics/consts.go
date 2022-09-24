package main

import "fmt"

func main() {
	const name string = "lizhaoyang" //常量的值不可以修改
	fmt.Println(name)
	const user string = " "
	fmt.Println(user)
	const (
		id int = 1
		id1
		id2
	) //常量在批量赋值的时候如果不指定其值与数据类型默认会与上一个常量的类型与值相同，必须是将值与类型全部省略
	fmt.Println(id, id1, id2)
	//枚举类型,可以直接做算数运算
	const (
		a1 = iota
		a2
		a3
		a4 = 2
		a5 = iota
	)
	fmt.Println(a1, a2, a3, a4, a5)
}
