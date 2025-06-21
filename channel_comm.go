package main

import (
	"fmt"
	"sync"
)

func main() {
	// 创建无缓冲通道
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	// 生产者协程
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i // 发送数据到通道
		}
		close(ch) // 发送完成后关闭通道
	}()

	// 消费者协程
	go func() {
		defer wg.Done()
		for num := range ch { // 从通道接收数据直到通道关闭
			fmt.Println("Received:", num)
		}
	}()

	wg.Wait() // 等待两个协程完成
}
