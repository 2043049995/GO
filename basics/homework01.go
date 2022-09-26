package main

import "fmt"

func main() {

	//找切片内最大的元素，不能用排序
	//假设某个值为最大值，比较每个值与这个值的关系
	//草稿
	var nums = []int{7, 9, 4, 5, 6, 78, 200, 1}
	maxNum := nums[0]
	//for i := 1; i < len(nums); i++ {
	//	if maxNum > nums[i] {
	//		fmt.Println("数值交换")
	//	} else {
	//		maxNum = nums[i]
	//		fmt.Println("最大值为：", nums[i])
	//	}
	//}
	for _, v := range nums {
		if v > maxNum {
			maxNum = v
		}
	}
	fmt.Println(maxNum)
}
