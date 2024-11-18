package supplement

import "fmt"

type girl struct {
	name string
	age  int
}

func Copy() {
	//值的拷贝(array, struct)
	g1 := girl{
		name: "晴雯",
		age:  21,
	}
	g2 := g1
	fmt.Printf("g1 == g2: %t \n", g1 == g2)
	fmt.Printf("&g1 == &g2: %t \n", &g1 == &g2)
	arr1 := [...]int{11, 22, 33}
	arr2 := arr1
	fmt.Printf("arr1 == arr2: %t \n", arr1 == arr2)
	fmt.Printf("&arr1 == &arr2: %t \n", &arr1 == &arr2)
	fmt.Println()

	//copy slice
	slice1 := []int{11, 22, 33, 44, 55}
	slice2 := make([]int, len(slice1), cap(slice1))
	copy(slice2, slice1)
	fmt.Printf("slice1: %p, value: %#v\n", &slice1, slice1)
	fmt.Printf("slice2: %p, value: %#v\n", &slice2, slice2)
	fmt.Println()

	//copy map (笨方法, 没有其他方法)
	oldMap := map[string]int{"one": 1, "two": 2}
	newMap := make(map[string]int)
	for key, value := range oldMap {
		newMap[key] = value
	}
}

//also see:
//https://pkg.go.dev/slices#Clone
//https://pkg.go.dev/maps#Clone
