package gotcha

//Why doesn’t this program compile?
//
//	func main() {
//		fruit := []string{
//			"apple",
//			"banana",
//			"cherry"
//		}
//		fmt.Println(fruit)
//	}
//
//	../main.go:5:11: syntax error: unexpected newline, expecting comma or }
//
//	Answer
//
//	In a multi-line slice, array or map literal, every line must end with a comma.
//
//	func main() {
//		fruit := []string{
//			"apple",
//			"banana",
//			"cherry", // comma added
//		}
//		fmt.Println(fruit) // "[apple banana cherry]"
//	}
//
//	This behavior is a consequence of the Go semicolon insertion rules.
//
//	As a result, you can add and remove lines without modifying the surrounding code.
//
//https://yourbasic.org/golang/gotcha-missing-comma-slice-array-map-literal/
