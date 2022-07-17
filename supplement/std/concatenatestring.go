package std

import (
	"bytes"
	"fmt"
	"strings"
)

func OldWay() { //Use the bytes package. It has a Buffer type which implements io.Writer.
	var buffer bytes.Buffer

	for i := 0; i < 10; i++ {
		buffer.WriteString("a")
	}

	fmt.Println(buffer.String())
}

func NewWay() { //From Go 1.10 there is a strings.Builder type.
	var sb strings.Builder

	for i := 0; i < 10; i++ {
		sb.WriteString("a")
	}

	fmt.Println(sb.String())
}
