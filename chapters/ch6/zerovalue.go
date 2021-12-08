package ch6

import (
	"fmt"
)

//一个type的零值是多少取决于这个type是什么具体类型的别名.
//
//如果具体类型是interface和引用类型(slice, map, pointer, channel, func),
//那么这个类型的零值就是nil;
//
//如果具体类型是基本数据类型, 那么零值就是对应类型的零值,
//比如 int: 0, float: 0., byte: '\0', rune: '\0', string: "", complex: 0+0i;
//
//如果具体类型是复合类型(array, struct), 零值就是所有字段都为零值的状态,
//任何时候一个复合类型实例都不是nil, 尝试用nil给一个复合类型变量赋值无法通过编译.

//nil is a predeclared identifier representing the zero value for
//a pointer, channel, func, interface, map, or slice type.
//Type must be a pointer, channel, func, interface, map, or slice type.

type AliasOfASlice []int

type AliasOfAStruct struct {
	FieldA int
	FieldB int
}

func ZeroValue() {

	//slice
	var p []int                            //nil
	q := []int{}                           //非nil
	j := make([]int, 0)                    //非nil
	fmt.Printf("p == nil: %t\n", p == nil) //true
	fmt.Printf("q == nil: %t\n", q == nil) //false
	fmt.Printf("j == nil: %t\n", j == nil) //false
	j = nil
	fmt.Printf("again: j == nil: %t\n", j == nil) //true
	fmt.Println()

	//type of slice
	var pp AliasOfASlice                     //nil
	qq := AliasOfASlice{}                    //非nil
	jj := make(AliasOfASlice, 0)             //非nil
	fmt.Printf("pp == nil: %t\n", pp == nil) //true
	fmt.Printf("qq == nil: %t\n", qq == nil) //false
	fmt.Printf("jj == nil: %t\n", jj == nil) //false
	jj = nil
	fmt.Printf("again: jj == nil: %t\n", jj == nil) //true
	fmt.Println()

	//type of struct
	var x AliasOfAStruct      //结构体实例, 零值, 非nil
	y := AliasOfAStruct{}     //结构体实例, 零值, 非nil
	pz := new(AliasOfAStruct) //指向结构体类型的指针实例, 非零值, 非nil; 被指向的结构体已经实例化, 是零值, 非nil.
	var pzz *AliasOfAStruct   //指向结构体类型的指针实例, 零值, nil; 被指向的结构体还未实例化, 是nil.

	////Cannot use 'nil' as the type AliasOfAStruct
	////Type must be a pointer, channel, func, interface, map, or slice type
	//x = nil

	////Cannot convert 'nil' to type 'AliasOfAStruct'
	////Type must be a pointer, channel, func, interface, map, or slice type
	//fmt.Printf("x == nil: %t\n", x == nil)

	////Cannot convert 'nil' to type 'AliasOfAStruct'
	////Type must be a pointer, channel, func, interface, map, or slice type
	//fmt.Printf("y == nil: %t\n", y == nil)

	fmt.Printf("pz == nil: %t\n", pz == nil) //false
	//pz: type *ch6.AliasOfAStruct value: &ch6.AliasOfAStruct{FieldA:0, FieldB:0} valueAsPointer: 0xc00000a0e0
	fmt.Printf("pz: type %T value: %#[1]v valueAsPointer: %[1]p\n", pz)

	fmt.Printf("pzz == nil: %t\n", pzz == nil) //true
	//pzz: type *ch6.AliasOfAStruct value: (*ch6.AliasOfAStruct)(nil) valueAsPointer: 0x0
	fmt.Printf("pzz: type %T value: %#[1]v valueAsPointer: %[1]p\n", pzz)
	fmt.Println()

	fmt.Printf("x: %#v\n", x)     //x: ch6.AliasOfAStruct{FieldA:0, FieldB:0}
	fmt.Printf("y: %#v\n", y)     //y: ch6.AliasOfAStruct{FieldA:0, FieldB:0}
	fmt.Printf("*pz: %#v\n", *pz) //*pz: ch6.AliasOfAStruct{FieldA:0, FieldB:0}
	//fmt.Printf("*pzz: %#v\n", *pzz) //panic: runtime error: invalid memory address or nil pointer dereference
	fmt.Println()
}

//1. nil channel 非常有用:
//	Given a nil channel c:
//		<-c receiving from c blocks forever
//		c <- v sending into c blocks forever
//		close(c) closing c panics
//	nil channel 常常用来禁用和激活select语句中的某个case
//2. make()出来的一个channel的值不是nil, 值为nil的channel是主动赋出来的
//3. 一个channel的类型不包含有无缓冲这一信息

func NilChannel() {
	//c == nil: false
	//c == nil: true
	//c == nil: false
	//c == nil: false
	//pushed.
	var c = make(chan int)
	fmt.Printf("c == nil: %v\n", c == nil)
	c = nil
	fmt.Printf("c == nil: %v\n", c == nil)
	c = make(chan int, 42)
	fmt.Printf("c == nil: %v\n", c == nil)
	c = make(chan int, 1)
	fmt.Printf("c == nil: %v\n", c == nil)
	c <- 100
	fmt.Printf("pushed.")
}
