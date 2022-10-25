package main

import (
	"fmt"
	"regexp"
)

func main() {
	//贪婪模式与非贪婪模式
	//re, _ := regexp.Compile("[0-9]+")
	//re, _ := regexp.Compile("[ab][0-9]+")//非贪婪模式，按最小的匹配，多个匹配到的字符串在一起的话会一起输出
	re, _ := regexp.Compile("(?U)[ab][0-9]+") //贪婪模式，按匹配出来的符合要求的数目字符串最大的匹配，将多个连续的符合要求的字符串都单独匹配出来

	//fmt.Println(re.FindAllString("0-a232-asda-b343", -1))
	fmt.Println(re.FindAllString("0-a232-asda-b343-a23b232", -1)[1])
	re.Longest() //将非贪婪模式转换为贪婪模式

}
