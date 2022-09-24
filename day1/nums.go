package main

import "fmt"

func main() {
	var arr [10]int
	//fmt.Println(arr)
	arr = [10]int{1, 2, 3, 6: 20, 7: 30}
	//int[1] = 10
	fmt.Println(arr)
	nums01 := [3]int{1, 3, 5}
	nums02 := [3]int{1, 24, 4}
	fmt.Println(nums01 != nums02)
	//获取数组长度
	fmt.Println(len(nums01))
	//索引
	nums02[2] = 1000
	fmt.Println(nums01[2], nums02[2])
	for k, v := range nums02 {
		fmt.Println(k, v)
		//使用简短声明的时候在左侧必须有个新的变量
		var value int
		//_, value = 1, 2, 3
		fmt.Println(value)
	}
	s := "123123123"
	s1 := s[1:4]
	nums03 := nums01[1:2] //nums[start:end:capcity]
	fmt.Printf("%T\n", s1)
	fmt.Printf("%T", nums03)
}
