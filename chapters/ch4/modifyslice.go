package ch4

import "fmt"

func changeSliceInWrongWay(sliceStructReplica []int) {
	newSliceStruct := append(sliceStructReplica, 21)
	fmt.Printf("%#v\n", sliceStructReplica)
	fmt.Printf("%#v\n", newSliceStruct)
}
func changeSliceInRightWay(sliceStructReplica []int) []int {
	newSliceStruct := append(sliceStructReplica, 42)
	fmt.Printf("%#v\n", sliceStructReplica)
	fmt.Printf("%#v\n", newSliceStruct)
	return newSliceStruct
}
func changeSliceInSecondRightWay(sliceStructPointerReplica *[]int) {
	*sliceStructPointerReplica = append(*sliceStructPointerReplica, 84)
	fmt.Printf("%#v\n", *sliceStructPointerReplica)
	return
}
func changeMap(m map[int]string) {
	m[42] = "42"
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
	--------------------
	outside: 1. map[int]string{}
	outside: 2. map[int]string{42:"42"}
*/
func ModifySliceAndMap() {
	sliceVar := []int{0, 1, 2, 3}
	fmt.Printf("outside: 1. %#v\n", sliceVar)
	changeSliceInWrongWay(sliceVar)
	fmt.Printf("outside: 2. %#v\n", sliceVar)
	sliceVar = changeSliceInRightWay(sliceVar)
	fmt.Printf("outside: 3. %#v\n", sliceVar)
	changeSliceInSecondRightWay(&sliceVar)
	fmt.Printf("outside: 4. %#v\n", sliceVar)
	fmt.Printf("--------------------\n")

	mapVar := make(map[int]string)
	fmt.Printf("outside: 1. %#v\n", mapVar)
	changeMap(mapVar)
	fmt.Printf("outside: 2. %#v\n", mapVar)
}

//order matters
func deleteElemAtIndex1(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

//order doesn't matter
func deleteElemAtIndex2(slice []int, i int) []int {
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
	nums = deleteElemAtIndex1(nums, 5)
	fmt.Println(nums)
	nums = deleteElemAtIndex2(nums, 3)
	fmt.Println(nums)

	fmt.Println("长度大于1的slice删除首尾元素:")
	nums = []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(nums)
	nums = deleteElemAtIndex1(nums, len(nums)-1)
	fmt.Println(nums)
	nums = deleteElemAtIndex1(nums, 0)
	fmt.Println(nums)
	fmt.Println("--------------------")
	nums = []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(nums)
	nums = deleteElemAtIndex2(nums, len(nums)-1)
	fmt.Println(nums)
	nums = deleteElemAtIndex2(nums, 0)
	fmt.Println(nums)

	fmt.Println("长度等于1的slice删除唯一元素:")
	nums = []int{0}
	fmt.Println(nums)
	nums = deleteElemAtIndex1(nums, len(nums)-1)
	fmt.Println(nums)

	nums = []int{0}
	fmt.Println(nums)
	nums = deleteElemAtIndex2(nums, len(nums)-1)
	fmt.Println(nums)
}
