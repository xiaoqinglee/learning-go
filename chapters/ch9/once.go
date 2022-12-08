package ch9

import (
	"fmt"
	"sync"
	"time"
)

// Once对象可以让传入Do()的函数只被调用一次,
// 不管两次Do()调用在一个goroutine中还是在多个goroutine中
// 不管传入Do()的函数是同一个函数还是不同的函数

func AnyFunc() {
	fmt.Println("AnyFunc called")
}

func Once1() {
	o := sync.Once{}
	o.Do(AnyFunc)
	o.Do(AnyFunc)
}

func Once2() {
	o := sync.Once{}
	for i := 0; i < 10; i++ {
		go func(idx int) {
			o.Do(func() {
				fmt.Printf("func {%d} called.\n", idx)
			})
		}(i)
	}
	time.Sleep(time.Second)
}
