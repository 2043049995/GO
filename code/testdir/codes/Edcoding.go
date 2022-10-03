package codes

import (
	"encoding/base64"
	"fmt"
)

func Edcoding() {
	//base64编解码实验,[]byte("123")作用是将字符串转为字节切片，字节切片不能进行比较，字符串可以进行比较
	x := base64.StdEncoding.EncodeToString([]byte("123")) //encodetostring表示将一个字符数组变成一个字符串
	fmt.Printf(x)
	y, err := base64.StdEncoding.DecodeString(x)
	fmt.Println(string(y), err)
}
