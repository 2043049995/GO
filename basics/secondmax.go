package main

import "fmt"

func main() {

	//获取第二个最大值
	var nums = []int{1, 5, 3, 0, 6, 4, 66, 44, 22, 2}
	maxNum := nums[0]
	secMax := 0
	for _, v := range nums {
		if v > maxNum {
			secMax = maxNum
			maxNum = v
		} else if v > secMax {
			secMax = v
		}
	}
	fmt.Println(secMax)
}
