package main

import "fmt"

func main() {
	var a map[string]int = map[string]int{"李朝阳": 2}

	b := a
	b["李朝阳"] = 3
	fmt.Println(a, b)
	//对于值类型来说，赋值会产生新的数据副本，引用类型赋值只是修改了新数据的指针指向了旧数据

}
