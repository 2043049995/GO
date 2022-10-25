package src

import (
	"flag"
	"fmt"
	"os/exec"
)

func FlagReception() {
	// 1 flag.String   flag.Int
	//第一个参数：命令行参数的名称
	//第二个参数：命令行不输入时的默认值
	//第三个参数：该参数的描述信息，help命令时会显示
	var (
		remotIp   string
		remoePort int
	)
	flag.StringVar(&remotIp, "host", "localhost", "This is default ip")
	flag.IntVar(&remoePort, "port", 22, "This is default port")
	flag.Parse()
	fmt.Println("IP为", remotIp, "端口为：", remoePort)
	cmd := exec.Command("ping", "8.8.8.8")
	cmd.Run()
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("output err:", err)
		return
	}
	fmt.Println(string(output))
}
