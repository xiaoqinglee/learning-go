package bank1

import "fmt"

var deposits = make(chan int)
var balances = make(chan int)

func monitor() {
	var balance int = 42 //balance被保护在当前goroutine中
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
			//do nothing
		}
	}
}

func init() { //package的初始化函数(一个package可以有多个init()初始化函数)
	go monitor()
}

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func UseMonitorGoroutine() {
	fmt.Printf("%d\n", Balance())
	Deposit(40)
	fmt.Printf("%d\n", Balance())
	Deposit(-2)
	fmt.Printf("%d\n", Balance())
}
