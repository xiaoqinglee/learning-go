package gotcha

//	matched, err := regexp.MatchString(`[0-9]*`, "12three45")
//	fmt.Println(matched) // true
//	fmt.Println(err)     // nil (regexp is valid)
//
//	The function regexp.MatchString (as well as most functions in the regexp package) does substring matching.
//
//	To check if a full string matches [0-9]*, anchor the start and the end of the regular expression:
//
//	the caret ^ matches the beginning of a text or line,
//	the dollar sign $ matches the end of a text.
//
//	matched, err := regexp.MatchString(`^[0-9]*$`, "12three45")
//	fmt.Println(matched) // false
//	fmt.Println(err)     // nil (regexp is valid)

//https://yourbasic.org/golang/gotcha-regexp-substring/
