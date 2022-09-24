package main

import "fmt"

// 输入内容
func main() {
	var yes string
	fmt.Println("有西瓜吗？")
	fmt.Scan(&yes)
	if yes == "Y" || yes == "y" {
		fmt.Println("有西瓜")
	} else {
		fmt.Println("没西瓜")
	}
}
