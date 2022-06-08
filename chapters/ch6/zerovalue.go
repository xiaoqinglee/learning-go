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

//对于slice, map和channel：
//使用var初始化时，三者都是nil值。
//使用make()方式初始化时，三者都是非nil。
//	nil slice 可直接拿来用。
//	nil map 不能直接使用，会panic。
//	nil channel 可以直接使用，但是send和receive操作会永远阻塞。

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

/*

1. nil channel 非常有用:
	Given a nil channel c:
		<-c receiving from c blocks forever
		c <- v sending into c blocks forever
		close(c) closing c panics
	nil channel 常常用来禁用和激活select语句中的某个case
2. 向一个 channel 变量直接赋值 nil 可以得到 nil channel;
	var varName chan 得到的 channel 变量, 它的值也是 nil;
	make() 出来的一个 channel 变量, 它的值不是 nil.
3. 一个channel的类型不包含有无缓冲这一信息

ch := make(chan int, 5)
ch 是一个指向 runtime.hchan 结构体的指针. 详见ch4.

*/

func NilChannel() {
	//c == nil: false
	//c == nil: true
	//c == nil: false
	//c == nil: false
	//
	//c2 == nil: true
	//c2 == nil: false

	var c = make(chan int)
	fmt.Printf("c == nil: %v\n", c == nil)
	c = nil
	fmt.Printf("c == nil: %v\n", c == nil)
	c = make(chan int, 42)
	fmt.Printf("c == nil: %v\n", c == nil)
	c = make(chan int, 1)
	fmt.Printf("c == nil: %v\n", c == nil)
	fmt.Println()

	var c2 chan int
	fmt.Printf("c2 == nil: %v\n", c2 == nil)
	c2 = make(chan int, 42)
	fmt.Printf("c2 == nil: %v\n", c2 == nil)
}

func TestVarInitialization() {

	//var struct_ struct{}========================================================
	//struct_: {}
	//struct_ == struct{}{}: true
	fmt.Println("var struct_ struct{}========================================================")
	var struct_ struct{}
	fmt.Printf("struct_: %v\n", struct_)
	fmt.Printf("struct_ == struct{}{}: %v\n", struct_ == struct{}{})

	//var pointer *string========================================================
	//pointer: <nil>
	//pointer == nil: true
	fmt.Println("var pointer *string========================================================")
	var pointer *string
	fmt.Printf("pointer: %v\n", pointer)
	fmt.Printf("pointer == nil: %v\n", pointer == nil)

	//var slice1 []int========================================================
	//slice1: []
	//slice1 == nil: true
	//len(slice1) == 0：true
	//after appending
	//slice1: [42]
	//slice1 == nil: false
	//len(slice1) == 0：false
	fmt.Println("var slice1 []int========================================================")
	var slice1 []int
	fmt.Printf("slice1: %v\n", slice1)
	fmt.Printf("slice1 == nil: %v\n", slice1 == nil)
	fmt.Printf("len(slice1) == 0：%v\n", len(slice1) == 0)
	slice1 = append(slice1, 42)
	fmt.Printf("after appending\n")
	fmt.Printf("slice1: %v\n", slice1)
	fmt.Printf("slice1 == nil: %v\n", slice1 == nil)
	fmt.Printf("len(slice1) == 0：%v\n", len(slice1) == 0)

	//slice2 := make([]int, 0)--------------------------------------------------------
	//slice2: []
	//slice2 == nil: false
	//len(slice2) == 0：true
	//after appending
	//slice2: [42]
	//slice2 == nil: false
	//len(slice2) == 0：false
	fmt.Println("slice2 := make([]int, 0)--------------------------------------------------------")
	slice2 := make([]int, 0)
	fmt.Printf("slice2: %v\n", slice2)
	fmt.Printf("slice2 == nil: %v\n", slice2 == nil)
	fmt.Printf("len(slice2) == 0：%v\n", len(slice2) == 0)
	slice2 = append(slice2, 42)
	fmt.Printf("after appending\n")
	fmt.Printf("slice2: %v\n", slice2)
	fmt.Printf("slice2 == nil: %v\n", slice2 == nil)
	fmt.Printf("len(slice2) == 0：%v\n", len(slice2) == 0)

	//var map_1 map[int]int========================================================
	//map_1: map[]
	//map_1 == nil: true
	//len(map_1) == 0：true
	//panic: assignment to entry in nil map
	fmt.Println("var map_1 map[int]int========================================================")
	var map_1 map[int]int
	fmt.Printf("map_1: %v\n", map_1)
	fmt.Printf("map_1 == nil: %v\n", map_1 == nil)
	fmt.Printf("len(map_1) == 0：%v\n", len(map_1) == 0)
	//map_1[42] = 42

	//map_2 := make(map[int]int)--------------------------------------------------------
	//map_2: map[]
	//map_2 == nil: false
	//len(map_2) == 0：true
	//after appending
	//map_2: map[42:42]
	//map_2 == nil: false
	//len(map_2) == 0：false
	fmt.Println("map_2 := make(map[int]int)--------------------------------------------------------")
	map_2 := make(map[int]int)
	fmt.Printf("map_2: %v\n", map_2)
	fmt.Printf("map_2 == nil: %v\n", map_2 == nil)
	fmt.Printf("len(map_2) == 0：%v\n", len(map_2) == 0)
	map_2[42] = 42
	fmt.Printf("after appending\n")
	fmt.Printf("map_2: %v\n", map_2)
	fmt.Printf("map_2 == nil: %v\n", map_2 == nil)
	fmt.Printf("len(map_2) == 0：%v\n", len(map_2) == 0)

	//var chan1 chan int========================================================
	//chan1: <nil>
	//chan1 == nil: true
	//len(chan1) == 0：true
	//waiting for receiving
	//fatal error: all goroutines are asleep - deadlock!
	//
	//goroutine 1 [chan send (nil chan)]...
	fmt.Println("var chan1 chan int========================================================")
	var chan1 chan int
	fmt.Printf("chan1: %v\n", chan1)
	fmt.Printf("chan1 == nil: %v\n", chan1 == nil)
	fmt.Printf("len(chan1) == 0：%v\n", len(chan1) == 0)
	//go func() {
	//	fmt.Println("waiting for receiving")
	//	<-chan1
	//	fmt.Println("receiving done")
	//}()
	//chan1 <- 42

	//chan2 := make(chan int, 0)--------------------------------------------------------
	//chan2: 0xc000052120
	//chan2 == nil: false
	//len(chan2) == 0：true
	//waiting for receiving
	//receiving done
	//after appending
	//chan2: 0xc000052120
	//chan2 == nil: false
	//len(chan2) == 0：true
	fmt.Println("chan2 := make(chan int, 0)--------------------------------------------------------")
	chan2 := make(chan int, 0)
	fmt.Printf("chan2: %v\n", chan2)
	fmt.Printf("chan2 == nil: %v\n", chan2 == nil)
	fmt.Printf("len(chan2) == 0：%v\n", len(chan2) == 0)
	go func() {
		fmt.Println("waiting for receiving")
		<-chan2
		fmt.Println("receiving done")
	}()
	chan2 <- 42
	fmt.Printf("after appending\n")
	fmt.Printf("chan2: %v\n", chan2)
	fmt.Printf("chan2 == nil: %v\n", chan2 == nil)
	fmt.Printf("len(chan2) == 0：%v\n", len(chan2) == 0)
}
