package main

import "fmt"

func main() {
	names := map[string]int{"李朝阳": 0, "李杰": 0, "刘杰": 0}
	var targetName string
	for {
		fmt.Println("输入要投给谁票：")
		fmt.Scan(&targetName)
		if _, v := names[targetName]; v {
			names[targetName]++
			for key, value := range names {
				fmt.Println("姓名"+key, "票数", value)
			}
		}

	}
}
