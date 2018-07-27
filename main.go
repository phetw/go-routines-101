package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	for i := 2; i < 13; i++ {
		wg.Add(1)
		go MultiplicationTable(i)
	}
	wg.Wait()
}

func MultiplicationTable(n int) {
	for i := 1; i < 13; i++ {
		result := n * i
		fmt.Printf("%d x %d = %d \n", n, i, result)
		time.Sleep(50 * time.Millisecond)
	}
	wg.Done()
}
