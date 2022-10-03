package codes

import (
	"fmt"
	"log"
)

func Logs() {
	log.Printf("我是printf日志%s", "xxx") //打印时会记录其生成的准确时间
	//log.Panic("我是panic")              //打印日志的时候触发panic
	log.SetPrefix("我是前缀") //为日志增加一个前缀，主要用于日志的分类
	fmt.Println("666")
	//log.Fatal("777") //打印完成后就会退出表示发生致命错误
	//log.Fatal("888")
	//设置日志打印的格式
	log.Flags()

}
