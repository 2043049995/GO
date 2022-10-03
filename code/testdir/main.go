package main

import (
	_struct "code/struct"
	"flag"
	"fmt"
)

// os模块
func main() {
	//var port int
	var host string
	var verbo bool
	var help bool
	//绑定命令行参数与变量关系
	//port := flag.Int("p", 22, "aaa")
	//flag.IntVar(&port, "p", 22, "sshhelp")
	flag.StringVar(&host, "h", "127.0.0.1", "hosthelp")
	flag.BoolVar(&verbo, "v", false, "detail log")
	flag.BoolVar(&help, "he", false, "helplog")
	//解析命令行参数
	fmt.Println(host)
	flag.Parse()
	fmt.Println(host)
	//flag.Usage = func() {
	//	fmt.Println("帮助信息")

	//}
	if help {
		flag.Usage()
		//flag.PrintDefaults()
	} else {
		//fmt.Println("。。。")
		//fmt.Println(flag.Args()) //没在命令中指定的可选变量可以通过args以切片的形式传进去
	}
	//fmt.Println(*port, host, verbo)
	//testdir.Time01()
	//testdir.Hash01()
	//testdir.Edcoding()
	//_struct.Anonymouse()
	//_struct.Anonymousstruct()
	//_struct.Method1()
	//_struct.Anonymousfunc()
	//_struct.Express()
	_struct.SliceSort()
}
