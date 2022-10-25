package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	addr := "127.0.0.1:9999"
	listen, err := net.Listen("tcp4", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()
	log.Printf("listen in %s\n", addr)
	conn, err := listen.Accept()
	if err != nil {
		log.Fatal(err)
	} else {
		defer conn.Close()
		log.Printf("client %s is connecting ...", conn.RemoteAddr())
		reader := bufio.NewReader(conn)
		for {

			line, _, err := reader.ReadLine()
			if err != nil {
				log.Fatal(err)
				break
			} else {
				fmt.Println("接收到的数据为：", string(line))
			}
		}
	}

}
