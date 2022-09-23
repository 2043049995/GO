package main

import "fmt"

func main() {
	var yes string
	fmt.Println("输入Y或N")
	fmt.Scan(&yes)
	switch {
	case yes == "y":
		{
			fmt.Println("yes")
		}
	case yes == "n":
		{
			fmt.Println("no")
		}
	default:
		{
			fmt.Println("输错了")
		}

		//switch yes {
		//case "y", "Y":
		//	{
		//		fmt.Println("输入了一个Y")
		//	}
		//case "n":
		//	{
		//		fmt.Println("输入了一个n")
		//	}

	}

}
