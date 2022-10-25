package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"rpc/data"
)

func main() {
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer conn.Close()
	//服务名称就是结构体名称
	//定义请求对象
	request := &data.CalcRequest{2, 4}
	//定义相应对象
	response := &data.CalcResponse{}
	//调用远程方法获取结果
	err1 := conn.Call("Calc.Add", request, response)
	fmt.Println(err1, response.Result)

}
