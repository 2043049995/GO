package main

import "fmt"

func main() {
	//“”可解释的字符串
	//··原生字符串
	var a string = "123123"
	var b string = `
	ss
`
	//字符串索引
	var desc string = "李朝阳"
	fmt.Println(string([]rune(desc[0:2])))
	fmt.Println(a)
	fmt.Println(b)
}
