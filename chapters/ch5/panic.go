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
		//do nothing
	case 1:
		panic(expectedPanic{})
	default:
		panic("unknown panic")
	}
}

func HandlePanic() (result int, err error) { //error类型的零值为nil //恢复产生自子调用的panic
	//选择性地从错误中恢复
	defer func() {
		switch panicValue := recover(); panicValue {
		case nil:
			fmt.Printf("未发生panic\n")
		case expectedPanic{}:
			fmt.Printf("发生了预期的panic\n")
			result, err = 0, fmt.Errorf("HandlePanic: xxx") //e.g. 元素不存在
		default:
			fmt.Printf("发生了未知的panic\n")
			panic(panicValue)
		}
	}()
	FunctionMightPanic()
	result, err = 42, nil
	return
}

func HandlePanic2() (result int, err error) { //恢复产生自自己的panic, recover()有这个能力, 只是没这个必要, 因为当前函数知道自己会产生哪些panic, 而且能改自己的代码, 不像调用一个外部的库函数
	defer func() {
		switch panicValue := recover(); panicValue {
		case nil:
			fmt.Printf("未发生panic\n")
		case expectedPanic{}:
			fmt.Printf("发生了预期的panic\n")
			result, err = 0, fmt.Errorf("HandlePanic: xxx")
		default:
			fmt.Printf("发生了未知的panic\n")
			panic(panicValue)
		}
	}()
	rand.Seed(time.Now().UnixNano())
	switch randomInt := rand.Intn(3); randomInt { //0, 1, 2
	case 0:
		//do nothing
	case 1:
		panic(expectedPanic{})
	default:
		panic("unknown panic")
	}
	result, err = 42, nil
	return
}

func TestPanic() {
	result, err := HandlePanic()
	if err != nil {
		fmt.Printf("error:Panic:%s\n", err.Error())
		return
	}
	fmt.Printf("result: %d\n", result)
}
func TestPanic2() {
	result, err := HandlePanic2()
	if err != nil {
		fmt.Printf("error:Panic:%s\n", err.Error())
		return
	}
	fmt.Printf("result: %d\n", result)
}

//https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-panic-recover/
//https://golang2.eddycjy.com/posts/appendix/02-goroutine-panic/

//1.recover()只能捕捉当前goroutine的panic, 跨goroutine的panic捕捉不了
//2.recover()只有放在defer里才能捕获panic
//3.如果某个goroutine的panic最终没有被捕获, 那么整个进程都会crash, 不管这个goroutine是不是main goroutine
