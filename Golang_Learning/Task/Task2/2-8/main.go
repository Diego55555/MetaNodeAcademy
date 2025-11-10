package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	count := 0
	mu := sync.Mutex{}
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("Count:", count)
}
