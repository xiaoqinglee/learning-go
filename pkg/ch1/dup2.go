package ch1

import (
	"bufio"
	"fmt"
	"os"
)

func Dup2() {
	counts := make(map[string]int)
	filename := "./input_source_file.txt"
	file, err := os.Open(filename)
	if err != nil {
		//忽略Fprintf可能产生的错误
		fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
	}
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
	}
	//忽略input.Err()中可能的错误
	for word, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, word)
		}
	}
}
