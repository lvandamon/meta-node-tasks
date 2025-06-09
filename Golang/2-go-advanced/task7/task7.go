package main

import (
	"fmt"
	"time"
)

func send(ch chan<- int) {
	for i := 0; i <= 10; i++ {
		ch <- i
		fmt.Println("发送:", i)
	}
	close(ch)
}

func receive(ch <-chan int) {
	for value := range ch {
		fmt.Println("接收到:", value)
	}
}

func main() {

	var ch = make(chan int)
	go send(ch)
	go receive(ch)
	time.Sleep(5 * time.Second)
	fmt.Println("程序结束")
}
