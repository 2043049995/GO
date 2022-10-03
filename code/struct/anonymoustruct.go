package _struct

import "fmt"

// 结构体同函数变量一样，变量名为大写才可以被包外函数访问，其结构体内部的属性值大小写也按照此规则，命名嵌入也适用于此规则
//结构体互相嵌入的可见性，匿名嵌套：
/*
匿名嵌套需要遵守的规则

结构体名大写，属性名是大写，嵌入到一个结构体小写的结构体里
结构体大写，属性名小写，嵌入到小写结构体里
小写的结构体，大写的属性，嵌入到大写结构体里
S A =>s 由于结构体名小写，包外不能创建结构体对象
S a =>s 由于结构体名小写，包外不能创建结构体对象

S A =>S 结构体名大写，包外可访问，属性也是大写，属性也可以访问.此处不用管嵌套的结构体是大写还是小写，只关注属性是大小写
S a =>S 可以创建结构体对象，属性不可访问

s A =>S 结构体名大写，包外可访问，属性也是大写，属性也可以访问
s a =>S 可以创建结构体对象，属性不可访问

s A =>s 由于结构体名小写，包外不能创建结构体对象
s a =>s 由于结构体名小写，包外不能创建结构体对象


*/

func Anonymousstruct() {
	//结构体的匿名嵌入
	type Address struct {
		qu   string
		xian string
	}
	type ss1 struct {
		id   int
		name string
		addr Address
	}
	type student struct { //命名嵌入，将一个结构体当做一种数据类型定义进另一个结构体
		ss1        //如果不声明其类型，只写了名字称为匿名嵌入，实际上是听过类名自己创建出了一个属性（前提是有这个名字的结构体声明）只是少写了匿名结构体的数据类型的声明
		gender int //引用与被引用的结构体参数名冲突时会优先访问当前对象寻找，找不到才会去嵌入的对象去找
		//如果嵌入了多个结构体且名称由冲突的话，必须去指定访问哪个结构体里发生冲突的名字。此时程序已经无法判断。
	}
	//命名与匿名嵌入的时候如果嵌入的是结构体的指针，只需要在赋值的时候赋地址进来即可
	//Addr:&Address{}
	var me student
	fmt.Printf("%v\n", me)
	me.gender = 1
	me.ss1.name = "李钊"
	fmt.Println(me)
	type teacher struct {
		id   int
		name string
		add  *Address
	}
	var t1 = teacher{
		id:   1,
		name: "刘",
		add:  &Address{qu: "广阳区", xian: "开阳县"},
	} //可以将结构体初始化的过程封装成一个函数，传入相应的数据生成新的对象

	t1.add.xian = "新"
	fmt.Printf("%v", *t1.add)
}
