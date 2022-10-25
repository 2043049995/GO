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

func lss(conn net.Conn) {
	file, _ := os.Open(".")
	defer file.Close()
	names, _ := file.Readdirnames()
	suffix := ""
	if len(names) > 0 {
		suffix = ":"
	}
	strings.Join(names, ":")
	fmt.Fprintf(conn, "%d|%s%s", len(names), strings.Join(names, ":"), suffix)
}
func cmd(conn net.Conn) (string, []string) {
	reader := bufio.NewReader(conn)
	op, err := reader.ReadString('|')
	log.Println("reader op error", err)
	cnttext, err := reader.ReadString('|')
	log.Println("conttext error", err)
	cont, err := strconv.Atoi(cnttext[:len(cnttext)-1])
	log.Println("conver error", err)
	args := make([]string, 0, cont)
	for cont > 0 {
		param, err := reader.ReadString('|')
		log.Println("read paramenter error", err)
		args = append(args, param[:len(param)-1])
		cont--
	}
	return op[:len(op)-1], args

}
func handleConn(conn net.Conn) {
	defer conn.Close()
	op, args := cmd(conn)
	switch op {
	case "ls":
		lss(conn)
	default:
		fmt.Println("参数错误")

	}
	fmt.Println(op, args)
}
func main() {

	logfile, _ := os.OpenFile("logfile.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	defer logfile.Close()
	log.SetOutput(logfile) //带缓冲io
	addr := "0.0.0.0:9999"
	listener, err := net.Listen("tcp4", addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accept error", err)
			continue
		}
		log.Printf("client %s online", conn.RemoteAddr())
		handleConn(conn)

	}

}
