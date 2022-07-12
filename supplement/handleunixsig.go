package supplement

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func HandleUnixSig() {
	sigs := make(chan os.Signal, 1)
	done := make(chan struct{})

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) //向操作系统注册感兴趣的信号

	go func() { //绑定信号处理逻辑
		sig := <-sigs
		fmt.Printf("sig value: %#v\n", sig)
		done <- struct{}{}
	}()
	fmt.Printf("waiting for signals\n")
	<-done
	fmt.Printf("done.\n")
}
