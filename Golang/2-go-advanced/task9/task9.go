package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count int
	mu    sync.Mutex
)

func increment() {
	mu.Lock()
	defer mu.Unlock()
	count++
}

func main() {
	// 启动10个goroutine同时增加计数
	for i := 0; i < 10; i++ {
		go func() {
			//1000次
			for j := 0; j < 1000; j++ {
				increment()
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println(count)
	fmt.Println("程序结束")
}
