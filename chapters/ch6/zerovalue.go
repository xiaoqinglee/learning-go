package ch6

import (
	"fmt"
)

//一个type的零值是多少取决于这个type是什么具体类型的别名.
//如果具体类型是pointer, channel, func, interface, map, or slice type,
//那么这个类型的零值就是nil;
//如果具体类型是其他类型, 那么零值就是对应类型的零值.
//struct类型的零值状态是所有字段都为零值的状态, 此时整个struct不是nil,
//任何时候一个struct实例都不是nil, 尝试用nil给一个struct实例赋值无法通过编译.

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
