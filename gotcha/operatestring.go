package gotcha

//	Why doesn’t this code compile?
//
//	s := "hello"
//	s[0] = 'H'
//	fmt.Println(s)
//
//	../main.go:3:7: cannot assign to s[0]
//
//	Answer
//
//	Go strings are immutable and behave like read-only byte slices (with a few extra properties).
//
//	To update the data, use a rune slice instead.
//
//	buf := []rune("hello")
//	buf[0] = 'H'
//	s := string(buf)
//	fmt.Println(s)  // "Hello"
//
//	If the string only contains ASCII characters, you could also use a byte slice.
//
//https://yourbasic.org/golang/gotcha-strings-are-immutable/
