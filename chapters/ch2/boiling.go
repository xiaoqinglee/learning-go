package ch2

import "fmt"

const boilingF = 212.0

func Boiling() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g F degree or %g C degree", f, c)
}
