package gotcha

//	Why doesn’t the iteration variable x notice that a[1] has been updated?
//
//	var a [2]int
//	for _, x := range a {
//	fmt.Println("x =", x)
//	a[1] = 8
//	}
//	fmt.Println("a =", a)
//
//	x = 0
//	x = 0        <- Why isn't this 8?
//	a = [0 8]
//
//	Answer
//
//	The range expression a is evaluated once before beginning the loop and a copy of the array is used to generate the iteration values.
//
//	To avoid copying the array, iterate over a slice instead.
//
//	var a [2]int
//	for _, x := range a[:] {
//	fmt.Println("x =", x)
//	a[1] = 8
//	}
//	fmt.Println("a =", a)
//
//	x = 0
//	x = 8
//	a = [0 8]

//https://yourbasic.org/golang/gotcha-range-copy-array/
