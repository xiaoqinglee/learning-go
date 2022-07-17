package gotcha

//	What’s up with the append function?
//
//	a := []byte("ba")
//
//	a1 := append(a, 'd')
//	a2 := append(a, 'g')
//
//	fmt.Println(string(a1)) // bag
//	fmt.Println(string(a2)) // bag
//
//	Answer
//
//	If there is room for more elements, append reuses the underlying array. Let's take a look:
//
//	a := []byte("ba")
//	fmt.Println(len(a), cap(a)) // 2 32
//
//	This means that the slices a, a1 and a2 will refer to the same underlying array in our example.
//
//	To avoid this, we need to use two separate byte arrays.
//
//	const prefix = "ba"
//
//	a1 := append([]byte(prefix), 'd')
//	a2 := append([]byte(prefix), 'g')
//
//	fmt.Println(string(a1)) // bad
//	fmt.Println(string(a2)) // bag
//
//	The scary case: It “worked” for me
//
//	In some Go implementations []byte("ba") only allocates two bytes, and then the code seems to work: the first string is "bad" and the second one "bag".
//
//	Unfortunately the code is still wrong, even though it seems to work. The program may behave differently when you run it in another environment.

//https://yourbasic.org/golang/gotcha-append/
