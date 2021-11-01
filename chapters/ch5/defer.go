package ch5

import (
	"fmt"
	"runtime"
)

func GetInvokingFunctionName() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fmt.Sprintf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
}

func returnOperator() int {
	fmt.Printf("calling %s\n", GetInvokingFunctionName())
	return 33
}

func deferFuncParam() int {
	fmt.Printf("calling %s\n", GetInvokingFunctionName())
	return 22
}
func FunctionContainingDefer() (returnedValue int) {
	returnedValue = 11
	fmt.Printf("calling %s\n", GetInvokingFunctionName())
	defer func(param int) {
		returnedValue = param
		fmt.Printf("calling %s\n", GetInvokingFunctionName())
	}(deferFuncParam())
	fmt.Printf("calling %s\n", GetInvokingFunctionName())
	return returnOperator()
}

//在遇到defer语句的时候会对defer后函数调用的参数进行计算.
//参数计算和函数调用不发生在同一时间
func Defer() {
	fmt.Printf("calling %s\n", GetInvokingFunctionName())
	returned := FunctionContainingDefer()
	fmt.Printf("calling %s\n", GetInvokingFunctionName())
	fmt.Printf("returnedValue: %d\n", returned)
}

/*
calling C:/Users/xiaoqing/GoLandProjects/LearningGo/chapters/ch5/defer.go:37 github.com/XiaoqingLee/LearningGo/chapters/ch5.Defer

calling C:/Users/xiaoqing/GoLandProjects/LearningGo/chapters/ch5/defer.go:27 github.com/XiaoqingLee/LearningGo/chapters/ch5.FunctionContainingDefer

calling C:/Users/xiaoqing/GoLandProjects/LearningGo/chapters/ch5/defer.go:22 github.com/XiaoqingLee/LearningGo/chapters/ch5.deferFuncParam

calling C:/Users/xiaoqing/GoLandProjects/LearningGo/chapters/ch5/defer.go:32 github.com/XiaoqingLee/LearningGo/chapters/ch5.FunctionContainingDefer

calling C:/Users/xiaoqing/GoLandProjects/LearningGo/chapters/ch5/defer.go:17 github.com/XiaoqingLee/LearningGo/chapters/ch5.returnOperator

calling C:/Users/xiaoqing/GoLandProjects/LearningGo/chapters/ch5/defer.go:30 github.com/XiaoqingLee/LearningGo/chapters/ch5.FunctionContainingDefer.func1

calling C:/Users/xiaoqing/GoLandProjects/LearningGo/chapters/ch5/defer.go:39 github.com/XiaoqingLee/LearningGo/chapters/ch5.Defer

returnedValue: 22

Process finished with the exit code 0
*/
