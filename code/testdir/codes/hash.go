package codes

import (
	"crypto/md5"
	"fmt"
)

func Hash01() {
	//加解密需要引入crypto包有若干种加密算法，sha256 sha512 md5等
	//在计算的时候必须传字节形式的数组，如果传入的是字符串需要用字节数组进行转换。会计算出其内部文字的哈希值
	fmt.Printf("%x", md5.Sum([]byte("123")))
	//如果需要计算md5的内容比较小的话直接计算即可，如果文件较大需要创建md5对象
	m := md5.New()                   //创建一个md5类型的对象
	m.Write([]byte("7777777777777")) //批量写一些数据,计算实际文件的md5值需要创建md5对象的方法
	m.Write([]byte("6666666666666")) //批量写一些数据

	fmt.Printf("%x", m.Sum(nil))
	//sha1 sha256 等算法原理一样

}
