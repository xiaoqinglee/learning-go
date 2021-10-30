package ch5

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Eof() {
	in := bufio.NewReader(os.Stdin)
	for {
		rune_, size, err := in.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Fprintf(os.Stderr, "ch3.Eof: %s\n", err.Error())
			}
		}
		fmt.Printf("rune: %c size: %d\n", rune_, size)
	}
}
