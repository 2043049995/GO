package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	addr := "127.0.0.1:9999"
	listen, err := net.Listen("tcp4", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()
	log.Printf("listen in %s\n", addr)
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		func() {
			defer conn.Close()
			log.Printf("client %s is connecting ...", conn.RemoteAddr())
			reader := bufio.NewReader(conn)
			for {

				line, _, err := reader.ReadLine()
				if err != nil {
					log.Println(err)
					break
				} else {

					fmt.Println("接收到的数据为：", string(line))
					time.Sleep(time.Second * 2)
					fmt.Fprintf(conn, "Time is : %s\n", time.Now().Format("2006-01-02"))
				}
			}

		}()

	}
}
