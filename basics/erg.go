package main

import "fmt"

func main() {
	s1 := "李朝阳go"
	for k, v := range s1 {
		fmt.Println(k, string(v))
	}
	//区分出英文字符与空格。英文字符在计算机内使用的是ascii编码，a-z,A-Z.
	//ch>="A"&ch <=Z,说明是一个大写的英文字母，小写字母判断方式与其类似
	map01 := make(map[string]int)
	map01["大写字母数量："] = 0
	map01["小写字母数量："] = 0
	map01["其他字符数量："] = 0

	fmt.Println(map01)
	s2 := `In as
sense we've come to our nation's capital to cash a check. When the architects of our republic wrote the
	magnificent words of the Constitution and the Declaration of Independence, they were signing a promissory note to 
	which every American was to fall heir. This note was a promise that all men, yes, black men as well as white men, 
	would be guaranteed the "unalienable Rights" of "Life, Liberty and the pursuit of Happiness." It is obvious today 
	that America has defaulted on this promissory note, insofar as her citizens of color are concerned. Instead of 
	honoring this sacred obligation, America has given the Negro people a bad check, a check which has come back marked
	"insufficient funds."`
	for _, value := range s2 {
		if value >= 'A' && value <= 'Z' {
			map01["大写字母数量："]++
		} else if value >= 'a' && value <= 'z' {
			map01["小写字母数量："]++
		} else {
			map01["其他字符数量："]++
		}
	}
	fmt.Println(map01)

}
