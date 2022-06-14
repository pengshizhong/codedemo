package main

import (
	"fmt"
	"sync"
)

type T struct {
	lock sync.Mutex
}
func (t *T) Lock() {
	t.lock.Lock()
}
func (t T) Unlock() {
	t.lock.Unlock()
}

func (t T) Unlock1() {
	fmt.Println(1)
}

func main() {
	t := T{lock: sync.Mutex{}}
	t.Lock()
	t.Unlock()
	t.Lock()
}