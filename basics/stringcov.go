package main

import (
	"fmt"
	"strconv"
)

func main() {
	//字符串到布尔类型的转换
	v, err := strconv.ParseBool("true")
	fmt.Println(v, err)
	//其他类型到字符串类型的转换
	sd := fmt.Sprintf("%d", 12) //借助fmt包将数字转换为字符串类型
	fmt.Printf("%T", sd)
	sd1 := fmt.Sprintf("%f", 12.0) //借助fmt包将数字转换为字符串类型
	fmt.Printf("%#v", sd1)

}
