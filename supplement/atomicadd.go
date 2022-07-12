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
	nGoroutines := 10000
	wg.Add(nGoroutines)
	for i := 0; i < nGoroutines; i++ {
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
