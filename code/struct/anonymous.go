package _struct

import "fmt"

func Anonymouse() {
	//匿名结构体，不定义结构体的名字，直接使用一个结构定义一个变量,在数据渲染，数据需要传递给模板的时候数据只会使用一次，可以用此方式
	a := struct {
		id   int
		name string
	}{}
	a.id = 2
	//结构体的一个属性可以使另一个结构体，称为结构体的嵌入
	type Address struct {
		qu   string
		xian string
	}
	type ss1 struct {
		id   int
		name string
		addr Address
	}
	addr := Address{"广阳区", "广西县"}
	var ss2_ = ss1{id: 11, name: "刘总", addr: addr}
	fmt.Println(ss2_.id, ss2_.name, ss2_.addr.qu, ss2_.addr.xian)
	ss2_.addr.qu = "北方大区"
	fmt.Println(ss2_)

}
