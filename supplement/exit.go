package supplement

import (
	"fmt"
	"os"
)

func Exit() {
	defer fmt.Printf("!\n") //os.Exit()退出的程序不会执行defer后面的表达式
	os.Exit(3)
}
