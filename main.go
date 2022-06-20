package main

import "fmt"

func main() {
	array1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Printf("array1: %#v\n", array1[len(array1):])
	fmt.Printf("array1: %#v\n", array1[len(array1):len(array1)])

	fmt.Printf("array1: %#v\n", array1[:0])
	fmt.Printf("array1: %#v\n", array1[0:0])
}
