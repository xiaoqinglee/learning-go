package ch4

import "fmt"

func changeSliceInWrongWay(sliceStructReplica []int) {
	newSliceStruct := append(sliceStructReplica, 42)
	fmt.Printf("%#v\n", sliceStructReplica)
	fmt.Printf("%#v\n", newSliceStruct)
}
func changeSliceInRightWay(sliceStructReplica []int) []int {
	newSliceStruct := append(sliceStructReplica, 42)
	fmt.Printf("%#v\n", sliceStructReplica)
	fmt.Printf("%#v\n", newSliceStruct)
	return newSliceStruct
}

func changeMap(m map[int]string) {
	m[42] = "42"
}

func modifySliceAndMap() {
	sliceVar := []int{0, 1, 2, 3}
	fmt.Printf("main: 1. %#v\n", sliceVar)
	changeSliceInWrongWay(sliceVar)
	fmt.Printf("main: 2. %#v\n", sliceVar)
	sliceVar = changeSliceInRightWay(sliceVar)
	fmt.Printf("main: 3. %#v\n", sliceVar)
	fmt.Printf("--------------------\n")

	mapVar := make(map[int]string)
	fmt.Printf("main: 1. %#v\n", mapVar)
	changeMap(mapVar)
	fmt.Printf("main: 2. %#v\n", mapVar)
}
