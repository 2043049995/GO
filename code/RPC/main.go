package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpc/rpcserver"
)

func main() {
	//注册服务
	rpc.Register(&rpcserver.Calc{})
	addr := "127.0.0.1:9999"
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	defer listen.Close()
	log.Println("listen success in ", addr)

	for {
		conn, err := listen.Accept()
		defer func() {
			conn.Close()
		}()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		log.Printf("client connected :%s\n", conn.RemoteAddr())
		//使用例程启动jsonrpc处理客户端请求
		log.Println("链接里的数据为：", conn)
		go jsonrpc.ServeConn(conn)
	}

}
