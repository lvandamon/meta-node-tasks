package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	count int64
	mu    sync.Mutex
)

func main() {
	// 启动10个goroutine同时增加计数
	for i := 0; i < 10; i++ {
		go func() {
			//1000次
			for j := 0; j < 1000; j++ {
				mu.Lock()
				atomic.AddInt64(&count, 1)
				mu.Unlock()
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println(count)
	fmt.Println("程序结束")
}
