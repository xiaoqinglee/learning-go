package supplement

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//ref
//https://www.cnblogs.com/apocelipes/p/10344011.html
//https://www.cnblogs.com/qcrao-2018/p/11007503.html

/*
	Done is provided for use in select statements:

		Stream generates values with DoSomething and sends them to out
		until DoSomething returns an error or ctx.Done is closed.

	func Stream(ctx context.Context, out chan<- Value) error {
		for {
			v, err := DoSomething(ctx)
			if err != nil {
				return err
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			case out <- v:
			}
		}
	}
*/
func TestContext() { //可以主动取消

	producerWaitingForPullingRequest := func(ctx context.Context) <-chan int {
		nums := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("ctx.Err(): %v\n", ctx.Err())
					deadline, ok := ctx.Deadline()
					fmt.Printf("deadline: %v, ok: %v\n", deadline, ok)
					return
				case nums <- n:
					n++
				}
			}
		}()
		return nums
	}

	ctx, cancel := context.WithCancel(context.Background())

	// A CancelFunc may be called by multiple goroutines simultaneously.
	// After the first call, subsequent calls to a CancelFunc do nothing.
	defer func() {
		cancel()
		time.Sleep(1 * time.Second) //测试使用, 便于case <-ctx.Done():后面的打印能正常输出
	}()

	for num := range producerWaitingForPullingRequest(ctx) {
		fmt.Printf("Got num: %v\n", num)
		if num > 7 {
			break
		}
	}
}

func TestTimeoutContext() { //可以自动取消, 也可以在自动取消前手动取消, 本用例自动取消

	childGoroutine := func(ctx context.Context, childGoroutineId int, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("ctx.Err(): %v\n", ctx.Err())
				deadline, ok := ctx.Deadline()
				fmt.Printf("deadline: %v, ok: %v\n", deadline, ok)
				return
			default:
				fmt.Printf("child goroutine %v is not canceled yet.\n", childGoroutineId)
			}
			fmt.Printf("child goroutine %v is working.\n", childGoroutineId)
			time.Sleep(1 * time.Second)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		cancel()
		time.Sleep(5 * time.Second) //测试使用, 便于case <-ctx.Done():后面的打印能正常输出
	}()

	wg := &sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go childGoroutine(ctx, i, wg)
	}
	wg.Wait()
}

func TestValueContext() {

	childGoroutine := func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		// Value(key interface{}) interface{}
		// 可以给context设置一些值，使用方法和map类似，key需要支持==比较操作，value需要是并发安全的
		traceId, ok := ctx.Value("parentGoroutineId").(string)
		if ok {
			fmt.Printf("parentGoroutineId %v.\n", traceId)
		} else {
			fmt.Printf("no value extracted from context.\n")
		}
	}

	wg := &sync.WaitGroup{}

	ctx := context.Background()
	wg.Add(1)
	go childGoroutine(ctx, wg)

	// 程序员有义务让 value 在自己编写的代码中互斥地使用或者使用并发安全的数据类型
	ctxWithValue := context.WithValue(ctx, "parentGoroutineId", "idFoo")
	wg.Add(1)
	go childGoroutine(ctxWithValue, wg)

	wg.Wait()
}
