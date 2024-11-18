package gotcha

//	Why does the copy disappear?
//
//	var src, dst []int
//	src = []int{1, 2, 3}
//	copy(dst, src) // Copy elements to dst from src.
//	fmt.Println("dst:", dst)
//
//	dst: []
//
//	Answer
//
//	The number of elements copied by the copy function is the minimum of len(dst) and len(src). To make a full copy, you must allocate a big enough destination slice.
//
//	var src, dst []int
//	src = []int{1, 2, 3}
//	dst = make([]int, len(src))
//	n := copy(dst, src)
//	fmt.Println("dst:", dst, "(copied", n, "numbers)")
//
//	dst: [1 2 3] (copied 3 numbers)
//
//	The return value of the copy function is the number of elements copied. See Copy function for more about the built-in copy function in Go.
//	Using append
//
//	You could also use the append function to make a copy by appending to a nil slice.
//
//	var src, dst []int
//	src = []int{1, 2, 3}
//	dst = append(dst, src...)
//	fmt.Println("dst:", dst)
//
//	dst: [1 2 3]
//
//https://yourbasic.org/golang/gotcha-copy-missing/

//also see:
//https://pkg.go.dev/slices#Clone
