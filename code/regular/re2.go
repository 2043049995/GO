package main

import (
	"fmt"
	"regexp"
)

func main() {
	re, _ := regexp.Compile("132\\d{8}")
	fmt.Println(re.MatchString("132xxx"))
	fmt.Println(re.MatchString("13222222222"))
	//替换
	fmt.Println(re.ReplaceAllString("我的电话是：132xxx", "132????????"))
	fmt.Println(re.ReplaceAllString("我的电话是：13222222222", "132????????"))
	//查找
	fmt.Println(re.FindAllString("我的电话是：13222222222,我的电话是：13222222223，我的电话是：15822222222", -1))
	//分割
	reg2, _ := regexp.Compile("[,，：:?]")
	fmt.Println(reg2.Split("我的电话是：13222222222,我的电话是：13222222223，我的电话是：15822222222", -1))
}
