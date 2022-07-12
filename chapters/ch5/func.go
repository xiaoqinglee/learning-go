package ch5

import (
	"fmt"
	"strings"
)

func sumA(x, y int) int {
	return x + y
}
func sumB(x int, y int) int {
	return x + y
}
func sumC(x, y int) (z int) {
	z = x + y
	//如果不写return会发生 Missing the 'return' statement at the end of the function.
	//如果函数有返回值, 那么必须有return语句, 空return也行, 不写不行.
	//另外 return != return nil,
	//return nil 会执行 z = nil
	//see bareReturn() 相关测试
	return
}

func TestSumC() {
	sumC(1, 2)
}
func ignoreSecondParam(x, _ int) int {
	return x
}

//不知道有什么用
func ignoreAllParam(int, int) int {
	return 42
}

func sumVariadicFunction(vars ...int) (sum int) {
	fmt.Printf("vars type: %T\n", vars)
	for _, var_ := range vars {
		sum += var_
	}
	return
}

func Function() {

	//函数类型
	fmt.Printf("function type: %T\n", sumA)
	fmt.Printf("function type: %T\n", sumB)
	fmt.Printf("function type: %T\n", sumC)
	fmt.Printf("function type: %T\n", ignoreSecondParam)
	fmt.Printf("function type: %T\n", ignoreAllParam)
	fmt.Println()

	//函数变量
	functions := []func(int, int) int{sumA, sumB, sumC, ignoreSecondParam, ignoreAllParam}
	for _, function := range functions {
		fmt.Printf("function type: %T\n", function)
	}
	fmt.Println()

	var sumD func(int, int) int
	//函数只能和nil比较, 两个函数变量不能互相比
	fmt.Printf("sumD type: %T\n", sumD)
	fmt.Printf("sumD == nil: %t\n", sumD == nil)
	//var sumE func(int, int) int
	////Invalid operation: sumD == sumE (the operator == is not defined on func(int, int) int)
	//fmt.Printf("sumD == sumE %t\n", sumD == sumE)
	////Invalid operation: sumA == sumB (the operator == is not defined on func(x int, y int) int)
	//fmt.Printf("sumA == sumB: %t\n", sumA == sumB)
	fmt.Println()

	//匿名函数
	fmt.Printf("%v\n", strings.Map(func(r rune) rune { return r + 1 }, "ABCabc"))

	//参数变长函数
	fmt.Printf("sumVariadicFunction(1,2,3) = %v\n", sumVariadicFunction(1, 2, 3))
	fmt.Printf("sumVariadicFunction(1,2,3) = %v\n", sumVariadicFunction([]int{11, 22, 33}...))
	var noneVariadicFunction func([]int) int
	fmt.Printf("function type %T\n", noneVariadicFunction)
	fmt.Printf("function type %T\n", sumVariadicFunction)

}
