package task1

import (
	"fmt"
	"io"
	"os"
)

//文件操作需要使用到OS包

func FileRead() {
	PATH := "user.txt"         //确认路径
	file, err := os.Open(PATH) //打开文件，会返回文件信息和错误信息，err表示打开失败时的错误信息
	if err != nil {
		fmt.Println(err)
	} else {
		var bytes []byte = make([]byte, 20)
		//连续读取，通过点断EOF来判断文件是否读取完成
		for {
			n, err := file.Read(bytes)
			fmt.Printf(string(bytes[:n]))
			if err == io.EOF {
				break
			}
		}
		//文件读取完成后需要关闭文件
		file.Close()
		//n, err := file.Read(bytes) //此处的n返回的就是目标文件文件的内容长度，由于此前定义了切片的长度，但是文件没有将切片填充满，就要通过返回的n来确定文件内有多少数据
		////最大的读取数量为初始化的切片的长度，如果想再读取后20个元素，只需要再将bytes切片用fule.read处理一次,文件不要关闭。如果文件过大，可以通过是否返回了EOF错误信息来判断文件有没有读取完成
		//
		//fmt.Println(n, err)
		//fmt.Println(string(bytes)) //会将对应文件内容中的数据的码点打印出来
		//fmt.Println(string(bytes[:n]))
		//n, err = file.Read(bytes)
		//fmt.Println(string(bytes[:n]))

	}
	//fmt.Println(file) //file 会返回一个指针，数据类型为*os.file类型
}
