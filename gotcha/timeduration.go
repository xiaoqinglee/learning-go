package gotcha

//	Why doesn’t this compile?
//
//	n := 100
//	time.Sleep(n * time.Millisecond)
//
//	invalid operation: n * time.Millisecond (mismatched types int and time.Duration)
//
//	Answer
//
//	There is no mixing of numeric types in Go. You can only multiply a time.Duration with
//
//	another time.Duration, or an untyped integer constant.
//
//	Here are three correct examples.
//
//	var n time.Duration = 100
//	time.Sleep(n * time.Millisecond)
//
//	const n = 100
//	time.Sleep(n * time.Millisecond)
//
//	time.Sleep(100 * time.Millisecond)

//https://yourbasic.org/golang/gotcha-multiply-duration-integer/
