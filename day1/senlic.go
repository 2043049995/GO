package main

import "fmt"

func main() {
	var b = []int{}
	a := []int{1, 2, 3}
	b = append(a, []int{6, 7, 8}...)
	fmt.Println(b)
	var c = []int{}
	c = b[:len(b)-1]
	fmt.Println(c)
	//数组和切片的区别
	//数组不可变，是值类型。在进行赋值的之后会进行值拷贝,此案例为切片赋值会将slice2指向slice1所指向的地址（值拷贝，将slice1中的值拷贝到slice2中）。并不会开辟新的空间.此时对slice进行修改slice1的值
	//也会跟着改变
	slice1 := []int{1, 2, 3}
	slice2 := slice1
	slice2[1] = 10
	fmt.Println(slice1)
	//数组赋值的话会在内存中申请一块和他内存大小相同的一块块空间用来存储新的数据，新数组数据的变化不会影响就的数组
	arr1 := [10]int{1, 2, 3}
	arr2 := arr1
	arr2[0] = 100
	fmt.Println(arr1)
}
