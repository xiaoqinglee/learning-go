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

	// A Context carries a deadline, a cancellation signal, and other values across
	// API boundaries.
	//
	// Context's methods may be called by multiple goroutines simultaneously.
	type Context interface {
		// Deadline returns the time when work done on behalf of this context
		// should be canceled. Deadline returns ok==false when no deadline is
		// set. Successive calls to Deadline return the same results.
		Deadline() (deadline time.Time, ok bool)

		// Done returns a channel that's closed when work done on behalf of this
		// context should be canceled. Done may return nil if this context can
		// never be canceled. Successive calls to Done return the same value.
		// The close of the Done channel may happen asynchronously,
		// after the cancel function returns.
		//
		// WithCancel arranges for Done to be closed when cancel is called;
		// WithDeadline arranges for Done to be closed when the deadline
		// expires; WithTimeout arranges for Done to be closed when the timeout
		// elapses.
		//
		// Done is provided for use in select statements:
		//
		//  // Stream generates values with DoSomething and sends them to out
		//  // until DoSomething returns an error or ctx.Done is closed.
		//  func Stream(ctx context.Context, out chan<- Value) error {
		//  	for {
		//  		v, err := DoSomething(ctx)
		//  		if err != nil {
		//  			return err
		//  		}
		//  		select {
		//  		case <-ctx.Done():
		//  			return ctx.Err()
		//  		case out <- v:
		//  		}
		//  	}
		//  }
		//
		// See https://blog.golang.org/pipelines for more examples of how to use
		// a Done channel for cancellation.
		Done() <-chan struct{}

		// If Done is not yet closed, Err returns nil.
		// If Done is closed, Err returns a non-nil error explaining why:
		// Canceled if the context was canceled
		// or DeadlineExceeded if the context's deadline passed.
		// After Err returns a non-nil error, successive calls to Err return the same error.
		Err() error

		// Value returns the value associated with this context for key, or nil
		// if no value is associated with key. Successive calls to Value with
		// the same key returns the same result.
		//
		// Use context values only for request-scoped data that transits
		// processes and API boundaries, not for passing optional parameters to
		// functions.
		//
		// A key identifies a specific value in a Context. Functions that wish
		// to store values in Context typically allocate a key in a global
		// variable then use that key as the argument to context.WithValue and
		// Context.Value. A key can be any type that supports equality;
		// packages should define keys as an unexported type to avoid
		// collisions.
		//
		// Packages that define a Context key should provide type-safe accessors
		// for the values stored using that key:
		//
		// 	// Package user defines a User type that's stored in Contexts.
		// 	package user
		//
		// 	import "context"
		//
		// 	// User is the type of value stored in the Contexts.
		// 	type User struct {...}
		//
		// 	// key is an unexported type for keys defined in this package.
		// 	// This prevents collisions with keys defined in other packages.
		// 	type key int
		//
		// 	// userKey is the key for user.User values in Contexts. It is
		// 	// unexported; clients use user.NewContext and user.FromContext
		// 	// instead of using this key directly.
		// 	var userKey key
		//
		// 	// NewContext returns a new Context that carries value u.
		// 	func NewContext(ctx context.Context, u *User) context.Context {
		// 		return context.WithValue(ctx, userKey, u)
		// 	}
		//
		// 	// FromContext returns the User value stored in ctx, if any.
		// 	func FromContext(ctx context.Context) (*User, bool) {
		// 		u, ok := ctx.Value(userKey).(*User)
		// 		return u, ok
		// 	}
		Value(key any) any
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

	// Deadline returns the time when work done on behalf of this context
	// should be canceled. Deadline returns ok==false when no deadline is
	// set. Successive calls to Deadline return the same results.
	deadline, ok := ctx.Deadline()
	fmt.Printf("deadline: %v, ok: %v\n", deadline, ok)

	defer close(nums)
	num := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("ctx.Err(): %v\n", ctx.Err())
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
	time.Sleep(5 * time.Second) //使得 case <-ctx.Done() 能够执行
}

