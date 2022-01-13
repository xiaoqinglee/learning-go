package gotcha

import "fmt"

type Type struct {
	a int
	b int
}

//f := []<type>{{...}, {...}}
//is roughly the same as:
//f := []<type>{<type>{...}, <type>{...}}

//f := []*<type>{{...}, {...}}
//is the same as...
//f := []*<type>{&<type>{...}, &<type>{...}}

func StructSliceInit() {
	foo := []Type{{1, 2}, {11, 22}}
	bar := []*Type{{1, 2}, {11, 22}}
	//gotcha.Type
	//*gotcha.Type gotcha.Type
	fmt.Printf("%T\n", foo[0])
	fmt.Printf("%T %T\n", bar[0], *bar[0])
}
