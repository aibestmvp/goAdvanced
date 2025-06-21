package main

import (
	"fmt"
	"sync"
	"time"
)

func printOdd(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		time.Sleep(1000)
		fmt.Println("Odd:", i)
	}
}

func printEven(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		time.Sleep(1000)
		fmt.Println("Even:", i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go printOdd(&wg)
	go printEven(&wg)

	wg.Wait()
	fmt.Println("Done!")
}
