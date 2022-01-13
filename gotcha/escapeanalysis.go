package gotcha

import (
	"fmt"
)

/**

ref: https://geektutu.com/post/hpg-escape-analysis.html
ref: https://pkg.go.dev/cmd/compile

编译器决定内存分配位置的方式，就称之为逃逸分析(escape analysis)。逃逸分析由编译器完成，作用于编译阶段。

-l
	Disable inlining.
-m
	Print optimization decisions. Higher values or repetition
	produce more detail.
go build -gcflags="-l -m" source_file.go

*/

//1.指针逃逸

func EscapeViaPointer() *int {
	escapes := 42
	return &escapes
}

//2.interface{} 动态类型逃逸
//空接口 interface{} 可以表示任意的类型，如果函数参数为 interface{}，编译期间很难确定其参数的具体类型，也会发生逃逸。

func println(param interface{}) {
	fmt.Println(param)
}
func EscapeViaEmptyInterface() {
	escapes := 42
	println(escapes)
}

//3.大变量,栈空间不足

func GenerateSmallSizedVariable() {
	doesntEscape := make([]int, 1)
	_ = len(doesntEscape)
}

func GenerateBigSizedVariable() {
	escapes := make([]int, 4000000)
	_ = len(escapes)
}

func GenerateUnknownSizedVariable(n int) {
	escapes := make([]int, n) // 不确定大小
	_ = len(escapes)
}

//4.闭包中保存状态的变量

func StatefulFunctionProvider() (statefulFunction func()) {
	timeBeingCalled := 0
	return func() {
		timeBeingCalled += 1
		fmt.Printf("I'am a stateful function.\nThis is the %d time being called.\n", timeBeingCalled)
	}
}

/**
PS C:\Users\xiaoqing\GoLandProjects\LearningGo> go build -gcflags="-l -m" .\gotcha\escapeanalysis.go
# command-line-arguments
gotcha\escapeanalysis.go:25:2: moved to heap: escapes
gotcha\escapeanalysis.go:32:14: leaking param: param
gotcha\escapeanalysis.go:33:13: ... argument does not escape
gotcha\escapeanalysis.go:37:9: escapes escapes to heap
gotcha\escapeanalysis.go:43:22: make([]int, 1) does not escape
gotcha\escapeanalysis.go:48:17: make([]int, 4000000) escapes to heap
gotcha\escapeanalysis.go:53:17: make([]int, n) escapes to heap
gotcha\escapeanalysis.go:60:2: moved to heap: timeBeingCalled
gotcha\escapeanalysis.go:61:9: func literal escapes to heap
gotcha\escapeanalysis.go:63:13: ... argument does not escape
gotcha\escapeanalysis.go:63:14: timeBeingCalled escapes to heap
*/
