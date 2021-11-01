package ch1

import (
	"bufio"
	"fmt"
	"os"
)

func Dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	//默认的分割函数是bufio.ScanLines 分割行
	input.Split(bufio.ScanWords) //分割单词
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
