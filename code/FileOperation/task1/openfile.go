package task1

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// 实际上filed的open与create方法都是调用的openfile方法，
func Openfiles() {
	//追加文件
	path := "user2.txt"
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err == nil {
		fmt.Printf("%T", strconv.FormatInt(time.Now().Unix(), 10))
		log.SetOutput(file)
		log.Printf(strconv.FormatInt(time.Now().Unix(), 10))
		file.Close()

	} else {
		fmt.Println(err)
	}

}
