package supplement

import (
	"fmt"
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
