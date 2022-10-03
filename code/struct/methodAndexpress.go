package _struct

//方法值就是获取一个对象后把他的方法赋值给另一个函数
import "fmt"

type Pig struct {
	id   int
	name string
}

func (pig *Pig) call() {
	fmt.Println("牛逼", pig.name)
}
func (pig Pig) GetPigName() string {
	return pig.name
}
func (pig *Pig) SetPigName(name string) {
	pig.name = name
}
func Express1() {
	//方法是为特定类型定义的，只能由该类型调用的函数
	/*
			方法是添加了接受者的函数，接受者必须是自定义的类型
			方法值与方法表达式
			方法也可以赋值给变量，存储在数组，切片，映射中，也可以作为参数传递给函数作为函数的返回值
		创建方法值后再对方法进行修改不会影响方法值
	*/
	pig := Pig{1, "猪"}
	m1 := pig.call //再赋值的时候会把这个值类型数据进行copy，复制到m1这个函数里，改之前的值不会影响到原始的dog值，
	//如果将call方法改为指针类型，此时dog会自动取引用，此时m1函数内存储的就是dog的指针，当初始值发生改变后由于m1调用的是指针，所以值也会跟着变

	m1()
	pig.SetPigName("狗")
	pig.call() //此值已发生改变
	m1()       //打印结果不会发生改变，因为其初始值已经赋给m1，不管后面的更改结果如何都不会更改m1方法里的dog。
	//pig1 := Pig{2, "666"}
	//pig1.call()
	//pig1.SetPigName("77")
	//pig1.call()
	pdog := &Dog{"w"}
	m2 := pdog.call //会自动解引用变为值类型
	m2()
	//所以如果是定义了一个方法值，这个值以后再调用的事后就不会发生改变。
	//因为这个值已经在定义的时候复制到函数中去了。
	//创建了方法值后再对对象进行修改不会影响方法值
	//如果是指针接收者的话在调用的时候会去拿现在的值，值类型接收者会进行copy
	//
}
