package main

import "fmt"

func main() {
	//a,b,c三个柱子，需要将a柱子的圆圈借助跳板b移动到c上，每次驿动的时候只能把笑的放到大的圆盘上面
	fmt.Println("3个盘子")
	tower("A", "B", "C", 3)
}

func tower(a, b, c string, layer int) { //abc三个柱子，layer个盘子
	if layer == 1 {
		fmt.Println(a, "->", c)
		return
	}
	//a n-1 借助c移动到b
	tower(a, c, b, layer-1)
	fmt.Println(a, "->", c)
	//b n -1 借助a移动到c
	tower(b, a, c, layer-1)
}
