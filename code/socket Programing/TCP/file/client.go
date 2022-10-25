package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func ls(conn net.Conn) {
	fmt.Fprintf(conn, "ls|0")
	reader := bufio.NewReader(conn)
	//fmt.Println(reader.ReadLine())
	sizeText, err := reader.ReadString('|')
	fmt.Println(sizeText)
	if err != nil {
		log.Fatal(err)
	}
	size, _ := strconv.Atoi(sizeText)
	for size > 0 {
		fmt.Println(reader.ReadString('|'))
		size--
	}
}
func main() {
	logfile, _ := os.OpenFile("client.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	defer logfile.Close()
	log.SetOutput(logfile)
	addr := "127.0.0.1:9999"
	conn, err := net.Dial("tcp4", addr)
	if err != nil {
		log.Printf("Connecterror: %#v", err)
	}
	defer conn.Close()
	log.Println("connect fileserver success")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("输入要执行的命令:")
		scanner.Scan()
		input := scanner.Text()
		cmds := strings.Split(input, " ")
		switch cmds[0] {
		case "ls":
			ls(conn)
		default:
			fmt.Println("Command Error")

		}
	}
}
