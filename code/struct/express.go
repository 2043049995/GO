package _struct

import "fmt"

// 方法表达式
type Dog01 struct {
	name string
}

func (dog Dog01) call() {
	fmt.Printf("%s", dog.name)
}

//自动生成
/*func (dog *Dog01) call() {
	fmt.Printf("%s", dog.name)
}
func (dog *Dog01) SetDogName(name string) {
	dog.name = name
}*/

// 在调用一个值类型的方法的时候可以用指针对象去调用，指针类型的方法不能用值类型去调用
// 因为go会根据值接收者方法自动生成对应的隐式指针接收者方法，而指针接收者方法不会自动生成值接收者方法
// 因为此时无论是指针调用还是值调用都不会改变原对象的值，而调用会改变原对象的值
func Express() {
	method1 := Dog01.call
	//method2 := (*Dog01).SetDogName
	//fmt.Printf("%T  %T", method1, method2)
	//func(_struct.Dog01)  func(*_struct.Dog01)
	//函数的调用,对于方法表达式来说在调用的时候必须传递一个对象进去
	//在函数的接收者的时候才会自动取引用和解引用，自己声明的方法不会自动取引用和解引用

	dogs1 := Dog01{"豆豆"}
	method1(dogs1)
	//dogs1.SetDogName("小黑")
	method1(dogs1)

}
