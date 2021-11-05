package bank2

import (
	"fmt"
	"sync"
)

var (
	rwmu    sync.RWMutex
	balance int = 42
)

func deposit(amount int) { //这个函数使用时要求上下文已经获取了互斥锁, balance变量不会出现race condition
	balance += amount
}

func Balance() int {
	rwmu.RLock()
	defer rwmu.RUnlock()
	return balance
}

func Deposit(amount int) {
	if amount < 0 {
		panic("Programming Error: Use Withdraw()")
	}
	rwmu.Lock()
	defer rwmu.Unlock()
	deposit(amount)
}

func Withdraw(amount int) (ok bool) {
	rwmu.Lock()
	defer rwmu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}

func UseLocks() {
	fmt.Printf("%d\n", Balance())
	Deposit(40)
	fmt.Printf("%d\n", Balance())
	ok := Withdraw(2)
	fmt.Printf("%t\n", ok)
	fmt.Printf("%d\n", Balance())
	ok = Withdraw(90)
	fmt.Printf("%t\n", ok)
	fmt.Printf("%d\n", Balance())
}
