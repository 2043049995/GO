package main

import (
	"fmt"
	"regexp"
)

func main() {
	//var s1 string = "^132\\d{8}$"
	//var s1 string = "^132[0-9]{8}$"
	var s1 string = "^(132|158)[0-9]{8}$"
	//^[a-zA-Z0-9]{1,32}@[a-z]{1,12}\\.edu$)
	//pattern1 := regexp.QuoteMeta("^ab") //将在正则中的特殊字符自动进行转义，匹配的就是传入的字符串
	s2 := "13215151515"
	s5 := "15815151515"
	s3 := "212121"
	s4 := "132ssss"
	fmt.Println(regexp.MatchString(s1, s5))
	fmt.Println(regexp.MatchString(s1, s2))
	fmt.Println(regexp.MatchString(s1, s3))
	fmt.Println(regexp.MatchString(s1, s4))
}
