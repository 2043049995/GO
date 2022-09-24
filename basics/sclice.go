package main

import "fmt"

func main() {
	//切片是长度可变的数组
	var nums []int
	fmt.Printf("%T\n", nums)
	fmt.Println(nums == nil) //此处判断的是指针，nil表示为空指针，切片定义出来未使用的话也是指向空
	//切片的赋值
	nums = []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	nums = []int{2, 2, 2, 2}
	fmt.Println(nums)
	nums[1] = 100
	fmt.Println(nums)
	///数组切片赋值
	var num1 [10]int = [10]int{1, 2, 3, 4}
	var num2 []int
	num2 = num1[:]
	fmt.Println(num2)
	fmt.Println(len(num2), cap(num2))
	//make函数声明切片,对切片里面的指针进行初始化
	var nums01 []int
	nums01 = make([]int, 5)
	fmt.Println(nums01)
	nums01[2] = 222
	fmt.Println(nums01)
	nums01 = make([]int, 9, 9) //改变切片的容量也会导致其内部数据丢失
	//fmt.Println(nums01)
	//元素操作
	//fmt.Println(nums01[0])
	//fmt.Println(nums01[2])
	//nums01[2] = 999
	fmt.Println(nums01[2])
	////fmt.Println(nums01[99])
	//fmt.Println(len(nums01), cap(nums01))
	//fmt.Printf("%p\n", &nums01)
	//容量是5，长度是3，后面的两个容量如何使用
	//fmt.Println(nums01)
	fmt.Printf("%p\n", &nums01[0])
	nums01 = append(nums01, 5) //在用append追加的时候会开辟一个新的数组，返回的是一个值类型，将nums01指到新的数组地址去，销毁旧的数组
	fmt.Printf("%p", &nums01[0])
	//fmt.Println(len(nums01), cap(nums01))
	fmt.Println(nums01)
	for _, v := range nums01 {
		fmt.Println(v)
	}
	fmt.Println("666")
	for i := 0; i < len(nums01); i++ {
		fmt.Println(nums01[i])
	}
	//fmt.Printf("%p", &nums01)
	//nums01[9] = 3333。数组的内存空间是连续的，对于切片来说要尽量减少开辟空间的次数，在长度与容量相同时再次申请会开辟一个比原始切片大1-1.5倍的容量
	//fmt.Println(nums01)
	//删除
	//fmt.Println(nums01)
	for i := 0; i < len(nums01); i++ {
		nums01[i] = i + 1
	}
	fmt.Println(nums01)
	//copy函数
	//var nums02 = []int{10, 20, 30}
	//copy(nums01, nums02)
	fmt.Println(nums01)
	//需要用copy实现删除的机制
	//nums03 := nums01[1:]
	copy(nums01[2:], nums01[3:])
	fmt.Println(nums01)
}
