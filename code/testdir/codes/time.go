package codes

import (
	"fmt"
	"time"
)

func Time01() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Format("2006-01-02 15:04:05.000")) //必须是这个值，否则时间不准 go语言诞生时间
	fmt.Println(now.Unix())
	fmt.Println(time.Now().Format("2006/01/02 15:04:05.000"))
	t, err := time.Parse("2006-01-02 15:04:05.000", "2006-01-02 15:04:05.000")
	fmt.Println(t, err)
	time.LoadLocation("Asia/Shanghai")
	time.Sleep(time.Second * 5) //使程序休眠,time.second 表示秒
	fmt.Println("666")
}
