package task1

import (
	"fmt"
	"os"
)

func FileWriet() {
	//创建文件，写入内容，关闭文件
	PATH := "write_text.txt"
	file, err := os.Create(PATH)
	if err != nil {
		fmt.Println(err)
	} else {
		//写文件
		//file.Write([]byte("a！@#￥")) // 写入数据要先转换为byte形式，用byte的话需要将字符串转换为字节数组
		//读的时候只能读字节数组，写的时候可以写字节切片或字符串
		file.WriteString("xswdf测试") //每次执行写动作的时候都会将文件之前的内容清空。先清除再写入，并不是覆盖
		//关闭文件
		file.Close()
	}

}
