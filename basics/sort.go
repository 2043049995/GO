package main

import (
	"fmt"
	"sort"
)

func main() {
	//可以使用原生自带的sort模块对切片进行排序
	slice1 := []int{1, 4, 6, 9, 3}
	sort.Ints(slice1)
	fmt.Println(slice1)

	//arr1 := [6]int{2, 5, 1, 3, 7, 4}
	//sort.i
	//fmt.Println(arr1)
	//对字符串进行排序
	slice3 := []string{"李", "123", "222", "22234"}
	sort.Strings(slice3)
	fmt.Println(slice3)
}
