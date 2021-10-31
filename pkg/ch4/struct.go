package ch4

import (
	"fmt"
	"math"
)

//当为某个类型定义多个方法的时候, 良好的风格是: 所有的方法的接受者要保持一致, 要么都是变量, 要么都是指向变量的指针

//对变量v取地址: &v
//对指向某个变量的指针p解引用: *p
//
//当某类型的方法的接收者形参为*typeFoo时, 实参可以使用typeFoo指针类型, Go会自动取地址
//当某类型的方法的接收者形参为typeFoo时, 实参可以使用*typeFoo指针类型, Go会自动解引用
//
//传递的实参到底是该类型实例的拷贝还是指向该类型实例的指针值的拷贝取决于函数签名中的形参
//接收者形参为*typeFoo时, 传递的是指针值的拷贝
//接收者形参为typeFoo时, 传递的是实例的拷贝
//
//除了方法(成员函数)外, 一个*typeFoo实例可以直接使用"."访问typeFoo的字段(成员变量), Go会自动解引用

////给int绑定方法, 错误!
////Cannot define new methods on the non-local type 'builtin.int'
//func (instance int) IncrementBoundToType() {
//	instance = instance+1
//}

type MyInt int

func (instance MyInt) IncrementBoundToType() {
	instance += 1
	fmt.Printf("inside method: instance value: %#v\n", instance)
	fmt.Printf("inside method: instance address: %p\n", &instance)
}
func (instancePointer *MyInt) IncrementBoundToTypePointer() {
	*instancePointer += 1
}

type Employee struct {
	ID                          int
	Name                        string
	Position                    string
	Salary                      int
	invisibleFieldOutsideThePkg string
	Manager                     *Employee
}

func (instance Employee) IncrementSalaryBoundToType(amount int) {
	instance.Salary += amount
	fmt.Printf("inside method: instance value: %#v\n", instance)
	fmt.Printf("inside method: instance address: %p\n", &instance)
}
func (instancePointer *Employee) IncrementSalaryBoundToTypePointer(amount int) {
	(*instancePointer).Salary += amount
}

func TypeMethod() {
	var i MyInt
	var e Employee

	fmt.Println("int alias:")
	fmt.Println("call type-bound method on var and pointer")
	fmt.Printf("%#v\n", i)
	i.IncrementBoundToType()
	fmt.Printf("%#v\n", i)
	(&i).IncrementBoundToType()
	fmt.Printf("%#v\n", i)
	fmt.Println("call pointer-bound method on var and pointer")
	fmt.Printf("%#v\n", i)
	i.IncrementBoundToTypePointer()
	fmt.Printf("%#v\n", i)
	(&i).IncrementBoundToTypePointer()
	fmt.Printf("%#v\n", i)

	fmt.Println("alias of a certain struct:")
	fmt.Println("call type-bound method on var and pointer")
	fmt.Printf("%#v\n", e)
	e.IncrementSalaryBoundToType(1)
	fmt.Printf("%#v\n", e)
	(&e).IncrementSalaryBoundToType(1)
	fmt.Printf("%#v\n", e)
	fmt.Println("call pointer-bound method on var and pointer")
	fmt.Printf("%#v\n", e)
	e.IncrementSalaryBoundToTypePointer(1)
	fmt.Printf("%#v\n", e)
	(&e).IncrementSalaryBoundToTypePointer(1)
	fmt.Printf("%#v\n", e)
}

func StructBasics() {
	//定义了type别名的struct的实例化

	//这种赋值方式需要按顺序初始化所有的field
	liu := Employee{
		0, "玄德", "统帅", 10000, "汉室姓刘", nil,
	}
	zhuge := Employee{
		ID:                          1,
		Name:                        "孔明",
		invisibleFieldOutsideThePkg: "奉诏讨贼, 匡扶汉室",
		Manager:                     &liu,
	}
	wolong := new(Employee)
	//自动解引用, 相当于(*wolong).ID = 1
	wolong.ID = 1
	wolong.Name = "孔明"
	wolong.invisibleFieldOutsideThePkg = "奉诏讨贼"
	wolong.Manager = &liu

	//比较
	fmt.Printf("zhuge == *wolong: %t\n", zhuge == *wolong)
	wolong.invisibleFieldOutsideThePkg += ", 匡扶汉室"
	fmt.Printf("zhuge == *wolong: %t\n", zhuge == *wolong)

	//字段可以取地址
	fmt.Printf("zhuge.Name: %v\n", zhuge.Name)
	*(&zhuge.Name) = "南阳诸葛先生"
	fmt.Printf("zhuge.Name: %v\n", zhuge.Name)

	//结构体字面量
	fmt.Printf("tempStruct: %v\n", struct {
		fieldA int
		fieldB string
		fieldC []float64
	}{42, "foobar", []float64{1949, 9.9, math.Pi}})
	fmt.Printf("emptyStruct: %v\n", struct{}{})
}

type point struct {
	X, Y int
}
type circle struct {
	point
	radius int
}
type Wheel struct {
	circle
	Spokes int
}

func AnonymousField() {
	wheelFoo := Wheel{
		circle: circle{
			point: point{
				X: 2,
				Y: 3,
			},
			radius: 4,
		},
		Spokes: 7,
	}
	fmt.Printf("%#v\n", wheelFoo)

	wheelBar := Wheel{}
	//下三行当前pkg和其他pkg都可以使用
	wheelBar.X = 22
	wheelBar.Y = 33
	wheelBar.Spokes = 77
	//下四行只能在当前pkg使用
	wheelBar.circle.point.X = 22
	wheelBar.circle.point.Y = 33
	wheelBar.radius = 44
	wheelBar.circle.radius = 44
	fmt.Printf("%#v\n", wheelBar)

}
func Struct() {
	//TypeMethod()
	//StructBasics()
	//AnonymousField()
}
