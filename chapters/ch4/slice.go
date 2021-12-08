package ch4

import (
	"fmt"
	"reflect"
)

func Slice() {
	months := [1 + 12]string{
		"", "Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
	}
	summer := months[6:9] //6,7,8月
	Q2 := months[4:7]     //4,5,6月
	fmt.Printf("month: %v, len: %d, type: %[1]T\n", months, len(months))
	fmt.Printf("summer: %v, len: %d, cap: %d, type: %[1]T\n", summer, len(summer), cap(summer))
	fmt.Printf("Q2: %v, len: %d, cap: %d, type: %[1]T\n", Q2, len(Q2), cap(Q2))
	fmt.Println()

	//initiation
	var s []int
	fmt.Printf("len(s) == 0: %t s == nil: %t\n", len(s) == 0, s == nil)
	s = nil
	fmt.Printf("len(s) == 0: %t s == nil: %t\n", len(s) == 0, s == nil)
	s = []int(nil)
	fmt.Printf("len(s) == 0: %t s == nil: %t\n", len(s) == 0, s == nil)
	s = []int{}
	fmt.Printf("len(s) == 0: %t s == nil: %t\n", len(s) == 0, s == nil)
	s = make([]int, 0) //蛋疼的写法
	fmt.Printf("len(s) == 0: %t s == nil: %t\n", len(s) == 0, s == nil)
	s = make([]int, 0, 0) //蛋疼的写法
	fmt.Printf("len(s) == 0: %t s == nil: %t\n", len(s) == 0, s == nil)
	fmt.Println()

	//comparison
	x := []int{1, 2, 3}
	y := []int{1, 2, 3}
	var p []int         //nil
	q := []int{}        //非nil
	j := make([]int, 0) //非nil
	fmt.Printf("x == nil: %t\n", x == nil)
	fmt.Printf("y == nil: %t\n", y == nil)
	fmt.Printf("p == nil: %t\n", p == nil) //true
	fmt.Printf("q == nil: %t\n", q == nil) //false
	fmt.Printf("j == nil: %t\n", j == nil) //false
	j = nil
	fmt.Printf("again: j == nil: %t\n", j == nil) //true

	//Invalid operation: x == y (the operator == is not defined on []int)
	//fmt.Printf("x == y: %t", x == y)
	fmt.Printf("reflect.DeepEqual(x,y): %t\n", reflect.DeepEqual(x, y)) //true
	fmt.Printf("p: %v\n", p)
	fmt.Printf("q: %v\n", q)
	fmt.Printf("j: %v\n", j)

	fmt.Printf("reflect.DeepEqual(p,p): %t\n", reflect.DeepEqual(p, p))     // true
	fmt.Printf("reflect.DeepEqual(p,nil): %t\n", reflect.DeepEqual(p, nil)) //!!! false
	fmt.Printf("reflect.DeepEqual(p,q): %t\n", reflect.DeepEqual(p, q))     //!!! false

	fmt.Printf("reflect.DeepEqual(q,j): %t\n", reflect.DeepEqual(q, j)) //true
	fmt.Println()

	//appending
	x = []int{1, 2, 3}
	y = []int{7, 8}
	fmt.Printf("%p -> %v\n", &x, x)
	fmt.Printf("len(x): %d\n", len(x)) //3
	fmt.Printf("cap(x): %d\n", cap(x)) //3
	fmt.Printf("%p -> %v\n", &y, y)

	x = append(x, 4, 5, 6)
	x = append(x, y...)
	fmt.Printf("%p -> %v\n", &x, x)
	fmt.Printf("len(x): %d\n", len(x)) //8
	fmt.Printf("cap(x): %d\n", cap(x)) //12
	fmt.Printf("%p -> %v\n", &y, y)
	fmt.Println()

	//copying, 取两个参数len中较小的那个, 永远不会出现越界
	x = []int{11, 22}
	y = []int{66, 77, 88}
	fmt.Printf("%p -> %v\n", &x, x)
	fmt.Printf("%p -> %v\n", &y, y)
	copy(x, y)
	fmt.Printf("%p -> %v\n", &x, x)
	fmt.Printf("%p -> %v\n", &y, y)

	x = []int{11, 22, 33}
	y = []int{66, 77}
	fmt.Printf("%p -> %v\n", &x, x)
	fmt.Printf("%p -> %v\n", &y, y)
	copy(x, y)
	fmt.Printf("%p -> %v\n", &x, x)
	fmt.Printf("%p -> %v\n", &y, y)
	fmt.Println()
}

/*
	go中索引操作的范围为闭区间[0,len(x)-1], 但是切片时允许使用[len(x):]获得一个空切片
		slice1: []int{}
		slice2: []int{}

*/
func SliceIndex() {
	array1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	var slice1 []int = array1[len(array1):]
	fmt.Printf("slice1: %#v\n", slice1)
	var slice2 []int = slice1[len(slice1):]
	fmt.Printf("slice2: %#v\n", slice2)
}

/*
	append 会在原slice变量的底层数组上进行原位操作, 而不是初始化一个新的底层数组.
		array1: [8]int{0, 1, 2, 3, 4, 5, 6, 7}
		slice1: []int{0, 1, 2, 3, 4, 5, 6, 7}
		after deletion:
		array1: [8]int{0, 1, 2, 4, 5, 6, 7, 7}
		slice1: []int{0, 1, 2, 4, 5, 6, 7, 7} <slice1已经改变>
		slice2: []int{0, 1, 2, 4, 5, 6, 7}
		----------------------------------
		array1: [8]int{0, 1, 2, 3, 4, 5, 6, 7}
		slice1: []int{0, 1, 2, 3, 4, 5, 6, 7}
		after deletion:
		array1: [8]int{0, 1, 2, 3, 4, 5, 6, 7}
		slice1: []int{0, 1, 2, 3, 4, 5, 6, 7} <slice1没有改变>
		slice2: []int{0, 1, 2, 4, 5, 6, 7}
*/
func UnderSlice() {
	array1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	var slice1 []int = array1[:]
	fmt.Printf("array1: %#v\n", array1)
	fmt.Printf("slice1: %#v\n", slice1)

	//删除索引为3的元素
	var slice2 []int = append(slice1[:3], slice1[3+1:]...)
	fmt.Printf("after deletion:\n")
	fmt.Printf("array1: %#v\n", array1)
	fmt.Printf("slice1: %#v\n", slice1)
	fmt.Printf("slice2: %#v\n", slice2)
	fmt.Printf("----------------------------------\n")

	array1 = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	slice1 = array1[:]
	fmt.Printf("array1: %#v\n", array1)
	fmt.Printf("slice1: %#v\n", slice1)

	//避免原位操作的方式
	slice2 = make([]int, len(slice1), cap(slice1))
	copy(slice2, slice1)
	slice2 = append(slice2[:3], slice2[3+1:]...)
	fmt.Printf("after deletion:\n")
	fmt.Printf("array1: %#v\n", array1)
	fmt.Printf("slice1: %#v\n", slice1)
	fmt.Printf("slice2: %#v\n", slice2)
}
