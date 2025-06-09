package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 10)
	var wg sync.WaitGroup

	//生产者
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			ch <- i
			fmt.Println("生产：", i)
		}
		close(ch)
	}()

	//消费者
	wg.Add(1)
	go func() {
		defer wg.Done()
		for value := range ch {
			fmt.Println("接收：", value)
		}
	}()

	wg.Wait() // 等待所有协程完成
	fmt.Println("程序结束")
}
