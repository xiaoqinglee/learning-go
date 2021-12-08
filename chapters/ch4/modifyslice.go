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
