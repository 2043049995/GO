package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "李朝阳时尚大师大所大所大所多"
	s2 := s1[1:3]
	fmt.Println(s2)
	//bytes:将buffer对象转换为字节切片
	//strings：将buffer对象转换成字符串
	strings.Split(s1, "666")
	fmt.Println(s1)
	fmt.Println(strings.Compare("abc", "ad"))
	fmt.Println(strings.ContainsAny("abc", "acb"))
	fmt.Println(strings.Count("aaasdad", "a"))
}
