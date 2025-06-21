package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	Name string
	Func func()
}

type TaskResult struct {
	Name     string
	Duration time.Duration
}

func runTask(task Task, wg *sync.WaitGroup, results chan<- TaskResult) {
	defer wg.Done()

	start := time.Now()
	task.Func()
	duration := time.Since(start)

	results <- TaskResult{
		Name:     task.Name,
		Duration: duration,
	}
}

func Schedule(tasks []Task) map[string]time.Duration {
	var wg sync.WaitGroup
	// 创建能容纳N个结果的通道（N=tasks数量）,缓冲量为tasks数量
	results := make(chan TaskResult, len(tasks))

	for _, task := range tasks {
		wg.Add(1)
		go runTask(task, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	stats := make(map[string]time.Duration)
	for result := range results {
		stats[result.Name] = result.Duration
	}

	return stats
}

func main() {
	// 定义一个task集合
	tasks := []Task{
		{"Task1", func() { time.Sleep(1 * time.Second) }},
		{"Task2", func() { time.Sleep(2 * time.Second) }},
		{"Task3", func() { time.Sleep(500 * time.Millisecond) }},
	}

	// 传入是任务列表，传出是每个任务的运行时间
	stats := Schedule(tasks)
	for name, duration := range stats {
		fmt.Printf("%s took %v\n", name, duration)
	}
}
