package supplement

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AtomicAdd() {
	var op int64
	var op2 int
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&op, 1)
			op2 += 1
		}()
	}
	wg.Wait()
	fmt.Println(op)
	fmt.Println(op2)
}
