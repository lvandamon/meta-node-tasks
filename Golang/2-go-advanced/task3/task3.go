package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

func odd() {
	mu.Lock()
	defer mu.Unlock()
	for i := 1; i <= 10; i += 2 {
		fmt.Printf("奇数: %d\n", i)
	}
}

func even() {
	mu.Lock()
	defer mu.Unlock()
	for i := 2; i <= 10; i += 2 {
		fmt.Printf("偶数: %d\n", i)
	}
}

func main() {
	go odd()
	go even()
	time.Sleep(time.Second)
}
