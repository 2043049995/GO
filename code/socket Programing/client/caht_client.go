package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr := "127.0.0.1:9999"
	conn, err := net.Dial("tcp4", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("连接成功")
	defer conn.Close()
	//fmt.Printf("%T", conn)
	reader := bufio.NewReader(conn)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("输入一些数据：")
		scanner.Scan()
		fmt.Fprintf(conn, " %s\n", scanner.Text())

		lines, _, _ := reader.ReadLine()
		fmt.Println(string(lines))

	}

}
