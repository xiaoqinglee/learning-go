package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/k0kubun/pp/v3"
	"github.com/xiaoqingLee/learning-go/chapters/ch8"
	"sync"
)

const (
	ProducerNum = 2
	ConsumerNum = 5
)

func Consumer(ch <-chan int, consumerId int) {
	for {
		elem, ok := <-ch
		if !ok {
			return
		}
		fmt.Printf("Consumer %v consume %v\n", consumerId, elem)
	}
}
func Producer(ch chan<- int, producerId int) {
	for elem := 0; elem < 5; elem++ {
		fmt.Printf("Producer %v produce %v\n", producerId, elem)
		ch <- elem
	}
}

func Main() {
	ch := make(chan int)

	var consumerWg sync.WaitGroup
	consumerWg.Add(ConsumerNum)
	for i := 1; i <= ConsumerNum; i++ {
		go func(id int) {
			defer consumerWg.Done()
			Consumer(ch, id)
		}(i)
	}
	var producerWg sync.WaitGroup
	producerWg.Add(ProducerNum)
	for i := 1; i <= ProducerNum; i++ {
		go func(id int) {
			defer producerWg.Done()
			Producer(ch, id)
		}(i)
	}
	producerWg.Wait()
	close(ch)
	consumerWg.Wait()
}

//func main() {
//	Main()
//}

func main() {
	spew.Dump(42)
	pp.Println(42)
	ch8.SwitchOnAndOff()
}
