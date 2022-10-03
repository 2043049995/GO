package _struct

import "fmt"

type UserType struct {
	id   int
	name string
}
type enginner struct {
	UserType
	Salary float64
}

func (user UserType) GetId() int {
	return user.id
}
func (user UserType) GetName() string {
	return user.name
}
func (user *UserType) SetName(name string) {
	user.name = name
}

func Anonymousfunc() {
	//结构体匿名方法调用
	me := enginner{UserType{id: 1, name: "李朝阳"}, 8.88}
	fmt.Println(me)
	fmt.Println(me.UserType.GetId())
	fmt.Println(me.UserType.GetName())
	me.UserType.SetName("牛逼")
	fmt.Println(me.GetName()) //如果嵌套的结构体某个名字只出现过一次，可以去掉匿名属性直接调用
	
}
