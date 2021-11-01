package ch1

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Dup3() {
	counts := make(map[string]int)
	filename := "./input_source_file.txt"
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		//忽略Fprintf可能产生的错误
		fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
	}
	text := string(bytes)
	words := strings.Split(text, "\n")
	for _, word := range words {
		counts[word]++
	}
	for word, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, word)
		}
	}
}
