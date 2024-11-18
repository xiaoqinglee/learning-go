package ch4

import "fmt"

func modifySliceInWrongWay(sliceStructReplica []int) {
	newSliceStruct := append(sliceStructReplica, 21)
	fmt.Printf("%#v\n", sliceStructReplica)
	fmt.Printf("%#v\n", newSliceStruct)
}
func modifySliceInRightWay1(sliceStructReplica []int) []int {
	newSliceStruct := append(sliceStructReplica, 42)
	fmt.Printf("%#v\n", sliceStructReplica)
	fmt.Printf("%#v\n", newSliceStruct)
	return newSliceStruct
}
func modifySliceInRightWay2(sliceStructPointerReplica *[]int) {
	*sliceStructPointerReplica = append(*sliceStructPointerReplica, 84)
	fmt.Printf("%#v\n", *sliceStructPointerReplica)
	return
}

/*
outside: 1. []int{0, 1, 2, 3}
[]int{0, 1, 2, 3}
[]int{0, 1, 2, 3, 21}
outside: 2. []int{0, 1, 2, 3}
[]int{0, 1, 2, 3}
[]int{0, 1, 2, 3, 42}
outside: 3. []int{0, 1, 2, 3, 42}
[]int{0, 1, 2, 3, 42, 84}
outside: 4. []int{0, 1, 2, 3, 42, 84}
*/
func ModifySlice() {
	sliceVar := []int{0, 1, 2, 3}
	fmt.Printf("outside: 1. %#v\n", sliceVar)
	modifySliceInWrongWay(sliceVar)
	fmt.Printf("outside: 2. %#v\n", sliceVar)
	sliceVar = modifySliceInRightWay1(sliceVar)
	fmt.Printf("outside: 3. %#v\n", sliceVar)
	modifySliceInRightWay2(&sliceVar)
	fmt.Printf("outside: 4. %#v\n", sliceVar)
}

// order matters
func deleteElemAtIndexOrdered(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

// order doesn't matter
func deleteElemAtIndexUnordered(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

/*
长度大于1的slice删除中间元素:
[0 1 2 3 4 5 6 7]
[0 1 2 3 4 6 7]
[0 1 2 7 4 6]
长度大于1的slice删除首尾元素:
[0 1 2 3 4 5 6 7]
[0 1 2 3 4 5 6]
[1 2 3 4 5 6]
--------------------
[0 1 2 3 4 5 6 7]
[0 1 2 3 4 5 6]
[6 1 2 3 4 5]
长度等于1的slice删除唯一元素:
[0]
[]
[0]
[]
*/
func DeleteElemAtIndex() {
	fmt.Println("长度大于1的slice删除中间元素:")
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(nums)
	nums = deleteElemAtIndexOrdered(nums, 5)
	fmt.Println(nums)
	nums = deleteElemAtIndexUnordered(nums, 3)
	fmt.Println(nums)

	fmt.Println("长度大于1的slice删除首尾元素:")
	nums = []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(nums)
	nums = deleteElemAtIndexOrdered(nums, len(nums)-1)
	fmt.Println(nums)
	nums = deleteElemAtIndexOrdered(nums, 0)
	fmt.Println(nums)
	fmt.Println("--------------------")
	nums = []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(nums)
	nums = deleteElemAtIndexUnordered(nums, len(nums)-1)
	fmt.Println(nums)
	nums = deleteElemAtIndexUnordered(nums, 0)
	fmt.Println(nums)

	fmt.Println("长度等于1的slice删除唯一元素:")
	nums = []int{0}
	fmt.Println(nums)
	nums = deleteElemAtIndexOrdered(nums, len(nums)-1)
	fmt.Println(nums)

	nums = []int{0}
	fmt.Println(nums)
	nums = deleteElemAtIndexUnordered(nums, len(nums)-1)
	fmt.Println(nums)
}

/*
go中索引操作的范围为闭区间[0,len(x)-1],
但是切片时允许使用[len(x):]获得尾部一个空切片, 允许使用[:0]获得一个头部的空切片
[len(x):]等价于[len(x):len(x)], [:0]等价于[0:0]

	array1: [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	slice1: []int{0, 1, 2, 3, 4, 5, 6, 7}
	-----------------
	slice2: []int{}
	slice3: []int{}
	slice4: []int{}
	slice5: []int{}
*/
func SliceIndex() {
	array1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	slice1 := array1[:]
	fmt.Printf("array1: %#v\n", array1)
	fmt.Printf("slice1: %#v\n", slice1)
	fmt.Printf("-----------------\n")

	var slice2 []int = array1[len(array1):]
	fmt.Printf("slice2: %#v\n", slice2)

	var slice3 []int = slice1[len(slice1):]
	fmt.Printf("slice3: %#v\n", slice3)

	var slice4 []int = array1[:0]
	fmt.Printf("slice4: %#v\n", slice4)

	var slice5 []int = slice1[:0]
	fmt.Printf("slice5: %#v\n", slice5)
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
func UnderSliceAppend() {
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

/**
切片高级语法: sliceFoo[lowIndex:highIndex:maxIndex]

对于 sliceFoo[lowIndex:highIndex]
lowIndex <= highIndex <= len(sliceFoo)

sliceFoo[lowIndex:highIndex] 等价于
sliceFoo[lowIndex:highIndex:cap(sliceFoo)]

对于 sliceFoo[lowIndex:highIndex:maxIndex]
lowIndex <= highIndex <= maxIndex <= cap(sliceFoo)
*/

func AdvancedSlice() {
	numbers := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var s1, s2 []int
	s1 = numbers[1:4]
	fmt.Printf("len:  %v, cap: %v\n", len(s1), cap(s1)) //len:  3, cap: 9
	s2 = numbers[1:4:4]
	fmt.Printf("len:  %v, cap: %v\n", len(s2), cap(s2)) //len:  3, cap: 3
}

//also see:
//https://pkg.go.dev/slices#Concat
//https://pkg.go.dev/slices#AppendSeq
//https://pkg.go.dev/slices#Delete
//https://pkg.go.dev/slices#Clone
