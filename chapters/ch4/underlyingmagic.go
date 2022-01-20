package ch4

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
	ref https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-make-and-new/

	当我们想要在 Go 语言中初始化一个结构时，可能会用到两个不同的关键字 — make 和 new。因为它们的功能相似，
	所以初学者可能会对这两个关键字的作用感到困惑，但是它们两者能够初始化的变量却有较大的不同。

	make 的作用是初始化内置的数据结构，也就是我们在前面提到的切片、哈希表和 Channel；
	new 的作用是根据传入的类型分配一片内存空间并返回指向这片内存空间的指针；

	我们在代码中往往都会使用如下所示的语句初始化这三类基本类型，这三个语句分别返回了不同类型的数据结构：

		slice := make([]int, 0, 100)
		hash := make(map[int]bool, 10)
		ch := make(chan int, 5)

		slice 是一个包含 data、cap 和 len 的结构体 reflect.SliceHeader；
		hash 是一个指向 runtime.hmap 结构体的指针；
		ch 是一个指向 runtime.hchan 结构体的指针；

	相比与复杂的 make 关键字，new 的功能就简单多了，它只能接收类型作为参数然后返回一个指向该类型的指针：

		i := new(int)

		var v int
		i := &v

	上述代码片段中的两种不同初始化方法是等价的，它们都会创建一个指向 int 零值的指针。
*/

/*
	ref https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-string/

	字符串在 Go 语言中的接口其实非常简单，每一个字符串在运行时都会使用如下的 reflect.StringHeader 表示，
	其中包含指向字节数组的指针和数组的大小：

		type StringHeader struct {
			Data uintptr
			Len  int
		}

	与切片的结构体相比，字符串只少了一个表示容量的 Cap 字段，而正是因为切片在 Go 语言的运行时表示与字符串高度相似，
	所以我们经常会说字符串是一个只读的切片类型。

		type SliceHeader struct {
			Data uintptr
			Len  int
			Cap  int
		}

	因为字符串作为只读的类型，我们并不会直接向字符串直接追加元素改变其本身的内存空间，
	所有在字符串上的写入操作都是通过拷贝实现的。
*/

/*
	ref https://stackoverflow.com/questions/36706843/how-to-get-the-underlying-array-of-a-slice-in-go

	To access the underlying array, you can use a combination of reflect and unsafe.
	In particular, reflect.SliceHeader contains a Data field which contains a pointer
	to the underlying array of a slice.

	Example adapted from the documentation of the unsafe package:

		s := []int{1, 2, 3, 4}
		hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
		data := *(*[4]int)(unsafe.Pointer(hdr.Data))
*/

func GetSliceUnderlyingArrayAddress(slice []int) unsafe.Pointer {
	slicePointer := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	return unsafe.Pointer(slicePointer.Data)
}

