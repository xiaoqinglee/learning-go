package gotcha

import (
	"github.com/davecgh/go-spew/spew"
	"math"
)

func TestFloat() {

	//有整数 X, X在左闭右闭区间[MinInt32, MaxInt32]内, 那么 X.0 <-> X 可以相互无损转换.
	//因为2的零次方等于1, 且 X 的绝对值较小, 所以 X 总能写成若干个 2**n (n 属于[0, +Inf)) 相加的形式

	//当整数 X 很大时候, 两个浮点数之间的步长会超过 1, 此时一个整数找不到其数值上相等的浮点数表示, 此时 float(X) != X

	var minInt32 int32 = math.MinInt32
	var maxInt32 int32 = math.MaxInt32
	//(bool) true
	//(int32) -2147483648
	//(int32) -2147483647
	//(int32) -2147483646
	//(int32) 2147483647
	//(int32) 2147483646
	//(int32) 2147483645
	spew.Dump(int32(float64(maxInt32)) == maxInt32)
	spew.Dump(int32(float64(minInt32)))
	spew.Dump(int32(float64(minInt32 + 1)))
	spew.Dump(int32(float64(minInt32 + 2)))
	spew.Dump(int32(float64(maxInt32)))
	spew.Dump(int32(float64(maxInt32 - 1)))
	spew.Dump(int32(float64(maxInt32 - 2)))

	var intGreaterThanMaxInt32LessThanMaxInt64 int64 = 8888888888888888888
	var maxInt64 int64 = math.MaxInt64
	//(int64) 8888888888888889344
	//(bool) false
	//(bool) false
	spew.Dump(int64(float64(intGreaterThanMaxInt32LessThanMaxInt64)))
	spew.Dump(int64(float64(intGreaterThanMaxInt32LessThanMaxInt64)) == intGreaterThanMaxInt32LessThanMaxInt64)
	spew.Dump(int64(float64(maxInt64)) == maxInt64)
}

//	Integers
//
//	There is no built-in abs function for integers, but it’s simple to write your own.
//
//	// Abs returns the absolute value of x.
//	func Abs(x int64) int64 {
//		if x < 0 {
//			return -x
//		}
//		return x
//	}
//
//	Warning: The smallest value of a signed integer doesn’t have a matching positive value.
//
//	math.MinInt64 is -9223372036854775808, but
//	math.MaxInt64 is 9223372036854775807.
//
//	Unfortunately, our Abs function returns a negative value in this case.
//
//	fmt.Println(Abs(math.MinInt64))
//	// Output: -9223372036854775808
//
//	(The Java and C libraries behave like this as well.)
//
//	Floats
//
//	The math.Abs function returns the absolute value of x.
//
//	func Abs(x float64) float64
//
//	Special cases:
//
//	Abs(±Inf) = +Inf
//	Abs(NaN) = NaN

//https://yourbasic.org/golang/absolute-value-int-float/
