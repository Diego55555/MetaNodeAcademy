package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	count := int64(0)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&count, 1)
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("Count:", count)
}
