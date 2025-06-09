package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 任务类型
type Task func()

// RunTasks 并发执行任务并统计时间
func RunTasks(tasks []Task) {
	var wg sync.WaitGroup

	for i, task := range tasks {
		wg.Add(1)
		go func(id int, t Task) {
			defer wg.Done()

			start := time.Now()
			t()
			duration := time.Since(start)

			fmt.Printf("任务%d 执行时间: %v\n", id+1, duration)
		}(i, task)
	}

	wg.Wait()
}

func main() {
	// 定义一组任务
	tasks := []Task{
		func() {
			time.Sleep(1 * time.Second)
			fmt.Println("任务1完成")
		},
		func() {
			time.Sleep(2 * time.Second)
			fmt.Println("任务2完成")
		},
		func() {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("任务3完成")
		},
	}

	// 执行任务
	RunTasks(tasks)
	fmt.Println("所有任务执行完毕")
}
