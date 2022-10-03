package codes

import (
	"flag"
	"fmt"
	"time"
)

func Flags() {
	var port int
	flag.IntVar(&port, "p", 22, "sshhelp")
	fmt.Println(port)
	time.Sleep(time.Second * 1)
}
