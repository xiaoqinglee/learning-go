package supplement

import (
	"context"
	"fmt"
	"github.com/k0kubun/pp/v3"
	"golang.org/x/time/rate"
	"time"
)

//使用有缓冲channel实现令牌桶限速

func RateLimit() {

	type request int
	requests := make(chan request, 30)
	for i := 1; i <= 42; i++ {
		requests <- request(i)
	}

	bucketSize := 3
	tokenBucket := make(chan time.Time, bucketSize)
	for i := 1; i <= bucketSize; i++ { //开始状态为充满
		tokenBucket <- time.Now()
	}
	go func() {
		//使用time.Tick实现的令牌桶峰值时需要处理bucketSize+2个请求
		for range time.Tick(1 * time.Second) {
			tokenBucket <- time.Now()
		}
	}()

	var itr int
	for token := range tokenBucket {
		itr += 1

		fmt.Println("token:", token.Second())

		req := <-requests
		fmt.Println("request:", req, "now:", time.Now().Second())

		if itr == 5 {
			fmt.Println("======sleeping======")
			time.Sleep(9 * time.Second)
		}
		fmt.Println("tokens in bucket:", len(tokenBucket))
	}
}

//https://api7.ai/learning-center/openresty/how-to-deal-with-bursty-traffic

// A Limiter controls how frequently events are allowed to happen.
// It implements a "token bucket" of size b, initially full and refilled
// at rate r tokens per second.
// Informally, in any large enough time interval, the Limiter limits the
// rate to r tokens per second, with a maximum burst size of b events.
// As a special case, if r == Inf (the infinite rate), b is ignored.
// See https://en.wikipedia.org/wiki/Token_bucket for more about token buckets.
//
// The zero value is a valid Limiter, but it will reject all events.
// Use NewLimiter to create non-zero Limiters.
//
// Limiter has three main methods, Allow, Reserve, and Wait.
// Most callers should use Wait.
//
// Each of the three methods consumes a single token.
// They differ in their behavior when no token is available.
// If no token is available, Allow returns false.
// If no token is available, Reserve returns a reservation for a future token
// and the amount of time the caller must wait before using it.
// If no token is available, Wait blocks until one can be obtained
// or its associated context.Context is canceled.
//
// The methods AllowN, ReserveN, and WaitN consume n tokens.

// Use ReserveN method if you wish to wait and slow down in accordance with the rate limit without dropping events.
// If you need to respect a deadline or cancel the delay, use Wait instead.
// To drop or skip events exceeding rate limit, use Allow instead.

func ProductionReadyTokenBucket() {
	limiter := rate.NewLimiter(1, 3) // 60s 放入一个 token, 峰值消耗动作为 3 min 积累的 token 总量.
	pp.Println("limiter.Limit():", limiter.Limit())
	pp.Println("limiter.Burst():", limiter.Burst())
	for i := 0; i < 6; i++ {
		pp.Println("i:", i)
		pp.Println("limiter.Tokens():", limiter.Tokens())
		pp.Println("limiter.Allow():", limiter.Allow())
		pp.Println("limiter.Tokens():", limiter.Tokens())
	}
	pp.Println("===========================================")
	limiter = rate.NewLimiter(1, 3) // 60s 放入一个 token, 峰值消耗动作为 3 min 积累的 token 总量.
	pp.Println("limiter.Limit():", limiter.Limit())
	pp.Println("limiter.Burst():", limiter.Burst())
	for i := 0; i < 6; i++ {
		pp.Println("i:", i)
		pp.Println("limiter.Tokens():", limiter.Tokens())
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Minute)
		pp.Println("limiter.Wait(ctx):", limiter.Wait(ctx))
		pp.Println("limiter.Tokens():", limiter.Tokens())
	}
	context.WithTimeout(context.Background(), 1*time.Minute)
}
