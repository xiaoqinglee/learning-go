package ch5

import (
	"fmt"
	"math/rand"
	"time"
)

type expectedPanic struct{} //预期的错误

func FunctionMightPanic() {
	rand.Seed(time.Now().UnixNano())
	switch randomInt := rand.Intn(3); randomInt { //0, 1, 2
	case 0:
		return
	case 1:
		panic(expectedPanic{})
	default:
		panic("unknown panic")
	}
}

func HandlePanic() (result int, err error) { //error类型的零值为nil
	//选择性地从错误中恢复
	defer func() {
		switch panicValue := recover(); panicValue {
		case nil:
			fmt.Printf("未发生panic\n")
		case expectedPanic{}:
			fmt.Printf("发生了预期的panic\n")
			err = fmt.Errorf("HandlePanic: xxx") //e.g. 元素不存在
		default:
			fmt.Printf("发生了未知的panic\n")
			panic(panicValue)
		}
	}()
	FunctionMightPanic()
	result = 42
	return
}

func Panic() {
	result, err := HandlePanic()
	if err != nil {
		fmt.Printf("error:Panic:%s\n", err.Error())
		return
	}
	fmt.Printf("result: %d\n", result)
}
