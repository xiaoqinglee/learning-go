package gotcha

/**
Python3
https://docs.python.org/zh-cn/3/library/hmac.html#hmac.compare_digest

hmac.compare_digest(a, b)

	返回 a == b。 此函数使用一种经专门设计的方式通过避免基于内容的短路行为来防止定时分析，使得它适合处理密码。
*/

/**
Golang
https://pkg.go.dev/crypto/subtle#ConstantTimeCompare

func ConstantTimeCompare(x, y []byte) int

	ConstantTimeCompare returns 1 if the two slices, x and y, have equal contents and 0 otherwise.
The time taken is a function of the length of the slices and is independent of the contents.
*/
