package supplement

import (
	"context"
	"fmt"
	"time"
)

/*
	[1]Introduction

	In Go servers, each incoming request is handled in its own goroutine. Request handlers often start additional
	goroutines to access backends such as databases and RPC services. The set of goroutines working on a request typically
	needs access to request-specific values such as the identity of the end user, authorization tokens, and the request’s
	deadline. When a request is canceled or times out, all the goroutines working on that request should exit quickly so
	the system can reclaim any resources they are using.

	At Google, we developed a context package that makes it easy to pass request-scoped values, cancellation signals, and
	deadlines across API boundaries to all the goroutines involved in handling a request. The package is publicly available
	as context.

	[2]Context

	The core of the context package is the Context type:

	// A Context carries a deadline, cancellation signal, and request-scoped values
	// across API boundaries. Its methods are safe for simultaneous use by multiple
	// goroutines.
	type Context interface {
		// Done returns a channel that is closed when this Context is canceled
		// or times out.
		Done() <-chan struct{}

		// Err indicates why this context was canceled, after the Done channel
		// is closed.
		Err() error

		// Deadline returns the time when this Context will be canceled, if any.
		Deadline() (deadline time.Time, ok bool)

		// Value returns the value associated with key or nil if none.
		Value(key interface{}) interface{}
	}

	The Done method returns a channel that acts as a cancellation signal to functions running on behalf of the Context:
	when the channel is closed, the functions should abandon their work and return. The Err method returns an error
	indicating why the Context was canceled. The Pipelines and Cancellation article discusses the Done channel idiom in
	more detail.

	A Context does not have a Cancel method for the same reason the Done channel is receive-only: the function receiving a
	cancellation signal is usually not the one that sends the signal. In particular, when a parent operation starts
	goroutines for sub-operations, those sub-operations should not be able to cancel the parent. Instead, the WithCancel
	function (described below) provides a way to cancel a new Context value.

	A Context is safe for simultaneous use by multiple goroutines. Code can pass a single Context to any number of
	goroutines and cancel that Context to signal all of them.

	The Deadline method allows functions to determine whether they should start work at all; if too little time is left, it
	may not be worthwhile. Code may also use a deadline to set timeouts for I/O operations.

	Value allows a Context to carry request-scoped data. That data must be safe for simultaneous use by multiple goroutines.

	[3]Derived contexts

	The context package provides functions to derive new Context values from existing ones. These values form a tree: when
	a Context is canceled, all Contexts derived from it are also canceled.

	Background is the root of any Context tree; it is never canceled:

	// Background returns an empty Context. It is never canceled, has no deadline,
	// and has no values. Background is typically used in main, init, and tests,
	// and as the top-level Context for incoming requests.
	func Background() Context

	WithCancel and WithTimeout return derived Context values that can be canceled sooner than the parent Context. The
	Context associated with an incoming request is typically canceled when the request handler returns. WithCancel is also
	useful for canceling redundant requests when using multiple replicas. WithTimeout is useful for setting a deadline on
	requests to backend servers:

	// WithCancel returns a copy of parent whose Done channel is closed as soon as
	// parent.Done is closed or cancel is called.
	func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

	// A CancelFunc cancels a Context.
	type CancelFunc func()

	// WithTimeout returns a copy of parent whose Done channel is closed as soon as
	// parent.Done is closed, cancel is called, or timeout elapses. The new
	// Context's Deadline is the sooner of now+timeout and the parent's deadline, if
	// any. If the timer is still running, the cancel function releases its
	// resources.
	func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

	WithValue provides a way to associate request-scoped values with a Context:

	// WithValue returns a copy of parent whose Value method returns val for key.
	func WithValue(parent Context, key interface{}, val interface{}) Context


	划重点:

	In particular, when a parent operation starts goroutines for sub-operations, those sub-operations should not be able to
	cancel the parent.

	A Context is safe for simultaneous use by multiple goroutines. Code can pass a single Context to any number of
	goroutines and cancel that Context to signal all of them.

	The context package provides functions to derive new Context values from existing ones. These values form a tree: when
	a Context is canceled, all Contexts derived from it are also canceled.

	Background is the root of any Context tree; it is never canceled.

	A Context provides a key-value mapping, where the keys and values are both of type interface{}. Key types must support
	equality, and values must be safe for simultaneous use by multiple goroutines.

*/

func canceledGoroutine(ctx context.Context, nums chan<- int) {
	defer close(nums)
	num := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("ctx.Err(): %v\n", ctx.Err())
			deadline, ok := ctx.Deadline()
			fmt.Printf("deadline: %v, ok: %v\n", deadline, ok)
			return
		case nums <- num:
			num++
			time.Sleep(1 * time.Second) //模拟耗时操作
		}
	}
}

func TestContext() { //可以主动取消
	ctx, cancel := context.WithCancel(context.Background())
	nums := make(chan int)
	go canceledGoroutine(ctx, nums)
	for num := range nums {
		fmt.Printf("Got num: %v\n", num)
		if num >= 7 {
			break
		}
	}
	// A CancelFunc may be called by multiple goroutines simultaneously.
	// After the first call, subsequent calls to a CancelFunc do nothing.
	cancel()
	time.Sleep(1 * time.Second) //便于case <-ctx.Done():后面的打印能正常输出
}

func TestTimeoutContext() { //可以超时自动取消, 也可以在自动取消前手动取消, 本用例超时自动取消
	ctx, _ := context.WithTimeout(context.Background(), 13*time.Second)
	nums := make(chan int)
	go canceledGoroutine(ctx, nums)
	for num := range nums { //死循环一直打印
		fmt.Printf("Got num: %v\n", num)
	}
}

func TestValueContext() {
}

//	Soham Kamani 讲解:
//	https://www.sohamkamani.com/golang/context-cancellation-and-values/
//
//	golang blog pipeline:
//	https://go.dev/blog/pipelines
//
//										   ┌───────────close()─────────┐
//		 ┌───────────────────────close()───│───────────────────────────│
//		 ↓                                 ↓                           │
//	sq(done) ────────(in)────────> merge(done) ────────(out)────────> main
//	│                 ↑              │                   ↑
//	└─────close()─────┘              └──────close()──────┘
//
//	goroutine: sq, merge, main
//	chan: in, out, done (都是无缓冲的)
