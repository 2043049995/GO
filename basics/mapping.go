package main

import "fmt"

func main() {
	//映射是存储一系列无序的key/value对，通过key对value进行访问,map为空的时候是不能进行任何操作的,必须将值初始化后才能进行操作
	var score map[string]int
	//	fmt.Printf("%T", score)
	//	fmt.Println(score == nil)
	score = make(map[string]int) //声明了一个个空映射
	fmt.Println(score)
	score = map[string]int{"李朝阳": 100, "李钊": 200}
	fmt.Println(score)
	score1 := make(map[string]int)
	fmt.Println(score1)
	score1 = map[string]int{"李杰": 1000}
	fmt.Println(score1)
	//增删改查
	fmt.Println(score1["李杰四"]) //访问一个不存在的值默认会返回空值
	//只靠判断反悔的是不是0值去确定某个key是否存在是不准确的
	//在查询某个key的时候会返回两个值，第二个值表示这个key是否存在
	if _, ok := score1["李杰"]; ok {
		fmt.Println(ok)
		fmt.Println("值存在")
	} else {
		fmt.Println(ok)
		fmt.Println("值不存在")
	}
	for k, v := range score {
		fmt.Println(k, v)
	}
	var score2 map[string]int
	score2 = map[string]int{"李朝阳": 100, "李钊": 200}
	//map的key无法修改，key写错的话只能删掉重新添加
	delete(score2, "李朝阳")
	fmt.Println(score2)
	fmt.Println(score2["李钊"])
	//添加元素
	score2["王鹏"] = 2222
	fmt.Println(score2)
	score2["王鹏"] = 111
	fmt.Println(score2)
	fmt.Println(len(score2))
	for k, v := range score2 {
		fmt.Println(k, v) //map是无序的遍历的时候与插入的时候顺序不一样
	}
	//key至少可以判断是否已存在
	//key至少可以进行一些== ！=的运算，bool、数组、int、字符串，value可以是任意类型
	var map1 map[string]string
	map1 = map[string]string{"李朝阳": "里"}
	fmt.Println(map1)
	var map2 map[string]map[string]string
	map2 = map[string]map[string]string{"李钊": {"223": "7E"}}
	fmt.Println(map2)
}
