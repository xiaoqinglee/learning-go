package main

import (
	"github.com/k0kubun/pp/v3"
	"github.com/xiaoqingLee/learning-go/chapters/ch3"
)

func main() {
	str := "中国A中42中国"
	for i := -3; i < 25; i++ {
		pp.Println("limit", i)
		pp.Println(ch3.SplitOnBoundary(str, i))
	}
}
