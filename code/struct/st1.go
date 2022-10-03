package _struct

import "fmt"

type Counter int //设计算数比较的话要先转换为相同类型才能比较，int(x)将x转换为int类型

//定义函数类型

func St1() {
	//自定义类型
	var count Counter
	fmt.Println(count)

}
