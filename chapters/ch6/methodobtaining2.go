package ch6

import (
	"fmt"
	"sync"
)

var mapWithLock = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup(key string) (value string, ok bool) {
	mapWithLock.Lock()
	v, ok := mapWithLock.mapping[key]
	mapWithLock.Unlock()
	return v, ok
}
func Put(key string, value string) (ok bool) {
	mapWithLock.Lock()
	mapWithLock.mapping[key] = value
	mapWithLock.Unlock()
	return true
}

func MethodObtaining() {
	Put("hello", "world")
	fmt.Println(Lookup("hello"))
}
