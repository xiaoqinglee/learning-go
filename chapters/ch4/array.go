package ch4

import "fmt"

func Array() {
	var a [3]int
	fmt.Printf("a[3]int: type: %T, value: %[1]v\n", a)
	var b [2]int
	fmt.Printf("b[2]int: type: %T, value: %[1]v\n", b)

	//initiation
	c := [3]int{11, 22, 33}
	fmt.Printf("c: type: %T, value: %[1]v\n", c)
	d := [3]int{111, 222}
	fmt.Printf("d: type: %T, value: %[1]v\n", d)
	e := [...]int{111, 222}
	fmt.Printf("e: type: %T, value: %[1]v\n", e)
	slice := []int{111, 222} //这不是数组
	fmt.Printf("slice: type: %T, value: %[1]v\n", slice)
	f := [...]int{9: 42}
	fmt.Printf("f: type: %T, value: %[1]v\n", f)
	g := [10]int{}
	//g[-1] = 42 //Invalid array index '-1' (must be non-negative)
	g[len(g)-1] = 42
	fmt.Printf("g: type: %T, value: %[1]v\n", g)
	fmt.Printf("&g: %[1]p\n", &g)
	fmt.Printf("&f: %[1]p\n", &f)
	fmt.Printf("two addresses &g == &f: %[1]v\n", &g == &f)
	fmt.Printf("two arrays g == f: %[1]v\n", g == f)
	////Invalid operation: e == slice (mismatched types [2]int and []int)
	//fmt.Printf("array and slice e == slice: %[1]v\n", e == slice)
	fmt.Println()

	//passing replica vs passing ref
	array := [3]int{11, 22, 33}
	fmt.Printf("array: %v\n", array)
	fmt.Printf("array address: %p\n", &array)
	fmt.Println()

	funcPassingCopyOfArray(array)
	fmt.Printf("array: %v\n", array)
	fmt.Println()

	funcPassingCopyOfArrayAddress(&array)
	fmt.Printf("array: %v\n", array)
	fmt.Println()

	//assignment
	array = [...]int{9, 9, 9}
	fmt.Printf("array: %v\n", array)

}

func funcPassingCopyOfArray(replica [3]int) {
	for index := range replica {
		replica[index] = 42
	}
	fmt.Printf("inside funcPassingCopyOfArray, &replica: %p\n", &replica)
	fmt.Printf("inside funcPassingCopyOfArray, replica: %v\n", replica)
}

func funcPassingCopyOfArrayAddress(replicaOfAddress *[3]int) {
	for index := range *replicaOfAddress {
		(*replicaOfAddress)[index] = 42
	}
	fmt.Printf("inside funcPassingCopyOfArrayAddress, replicaOfAddress: %p\n", replicaOfAddress)
	fmt.Printf("inside funcPassingCopyOfArrayAddress, array: %v\n", *replicaOfAddress)
}
