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
	conn, err := net.Dial("tcp4", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("连接成功")
	defer conn.Close()
	//fmt.Printf("%T", conn)
	reader := bufio.NewReader(conn)
	for {
		fmt.Fprintf(conn, "Time is : %s\n", time.Now().Format("2006-01-02"))
		fmt.Fprintf(conn, "over")
		lines, _, _ := reader.ReadLine()
		fmt.Println(string(lines))
		fmt.Fprintf(conn, "over")

	}

}