/*
	共享底层数组:
	slice1: []int{0, 11, 2}, slice2: []int{0, 11, 2}
	slice1 underlying array's address: (unsafe.Pointer)(0xc0000ae090)
	slice2 underlying array's address: (unsafe.Pointer)(0xc0000ae090)
	--------------------------------------------------------------
	共享底层数组:
	array1: [3]int{0, 11, 2}, slice1: []int{0, 11, 2}, slice2: []int{0, 11, 2}
	array1's address: 0xc0000ae0a8
	slice1 underlying array's address: (unsafe.Pointer)(0xc0000ae0a8)
	slice2 underlying array's address: (unsafe.Pointer)(0xc0000ae0a8)
	--------------------------------------------------------------
	append 操作没有在底层数组上发生越界, 所以 slice1 slice2 仍然继续共享底层数组:
	array1: [3]int{0, 1, 2}, slice1: []int{0, 1}, slice2: []int{0, 1, 2}
	array1's address: 0xc0000ae0a8
	slice1 underlying array's address: (unsafe.Pointer)(0xc0000ae0a8)
	slice2 underlying array's address: (unsafe.Pointer)(0xc0000ae0a8)
	array1: [3]int{0, 1, 22}, slice1: []int{0, 1, 22}, slice2: []int{0, 1, 22}
	array1's address: 0xc0000ae0a8
	slice1 underlying array's address: (unsafe.Pointer)(0xc0000ae0a8)
	slice2 underlying array's address: (unsafe.Pointer)(0xc0000ae0a8)
	--------------------------------------------------------------
	append 操作超出了底层数组的边界, 所以 slice1 的底层数组自立门户:
	array1: [3]int{0, 1, 2}, slice1: []int{0, 1}, slice2: []int{0, 1, 2}
	array1's address: 0xc0000ae0a8
	slice1 underlying array's address: (unsafe.Pointer)(0xc0000ae0a8)
	slice2 underlying array's address: (unsafe.Pointer)(0xc0000ae0a8)
	array1: [3]int{0, 1, 2}, slice1: []int{0, 1, 22, 33}, slice2: []int{0, 1, 2}
	array1's address: 0xc0000ae0a8
	slice1 underlying array's address: (unsafe.Pointer)(0xc0000cc060)
	slice2 underlying array's address: (unsafe.Pointer)(0xc0000ae0a8)
	array1: [3]int{42, 1, 2}, slice1: []int{0, 1, 22, 33}, slice2: []int{42, 1, 2}
	array1's address: 0xc0000ae0a8
	slice1 underlying array's address: (unsafe.Pointer)(0xc0000cc060)
	slice2 underlying array's address: (unsafe.Pointer)(0xc0000ae0a8)
	--------------------------------------------------------------
*/
func TestSliceAssignment() {
	fmt.Printf("共享底层数组:\n")
	slice1 := []int{0, 1, 2}
	slice2 := slice1
	slice1[1] = 11
	fmt.Printf("slice1: %#v, slice2: %#v\n", slice1, slice2)
	fmt.Printf("slice1 underlying array's address: %#v\n", GetSliceUnderlyingArrayAddress(slice1))
	fmt.Printf("slice2 underlying array's address: %#v\n", GetSliceUnderlyingArrayAddress(slice2))
	fmt.Printf("--------------------------------------------------------------\n")

	fmt.Printf("共享底层数组:\n")
	array1 := [...]int{0, 1, 2}
	slice1 = array1[:]
	slice2 = slice1
	slice1[1] = 11
	fmt.Printf("array1: %#v, slice1: %#v, slice2: %#v\n", array1, slice1, slice2)
	fmt.Printf("array1's address: %p\n", &array1)
	fmt.Printf("slice1 underlying array's address: %#v\n", GetSliceUnderlyingArrayAddress(slice1))
	fmt.Printf("slice2 underlying array's address: %#v\n", GetSliceUnderlyingArrayAddress(slice2))
	fmt.Printf("--------------------------------------------------------------\n")

	fmt.Printf("append 操作没有在底层数组上发生越界, 所以 slice1 slice2 仍然继续共享底层数组:\n")
	array1 = [...]int{0, 1, 2}
	slice1 = array1[:2] // 0, 1
	slice2 = array1[:]  // 0, 1, 2
	fmt.Printf("array1: %#v, slice1: %#v, slice2: %#v\n", array1, slice1, slice2)
	fmt.Printf("array1's address: %p\n", &array1)
	fmt.Printf("slice1 underlying array's address: %#v\n", GetSliceUnderlyingArrayAddress(slice1))
	fmt.Printf("slice2 underlying array's address: %#v\n", GetSliceUnderlyingArrayAddress(slice2))
	slice1 = append(slice1, 22)
	fmt.Printf("array1: %#v, slice1: %#v, slice2: %#v\n", array1, slice1, slice2)
	fmt.Printf("array1's address: %p\n", &array1)
	fmt.Printf("slice1 underlying array's address: %#v\n", GetSliceUnderlyingArrayAddress(slice1))
	fmt.Printf("slice2 underlying array's address: %#v\n", GetSliceUnderlyingArrayAddress(slice2))
	fmt.Printf("--------------------------------------------------------------\n")

	fmt.Printf("append 操作超出了底层数组的边界, 所以 slice1 的底层数组自立门户:\n")
	array1 = [...]int{0, 1, 2}
	slice1 = array1[:2] // 0, 1
	slice2 = array1[:]  // 0, 1, 2
	fmt.Printf("array1: %#v, slice1: %#v, slice2: %#v\n", array1, slice1, slice2)
	fmt.Printf("array1's address: %p\n", &array1)
	fmt.Printf("slice1 underlying array's address: %#v\n", GetSliceUnderlyingArrayAddress(slice1))
	fmt.Printf("slice2 underlying array's address: %#v\n", GetSliceUnderlyingArrayAddress(slice2))

	//因为最终array1的最终结果为[3]int{0, 1, 2}, 而不是[3]int{0, 1, 22}
	//所以我们可以知道append函数会先计算append动作结束后底层数组右边界是否有越界从而决定是否新开辟底层数组,
	//然后开始追加所有元素.
	//而不是将多个elem逐一append到旧的底层数组边追加边检测是否越界.
	slice1 = append(slice1, 22, 33)
	fmt.Printf("array1: %#v, slice1: %#v, slice2: %#v\n", array1, slice1, slice2)
	fmt.Printf("array1's address: %p\n", &array1)
	fmt.Printf("slice1 underlying array's address: %#v\n", GetSliceUnderlyingArrayAddress(slice1))
	fmt.Printf("slice2 underlying array's address: %#v\n", GetSliceUnderlyingArrayAddress(slice2))

	array1[0] = 42 //再次检查
	fmt.Printf("array1: %#v, slice1: %#v, slice2: %#v\n", array1, slice1, slice2)
	fmt.Printf("array1's address: %p\n", &array1)
	fmt.Printf("slice1 underlying array's address: %#v\n", GetSliceUnderlyingArrayAddress(slice1))
	fmt.Printf("slice2 underlying array's address: %#v\n", GetSliceUnderlyingArrayAddress(slice2))
	fmt.Printf("--------------------------------------------------------------\n")

}

