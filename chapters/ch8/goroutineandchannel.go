package ch8

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func UnbufferedChannel() {
	ch := make(chan int)
	go func() {
		fmt.Println("pusher: trying to push elem 42 to ch")
		ch <- 42 //block until there is a puller
		fmt.Println("pusher: elem 42 pushed to ch")

		fmt.Println("pusher: trying to push elem 84 to ch")
		ch <- 84 //block until there is a puller
		fmt.Println("pusher: elem 84 pushed to ch")
	}()
	fmt.Println("puller: trying to pull elem from ch")
	x := <-ch
	fmt.Printf("puller: elem %d pulled from ch\n", x)
	time.Sleep(time.Second * 5)
}

func UnbufferedChannelWaitSignal() { //实现goroutine之间的等停(完成condition variable功能只一对一)
	done := make(chan struct{})
	go func() {
		fmt.Println("doing long-time job")
		time.Sleep(time.Second * 5)
		done <- struct{}{}
	}()
	fmt.Println("waiting for signal to continue")
	<-done
	fmt.Println("got signal, continue. ")
}

func UnbufferedChannelWaitBroadcast() { //实现goroutine之间的等停(完成condition variable功能之一对多)

	//四个waiting goroutine等待waited goroutine, main goroutine等待四个waiting goroutine
	waitedGoroutineDone := make(chan struct{})
	var wg sync.WaitGroup
	waitingGoroutinesDone := make(chan struct{})

	//waited goroutine
	go func() {
		fmt.Println("waited goroutine: doing long-time job")
		time.Sleep(time.Second * 5)
		close(waitedGoroutineDone) //使用close, 一对多广播
	}()

	//waiting goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("waiting goroutine: waiting")
		var got interface{} = <-waitedGoroutineDone
		fmt.Printf("waiting goroutine: got %#v\n", got) //收到的是通道元素对应的零值, (不一定是nil)
		fmt.Println("waiting goroutine: continue")
	}()
	//waiting goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("waiting goroutine: waiting")
		var got interface{} = <-waitedGoroutineDone
		fmt.Printf("waiting goroutine: got %#v\n", got)
		fmt.Println("waiting goroutine: continue")
	}()
	//waiting goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("waiting goroutine: waiting")
		var got interface{} = <-waitedGoroutineDone
		fmt.Printf("waiting goroutine: got %#v\n", got)
		fmt.Println("waiting goroutine: continue")
	}()
	//waiting goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("waiting goroutine: waiting")
		var got interface{} = <-waitedGoroutineDone
		fmt.Printf("waiting goroutine: got %#v\n", got)
		fmt.Println("waiting goroutine: continue")
	}()
	//closer goroutine of waiting goroutine
	go func() {
		fmt.Println("closer goroutine: counting")
		wg.Wait()
		fmt.Println("closer goroutine: counting finished, continue")
		//因为这里的接收者只有一个goroutine, 所以close和推入一个空元素效果是一样的.
		close(waitingGoroutinesDone)

	}()
	//join point in main goroutine
	fmt.Println("main: waiting")
	<-waitingGoroutinesDone
	fmt.Println("main: continue")
}

func TryClosedChannel() {
	ch1 := make(chan int)
	go func() {
		ch1 <- 42
		ch1 <- 63
		close(ch1)
		//close(ch1) //panic: close of closed channel
		//ch1 <- 84 //panic: send on closed channel
	}()
	fmt.Println("main: waiting")
	var x int
	var ok bool
	x, ok = <-ch1
	fmt.Printf("main: x: %v ok: %t\n", x, ok)
	x = <-ch1
	fmt.Printf("main: x: %v ok: %t\n", x, ok)
	x, ok = <-ch1
	fmt.Printf("main: x: %v ok: %t\n", x, ok)
	x = <-ch1
	fmt.Printf("main: x: %v ok: %t\n", x, ok)
	fmt.Println()

	ch2 := make(chan int)
	go func() {
		ch2 <- 42
		ch2 <- 0
		ch2 <- 63
		ch2 <- 0
		close(ch2)
	}()
	for elem := range ch2 {
		fmt.Printf("main: elem: %v\n", elem)
	}
}

func incrementer(out chan<- int, in <-chan int) {
	for intElem := range in {
		out <- intElem + 1
	}
}
func squarer(out chan<- int, in <-chan int) {
	for intElem := range in {
		out <- intElem * intElem
	}
}
func printer(in <-chan int) {
	for intElem := range in {
		fmt.Println("printer:", intElem)
	}
}

