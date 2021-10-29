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
