package ch6

import (
	"fmt"
	"strconv"
)

type MyInt int

func (p *MyInt) Increment() {
	*p = MyInt(int(*p) + 1)
}
func (p *MyInt) IncrementBy(by MyInt) {
	*p = MyInt(int(*p) + int(by))
}
func (p *MyInt) Plus(other MyInt) MyInt {
	return MyInt(int(*p) + int(other))
}
func (p *MyInt) String() string {
	return strconv.Itoa(int(*p))
}

func MethodVariable() {
	//绑定了receiver实例的方法变量
	var p *MyInt
	//p == nil: true
	//p: type *ch6.MyInt value: (*ch6.MyInt)(nil) valueAsPointer: 0x0
	fmt.Printf("p == nil: %t\n", p == nil)
	fmt.Printf("p: type %T value: %#[1]v valueAsPointer: %[1]p\n", p)
	myInt := MyInt(42)
	p = &myInt
	//p == nil: false
	//p: type *ch6.MyInt value: (*ch6.MyInt)(0xc00000a0d0) valueAsPointer: 0xc00000a0d0
	fmt.Printf("p == nil: %t\n", p == nil)
	fmt.Printf("p: type %T value: %#[1]v valueAsPointer: %[1]p\n", p)
	fmt.Println()

	pIncrement := p.Increment
	pIncrementBy := p.IncrementBy
	pPlus := p.Plus
	fmt.Printf("type: %T value: %[1]p\n", pIncrement)
	fmt.Printf("type: %T value: %[1]p\n", pIncrementBy)
	fmt.Printf("type: %T value: %[1]p\n", pPlus)
	fmt.Printf("*p = %s\n", p.String())
	pIncrement()
	fmt.Printf("*p = %s\n", p.String())
	pIncrementBy(MyInt(-42))
	fmt.Printf("*p = %s\n", p.String())
	fmt.Printf("pPlus(MyInt(10)) = %#v\n", pPlus(MyInt(10)))
	fmt.Println()

	//未绑定receiver实例的方法变量
	paramIncrement := (*MyInt).Increment
	paramIncrementBy := (*MyInt).IncrementBy
	paramPlus := (*MyInt).Plus
	fmt.Printf("type: %T value: %[1]p\n", paramIncrement)
	fmt.Printf("type: %T value: %[1]p\n", paramIncrementBy)
	fmt.Printf("type: %T value: %[1]p\n", paramPlus)

	myInt = MyInt(42)
	p2 := &myInt
	fmt.Printf("*p2 = %s\n", p2.String())
	paramIncrement(p2)
	fmt.Printf("*p2 = %s\n", p2.String())
	paramIncrementBy(p2, MyInt(-42))
	fmt.Printf("*p2 = %s\n", p2.String())
	fmt.Printf("paramPlus(p2, MyInt(10)) = %#v\n", paramPlus(p2, MyInt(10)))
	fmt.Println()
}
