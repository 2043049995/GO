package main

//go不支持给参数设置默认值
//函数所有的分支必须有返回值
import (
	"fmt"
)

func main() {
	fmt.Println(sums(4, 6))
	//a, b, c, d := math(4, 8)
	//fmt.Println(a, b, c, d)
}

// 函数的返回值也可以是变量名
func sums(a, b int) (sum int, jian int) {
	sum = a + b
	jian = a - b
	return

}

// 函数的返回值有多个，接受的时候要用相同数量的变量去接收
func math(a, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}

// 切片内指定位置数据删除
func delslice() {
	slices := []int{1, 23, 4, 5, 6, 7, 8}
	slices1 := slices[:1]
	slices2 := slices[2:]
	slices1 = append(slices1, slices2...)
	fmt.Println(slices1)
}

//	func nopa()int {
//		var slice1 = []int{1,2,3}
//		return nopa1(slice1...)
//	}
//
//	func nopa1(args ...int)[]int {
//		return args
//	}
func cacl(op string, a, b int, args ...int) int {
	switch op {
	case "add":
		return addN(a, b, args...) ///通过...可以对切片进行解包，也只能对切片进行解包

	}
	return -1
}
func addnum(a, b int) int { //函数参数的类型合并
	fmt.Println("返回前")
	return a + b
}
func addN(a, b int, args ...int) int { ///... int 表示可选参数,int为其数据类型,可变参数必须放在函数参数的最后一个，一个函数只能有一个可变参数
	fmt.Println(a, b, args)
	//fmt.Printf("%T", args)
	c := a + b
	for _, v := range args {
		c = c + v
	}
	return c

}
