package _struct

import (
	"fmt"
	"time"
)

type User struct {
	ID       int
	name     string
	birthday time.Time
}

type UserPoint *User

func Std2() {
	var me User = User{1, "李朝阳", time.Now()}
	fmt.Println(me.name)
	m1 := &User{2, "李钊", time.Now()}
	(*m1).ID = 33    //将指针放在括号里解引用后可以直接赋值
	m1.name = "刘解" //go语言里可以省略解引用直接去访问
	
	fmt.Printf("%v", m1)
}
