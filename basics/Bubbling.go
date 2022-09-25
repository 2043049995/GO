package main

import "fmt"

func main() {
	height1 := []int{1, 4, 10, 5, 6, 7, 2, 3}
	//从第一个人开始，两两进行比较
	for a := 0; a < len(height1); a++ {
		for i := 0; i < len(height1)-1; i++ {
			if height1[i] > height1[i+1] {
				height1[i], height1[i+1] = height1[i+1], height1[i]
			}
		}
		fmt.Println(height1)

	}

}
