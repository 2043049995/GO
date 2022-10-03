package _struct

import "fmt"

type Dog struct {
	name string
}
type Cat struct {
	Name string
}

func Method1() {
	//方法的定义和调用

	dog1 := Dog{"黑点"}
	dog1.call()
	dog1.SetName("李") //直接传入引用类型
	dog1.call()
	(&dog1).SetName("6") //取引用。传入其引用类型，对需要单独处理的数据要加上小括号
	dog1.SetName("6")    //go语言里如果调用方法的对象是一个值类型，而函数的接收这是一个指针类型则会自动取引用。
	dog1.call()          //下方的相反的结构体传值与上面类似，会自动解引用
	pdog := Dog{"豆花"}  //指针类型可以直接调用，正常使用需要前面加*变为值类型
	pdog.call()          //需要保持原本的数据类型去调用接收对应传入参数的方法
	//var dogs *Dog        //如果定义的指针类型的数据没有初始化无法直接使用，指针类型默认值为nil，无法调用任何方法
	//自定义方法名如果为String的话，需要做的也是格式化输出，直接用变量名.结构体里的项就可以直接调用这个方法
	//这个方法必须用在fmt.println（）里面，是由这个函数去调用的，并不能单独使用

}
func (cat Cat) ca() { //自动取引用和解引用只会发生在函数接受者的位置，在函数的方法参数是不会进行解引用和取引用的。
	fmt.Println("喵喵", cat.Name)
}
func (dog Dog) call() {
	fmt.Println("汪汪", dog.name)
}
func (dog *Dog) SetName(name string) { //结构体默认是值类型，无法直接修改其值
	dog.name = name
}