func TestTimeoutContext() { //可以超时自动取消, 也可以在自动取消前手动取消, 本用例超时自动取消
	ctx, _ := context.WithTimeout(context.Background(), 13*time.Second)
	nums := make(chan int)
	go canceledGoroutine(ctx, nums)
	for num := range nums { //死循环一直打印
		fmt.Printf("Got num: %v\n", num)
	}
}

// WithValue returns a copy of parent in which the value associated with key is
// val.
// The provided key must be comparable and should not be of type
// string or any other built-in type to avoid collisions between
// packages using context. Users of WithValue should define their own
// types for keys.

// 为了防止和导入的标准库和第三方工具包里面的ctx的key碰撞, 不小心覆盖旧key-val,
// 所以要使用自定义的key类型而不是使用内建的string int64等类型,
// 而且这个类型应该是不可导出的.

func TestDerivedValueContext() {

	// 注意 context实例的key是any类型, 所以 "foo" 和 favContextKey("foo")是两个不同的key
	type favContextKey string

	testKeys := func(ctx context.Context) {
		for _, key := range []string{"OldKey1", "OldKey2", "OldKey3", "NewKey"} {
			if v := ctx.Value(favContextKey(key)); v != nil {
				fmt.Println("found:", key, v)
			} else {
				fmt.Println("not found:", key, v)
			}
		}
	}

	origin := context.Background()
	ctx := context.WithValue(origin, favContextKey("OldKey1"), "OldValue1")
	ctx = context.WithValue(ctx, favContextKey("OldKey2"), "OldValue2")
	ctx = context.WithValue(ctx, favContextKey("OldKey3"), "OldValue3")

	derived := context.WithValue(ctx, favContextKey("OldKey1"), "NewValue") //"覆盖"
	derived = context.WithValue(derived, favContextKey("OldKey2"), nil)     //"删除"
	derived = context.WithValue(derived, favContextKey("NewKey"), "NewVal") //"新增"

	testKeys(ctx)
	testKeys(derived)

	//context对象是immutable的, 只能拷贝.
	//在derived context的过程中可以使用WithValue方法得一个拥有旧key新value的新context实例,
	//表面上达到了覆盖(或删除)的效果, 但是实际上是再封装, 读取一个context的key的过程是从外到内剥洋葱寻找第一个匹配到的目标.

	//found: OldKey1 OldValue1
	//found: OldKey2 OldValue2
	//found: OldKey3 OldValue3
	//not found: NewKey <nil>
	//found: OldKey1 NewValue
	//not found: OldKey2 <nil>
	//found: OldKey3 OldValue3
	//found: NewKey NewVal
}

func TestContextAsParameter() {

	type favContextKey string

	testKeys := func(ctx context.Context) {
		for _, key := range []string{"oldKey", "newKey"} {
			if v := ctx.Value(favContextKey(key)); v != nil {
				fmt.Println("found:", key, v)
			} else {
				fmt.Println("not found:", key, v)
			}
		}
	}

	changeParam := func(ctx context.Context) {
		ctx = context.WithValue(ctx, favContextKey("newKey"), "newValue")
		fmt.Println("in func:")
		testKeys(ctx)
	}

	origin := context.Background()
	origin = context.WithValue(origin, favContextKey("oldKey"), "oldValue")

	fmt.Println("out func:")
	testKeys(origin)
	changeParam(origin)
	fmt.Println("out func:")
	testKeys(origin)

	//out func:
	//found: oldKey oldValue
	//not found: newKey <nil>
	//in func:
	//found: oldKey oldValue
	//found: newKey newValue
	//out func:
	//found: oldKey oldValue
	//not found: newKey <nil>
}

//	Soham Kamani 讲解:
//	https://www.sohamkamani.com/golang/context-cancellation-and-values/
//
//	golang blog:
//  https://go.dev/blog/context
//	https://go.dev/blog/pipelines
//`
//                                           ┌───────────close()─────────┐
//         ┌───────────────────────close()───│───────────────────────────│
//         ↓                                 ↓                           │
//    sq(done) ────────(in)────────> merge(done) ────────(out)────────> main
//    │                 ↑              │                   ↑
//    └─────close()─────┘              └──────close()──────┘
//
//`
//	goroutine: sq, merge, main
//	chan: in, out, done (都是无缓冲的)
