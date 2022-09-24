package main

import "fmt"

func main() {
	//数组实现队列
	var nums = []int{}
	//nums = make([]int, 0)
	//var num int = 1
	var adds int
	var index int
	var inputs string
	for {
		fmt.Println("输入要做的操作，增加、删除、修改、查询")
		fmt.Scan(&inputs)
		switch inputs {

		case "增加":
			{
				fmt.Println("输入需要增加的数字")
				fmt.Scan(&adds)
				nums = add(nums, adds)
				//fmt.Println("666")
				fmt.Println(nums)
				break
			}
		case "查看":
			{
				fmt.Println("输入要查看数字的下标：")
				fmt.Scan(&adds)
				fmt.Println(nums[adds])
			}
		case "删除":
			{
				//fmt.Println("输入要删除数字的下标：")
				//fmt.Scan(&index)
				nums = delete(nums, index)
				fmt.Println(nums)
			}
		}
	}
}
func delete(nums []int, index int, numOptional ...int) []int {
	//num := 1
	//if len(numOptional) > 0 {
	//num = numOptional[0]
	//}
	//nums = nums[1:]
	fmt.Println("删除完成后的值：", nums)
	return nums
}

func add(nums []int, adds int) []int {
	nums = append(nums, adds)
	return nums
}
