package ch5

import (
	"fmt"
	"time"
)

func WrongTimer() {
	startedAt := time.Now()
	defer fmt.Println(time.Since(startedAt))

	time.Sleep(time.Second * 2)
}

func CorrectTimer() {
	startedAt := time.Now()
	defer func() { fmt.Println(time.Since(startedAt)) }()

	time.Sleep(time.Second * 2)
}
