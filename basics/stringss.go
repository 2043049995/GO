package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var byte1 []byte
	byte1 = []byte{'s', 'k'}
	fmt.Println(string(byte1)) //string可以将字符切片转换为字符串
	s1 := "gogogo测试"
	fmt.Println(len(s1))
	s2 := []rune(s1)
	fmt.Println(len(s2))
	fmt.Println(utf8.RuneCountInString(s1))

}