func AddOrDeleteMapElem(m map[int]string) {
	m[42] = "42"
}

/*
	outside: m value: map[42:42]
	outside: m as pointer: 0xc00007a480 &m: 0xc000006028
	inside: m value: map[42:42]
	inside: m as pointer: 0xc00007a480 &m: 0xc000006038
	inside: m value: map[]
	inside: m as pointer: 0xc00007a510 &m: 0xc000006038
	outside: m value: map[42:42]
	outside: m as pointer: 0xc00007a480 &m: 0xc000006028
*/
func ModifyMapInWrongWay() {
	//map类型是pointer, pointer做参数时会发生值拷贝, 在函数栈上初始化新的pointer变量. 所以inside outside &m不一样.
	//m是指针, &m是指针变量的地址, 也即指向指针的指针.
	m := make(map[int]string)
	m[42] = "42"
	fmt.Printf("outside: m value: %v\n", m)
	fmt.Printf("outside: m as pointer: %p &m: %p\n", m, &m)
	func(m map[int]string) {
		fmt.Printf("inside: m value: %v\n", m)
		fmt.Printf("inside: m as pointer: %p &m: %p\n", m, &m)
		m = make(map[int]string)
		fmt.Printf("inside: m value: %v\n", m)
		fmt.Printf("inside: m as pointer: %p &m: %p\n", m, &m)
	}(m)
	fmt.Printf("outside: m value: %v\n", m)
	fmt.Printf("outside: m as pointer: %p &m: %p\n", m, &m)
}

/*
	outside: m value: map[42:42]
	outside: m as pointer: 0xc00007a480 &m: 0xc000006028
	inside: m value: map[42:42]
	inside: m as pointer: 0xc00007a480 &m: 0xc000006028
	inside: m value: map[]
	inside: m as pointer: 0xc00007a510 &m: 0xc000006028
	outside: m value: map[]
	outside: m as pointer: 0xc00007a510 &m: 0xc000006028

*/
func ModifyMapInRightWay1() { //inside outside &m一样.
	m := make(map[int]string)
	m[42] = "42"
	fmt.Printf("outside: m value: %v\n", m)
	fmt.Printf("outside: m as pointer: %p &m: %p\n", m, &m)
	func(pm *map[int]string) {
		fmt.Printf("inside: m value: %v\n", *pm)
		fmt.Printf("inside: m as pointer: %p &m: %p\n", *pm, pm)
		*pm = make(map[int]string)
		fmt.Printf("inside: m value: %v\n", *pm)
		fmt.Printf("inside: m as pointer: %p &m: %p\n", *pm, pm)
	}(&m)
	fmt.Printf("outside: m value: %v\n", m)
	fmt.Printf("outside: m as pointer: %p &m: %p\n", m, &m)
}

/*
	outside: m value: map[42:42]
	outside: m as pointer: 0xc0000c2450 &m: 0xc0000ce018
	inside: m value: map[42:42]
	inside: m as pointer: 0xc0000c2450 &m: 0xc0000ce028
	inside: m value: map[]
	inside: m as pointer: 0xc0000c24e0 &m: 0xc0000ce028
	outside: m value: map[]
	outside: m as pointer: 0xc0000c24e0 &m: 0xc0000ce018
*/
func ModifyMapInRightWay2() { //inside outside &m不一样
	m := make(map[int]string)
	m[42] = "42"
	fmt.Printf("outside: m value: %v\n", m)
	fmt.Printf("outside: m as pointer: %p &m: %p\n", m, &m)
	m = func(m map[int]string) map[int]string {
		fmt.Printf("inside: m value: %v\n", m)
		fmt.Printf("inside: m as pointer: %p &m: %p\n", m, &m)
		m = make(map[int]string)
		fmt.Printf("inside: m value: %v\n", m)
		fmt.Printf("inside: m as pointer: %p &m: %p\n", m, &m)
		return m
	}(m)
	fmt.Printf("outside: m value: %v\n", m)
	fmt.Printf("outside: m as pointer: %p &m: %p\n", m, &m)
}
