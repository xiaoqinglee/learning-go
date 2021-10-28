package ch1

import (
	"fmt"
	"os"
)

func Echo2() {
	var s, sep string
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
