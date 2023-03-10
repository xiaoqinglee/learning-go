package supplement

import (
	"fmt"
	"math"
	"time"
)

func CastBetweenIntAndFloat() {
	//math.MaxInt64     9223372036854775807
	//math.MaxInt32     2147483647
	//time.Now().Unix() 1678440318
	//below math.MaxInt32 is ok to cast between float64 and int
	//not equal at:     16777217

	fmt.Println("math.MaxInt64", math.MaxInt64)
	fmt.Println("math.MaxInt32", math.MaxInt32)
	fmt.Println("time.Now().Unix()", time.Now().Unix())
	var f float64 = 1.0
	var i int64 = 1
	for {
		if i != int64(f) {
			fmt.Println("not equal at:", i)
			break
		}
		if i >= math.MaxInt32 {
			fmt.Println("below math.MaxInt32 is ok to cast between float64 and int")
			break
		}
		f += 1.0
		i += 1
	}

	var f32 float32 = 1.0
	i = 1
	for {
		if i != int64(f32) {
			fmt.Println("not equal at:", i)
			break
		}
		if i >= math.MaxInt32 {
			fmt.Println("below math.MaxInt32 is ok to cast between float32 and int")
			break
		}
		f32 += 1.0
		i += 1
	}
}
