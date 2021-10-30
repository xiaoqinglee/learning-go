package ch5

import "fmt"

func statefulFunctionProvider() (statefulFunction func()) {
	timeBeingCalled := 0
	return func() {
		timeBeingCalled += 1
		fmt.Printf("I'am a stateful function.\nThis is the %d time being called.\n", timeBeingCalled)
	}
}

func Closure() {
	functionInstance := statefulFunctionProvider()
	for i := 0; i < 5; i++ {
		functionInstance()
	}
	//保存anotherFunctionInstance的状态的timeBeingCalled变量和保存functionInstance的状态的timeBeingCalled变量不是一个实例
	anotherFunctionInstance := statefulFunctionProvider()
	for i := 0; i < 5; i++ {
		anotherFunctionInstance()
	}
}
