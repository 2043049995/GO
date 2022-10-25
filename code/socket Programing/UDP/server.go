package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ipAddress := "0.0.0.0:9888"
	packageConn, err := net.ListenPacket("udp", ipAddress)
	defer packageConn.Close()
	if err != nil {
		log.Fatal(err)
	}
	ctx := make([]byte, 1024)
	for {
		n, addr, err := packageConn.ReadFrom(ctx)
		if err != nil {
			log.Fatal(err)
			continue
		}
		//fmt.Println(n, addr)
		fmt.Println(string(ctx[:n]))
		packageConn.WriteTo([]byte("This is server"), addr)
	}

}
