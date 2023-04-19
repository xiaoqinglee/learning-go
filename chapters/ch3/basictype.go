package ch3

import (
	"bytes"
	"fmt"
	"math"
	"math/cmplx"
	"os"
	"strconv"
	"unicode/utf8"
)

func Integer() {
	fmt.Println("5/4 =", 5/4)
	fmt.Println("5.0/4 =", 5.0/4)
	fmt.Println("5/4.0 =", 5/4.0)
	fmt.Println("5.0/4.0 =", 5.0/4.0)

	var appleCount int32
	var orangeCount int64
	//var sum = appleCount + orangeCount //(mismatched types int32 and int64)
	var sum = int(appleCount) + int(orangeCount)
	fmt.Println("sum =", sum)

	o := 0666
	fmt.Printf("o: %d %[1]o %#[1]o\n", o)
	x := 0xdeadbeef
	fmt.Printf("o: %d %[1]x %#[1]x\n", x)
}
func Rune() {
	ascii := 'a'
	unicode := '国'
	newline := '\n'

	fmt.Printf("ascii:   %d %[1]c %[1]q\n", ascii)
	fmt.Printf("unicode: %d %[1]c %[1]q\n", unicode)
	fmt.Printf("newline: %d %[1]c %[1]q\n", newline)
}
func Float() {
	var z float64
	fmt.Println("z, -z, 1/z, -1/z, z/z:", z, -z, 1/z, -1/z, z/z)
	fmt.Println()
	fmt.Println("math.IsNaN(z/z):", math.IsNaN(z/z))
	fmt.Println("math.IsNaN(math.NaN()):", math.IsNaN(math.NaN()))
	fmt.Println("math.NaN() == z/z:", math.NaN() == z/z)
	fmt.Println()
	fmt.Println("math.IsInf(1/z, +1):", math.IsInf(1/z, +1))
	fmt.Println("math.IsInf(math.Inf(+1), +1)):", math.IsInf(math.Inf(+1), +1))
	fmt.Println("math.Inf(+1) == 1/z:", math.Inf(+1) == 1/z)
	fmt.Println()
	fmt.Println("math.IsInf(-1/z, -1):", math.IsInf(-1/z, -1))
	fmt.Println("math.IsInf(math.Inf(-1), -1)):", math.IsInf(math.Inf(-1), -1))
	fmt.Println("math.Inf(-1) == -1/z:", math.Inf(-1) == -1/z)
	fmt.Println()
}
func Complex() {
	fmt.Println("1-2i == 1+2i:", 1-2i == 1+2i)
	fmt.Println("1-2i == 1-2i:", 1-2i == 1-2i)
	//fmt.Println("(1+i)*(1-i) =", (1+i)*(1-i)) //Unresolved reference 'i'
	fmt.Println("(1+i)*(1-i) =", complex(1, 1)*complex(1, -1))
	fmt.Println("cmplx.Sqrt(-1) =", cmplx.Sqrt(-1))
}
func boolTestFuncReturningFalse() bool {
	fmt.Println("call func boolTestFuncReturningFalse()")
	return false
}
func boolTestFuncReturningTrue() bool {
	fmt.Println("call func boolTestFuncReturningTrue()")
	return true
}
func Boolean() {
	//fmt.Println(bool(1)) //Cannot convert an expression of the type 'int' to the type 'bool'
	//fmt.Println(bool(0)) //Cannot convert an expression of the type 'int' to the type 'bool'
	fmt.Println(true || boolTestFuncReturningFalse()) //短路
	fmt.Println(false && boolTestFuncReturningTrue()) //短路
}
func String() {
	var bytesBuffer bytes.Buffer
	bytesBuffer.WriteByte('A')
	bytesBuffer.WriteRune('甲')
	bytesBuffer.WriteString("42")
	fmt.Println("bytesBuffer:", bytesBuffer)
	fmt.Printf("bytesBuffer: %v\n", bytesBuffer)
	fmt.Printf("bytesBuffer: %q\n", bytesBuffer)
	fmt.Println("bytesBuffer.String():", bytesBuffer.String())
	fmt.Println("[]byte(bytesBuffer.String()):", []byte(bytesBuffer.String()))
	fmt.Println("[]rune(bytesBuffer.String()):", []rune(bytesBuffer.String()))
	fmt.Println()

	str := "A甲42"
	fmt.Println(str)
	fmt.Println()

	//按[]byte处理
	fmt.Println("len(str):", len(str))
	fmt.Println("str[:4]:", str[:4])
	for i, j := range []byte(str) {
		fmt.Println("i=>j:", i, j)
	}
	fmt.Println()

	//按[]rune处理
	fmt.Println("utf8.RuneCountInString(str):", utf8.RuneCountInString(str))
	for i, j := range []rune(str) {
		fmt.Println("i=>j:", i, j)
	}
	fmt.Println()

	//两不像, i是[]byte索引, j是[]rune值
	for i, j := range str {
		fmt.Println("i=>j:", i, j)
	}
	fmt.Println()

	integer, err := strconv.Atoi("42")
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(1 + integer)
	fmt.Println("1" + strconv.Itoa(42))
}
func Const() {
	fmt.Printf("type of math.Pi: %T\n", math.Pi)
	fmt.Printf("type of float32(math.Pi): %T\n", float32(math.Pi))
	fmt.Printf("type of float64(math.Pi): %T\n", float64(math.Pi))
	fmt.Println()

	const (
		jiu = 9
		jiujiu
		jiujiujiu
		ling = 0
		lingling
	)
	fmt.Printf("[]int{jiu, jiujiu, jiujiujiu, ling, lingling}: %v\n", []int{jiu, jiujiu, jiujiujiu, ling, lingling})
	fmt.Println()

	type Count int
	const (
		Zero Count = iota
		One
		Two
		Three
	)
	fmt.Printf("[]Count{Zero, One, Two, Three}: %v\n", []Count{Zero, One, Two, Three})
	fmt.Println()

	type Weekday int
	const (
		Mon Weekday = iota + 1
		Tue
		Wed
		Thur
	)
	fmt.Printf("[]Weekday{Mon, Tue, Wed, Thur}: %v\n", []Weekday{Mon, Tue, Wed, Thur})
	fmt.Println()

	type Month int
	const (
		_ Month = iota
		Jan
		Feb
		Mar
		Apr
	)
	fmt.Printf("[]Month{Jan, Feb, Mar, Apr}: %v\n", []Month{Jan, Feb, Mar, Apr})
	fmt.Println()

	type Exp int
	const (
		two Exp = 1 << (iota + 1)
		four
		eight
		sixteen
	)
	fmt.Printf("[]Exp{two, four, eight, sixteen}: %v\n", []Exp{two, four, eight, sixteen})
	fmt.Println()
}

func TestUnsignedIntSub() {
	//x= 1
	//y= 18446744073709551615
	var yi uint64 = 1
	var er uint64 = 2
	x := er - yi // var x uint64 = er - yi
	y := yi - er // var y uint64 = yi - er
	fmt.Println("x=", x)
	fmt.Println("y=", y)
}
