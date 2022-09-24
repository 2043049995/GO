package main

import "fmt"

func main() {
	i := 0
	result := 1
	for i < 100 {
		result = result + i
		i++
		if i == 50 {
			goto END
		}
	}
END:
	{
		fmt.Println("循环结束", i)
	}
}