func UnidirectionalChannel() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	////向通道内发送元素的时候一定要提前准备好接受者, 不然就死锁了
	////fatal error: all goroutines are asleep - deadlock!
	//for i := 0; i < 4; i++ {
	//	ch1 <- i + 1
	//}
	go incrementer(ch2, ch1)
	go squarer(ch3, ch2)
	go printer(ch3)
	for i := 0; i < 4; i++ {
		ch1 <- i
	}

	time.Sleep(time.Second * 5)
}
func BufferedChannel() {
	ch := make(chan int, 2)
	fmt.Printf("size: %d\n", cap(ch))
	fmt.Printf("elem count now: %d\n", len(ch))
	go func() {
		fmt.Println("pusher: trying to push elem 42 to ch")
		ch <- 42 //non-block
		fmt.Println("pusher: elem 42 pushed to ch")
		fmt.Printf("elem count now: %d\n", len(ch))

		fmt.Println("pusher: trying to push elem 84 to ch")
		ch <- 84 //non-block
		fmt.Println("pusher: elem 84 pushed to ch")
		fmt.Printf("elem count now: %d\n", len(ch))

		fmt.Println("pusher: trying to push elem 96 to ch")
		ch <- 96 //blocked util any previous elems pop out
		fmt.Println("pusher: elem 96 pushed to ch")
		fmt.Printf("elem count now: %d\n", len(ch))
	}()
	time.Sleep(time.Second * 5)
	fmt.Println()

	fmt.Println("puller: trying to pull elem from ch")
	x := <-ch
	fmt.Printf("puller: elem %d pulled from ch\n", x)
	time.Sleep(time.Second * 5)
	fmt.Println()
}

func workerNeedingToken(tokenChan chan struct{}) {
	//获得令牌(在buffered channel中占一个位置)
	tokenChan <- struct{}{}
	defer func() {
		//释放令牌(释放在buffered channel中占的位置)
		<-tokenChan
	}()

	//do jobs
	fmt.Println("doing jobs")
	time.Sleep(time.Second * 1)
}

func SemaphoreUsingChannelElemAsToken() { //多元信号量 (01信号量实现互斥锁, 多元信号量控制并发数量)
	//需求: 活跃的worker goroutine限制在5以内

	tokens := make(chan struct{}, 5)

	for i := 0; i < 30; i++ {
		go workerNeedingToken(tokens)
	}
	time.Sleep(time.Minute * 10)
}

/*
select{
	case <-ch1: //可能会阻塞
		//...
	case x := <-ch2: //可能会阻塞
		//...
	case ch3 <- y: //可能会阻塞
		//...
	default: //如果存在这个分支, 那么这个分支永远不阻塞, 但是当前分支优先级低, 所有case均阻塞时, 这个default后面的内容才会执行
		//...
}
*/
var cancelSignal = make(chan struct{})

//另一个goroutine close cancelSignal通道后, 在当前goroutine内调用cancelled()就可以感知到外部传进来的消息了
//注意一定要使用close, 而不是向通道推入一个元素,
//因为多个goroutine都需要使用cancelled()时, 无法保证元素推入数量和消费数量一致从而导致内存泄漏.
func cancelled() bool {
	select {
	case <-cancelSignal:
		return true
	default:
		return false
	}
}

func Select() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ { //收发交替
		fmt.Printf("i: %d\n", i)
		select {
		case x := <-ch:
			fmt.Printf("receive: %d\n", x)
		case ch <- i:
			fmt.Printf("send: %d\n", i)
		}
	}
}

func LaunchRocket1() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) //从stdin中读取一个字节放到cap为1的一个字节slice中
		close(abort)
	}()

	var timerChan <-chan time.Time = time.After(10 * time.Second) //从此句开始计时
	select {
	case <-timerChan:
		fmt.Printf("lanch!\n")
		return
	case <-abort:
		fmt.Printf("lanch abort.\n")
		return
	}
}

func LaunchRocket2() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) //从stdin中读取一个字节放到cap为1的一个字节slice中
		close(abort)
	}()

	var ticker *time.Ticker = time.NewTicker(1 * time.Second)
	var tickerChan <-chan time.Time = ticker.C
	fmt.Printf("倒计时:\n")
	for i := 10; i > 0; i-- {
		select {
		case <-tickerChan:
			fmt.Printf("%d\n", i)
		case <-abort:
			fmt.Printf("launch abort.\n")
			return
		}
	}
	ticker.Stop()
	fmt.Printf("lanch!\n")
}
