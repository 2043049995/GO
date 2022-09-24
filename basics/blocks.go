package main

import "fmt"

//作用域：定义标识符作用的范围

func main() {
	outer := 1
	{
		fmt.Println(outer)
		outer := 2
		//outer1 := 2//无法被外部使用。子语句可以使用父语句块中的标识符，父不能使用子语句块的
		{
			fmt.Println(outer)
		}
	}
	fmt.Println(outer) //语句块内定义的变量外边无法使用
}
