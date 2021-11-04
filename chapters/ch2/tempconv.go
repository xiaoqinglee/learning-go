package ch2

import "fmt"

type Celsius float64
type Fahrenhait float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenhait {
	return Fahrenhait(c*9/5 + 32)
}
func FToC(f Fahrenhait) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
func (c Celsius) String() string {
	return fmt.Sprintf("%g °C", c)
}

func (f Fahrenhait) String() string {
	return fmt.Sprintf("%g °F", f)
}

func TempConv() {

	//两个Celsius实例之间可以 +-*/ 运算
	//Celsius实例和无类型浮点数常量之间可以 +-*/ 运算
	//Celsius实例和有类型浮点数字面量之间不可以 +-*/ 运算
	//一个Celsius实例和和Fahrenhait实例之间就不能 +-*/ 运算了, 虽然他们都是float64的别名.
	var c Celsius
	c = 42.0
	var cc Celsius
	cc = 40.0
	c += cc
	c += 2.0
	//c += float64(2.0) //Invalid operation: c += float64(2.0) (mismatched types Celsius and float64)
	fmt.Printf("c: %f\n", c)

	////Invalid operation: c += f (mismatched types Celsius and Fahrenhait)
	//var f Fahrenhait
	//f = 8.0
	//c += f

	fmt.Printf("%v : %v 或 %v\n", BoilingC, CToF(BoilingC), FToC(CToF(BoilingC)))
	fmt.Printf("%v : %v 或 %v\n", FreezingC, CToF(FreezingC), FToC(CToF(FreezingC)))
	fmt.Printf("%v : %v 或 %v\n", AbsoluteZeroC, CToF(AbsoluteZeroC), FToC(CToF(AbsoluteZeroC)))
	fmt.Println()
	fmt.Printf("%v: %v\n", BoilingC, CToF(BoilingC))
	fmt.Printf("%s: %s\n", BoilingC.String(), CToF(BoilingC).String())
	fmt.Printf("%s: %s\n", BoilingC, CToF(BoilingC))
	fmt.Printf("%f: %f\n", BoilingC, CToF(BoilingC))
	fmt.Printf("%g: %g\n", BoilingC, CToF(BoilingC))
	fmt.Printf("%e: %e\n", BoilingC, CToF(BoilingC))
}
