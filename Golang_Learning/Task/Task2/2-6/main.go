package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Send Value:", i)
			channel <- i
		}
	}()

	go func() {
		for {
			select {
			case value, ok := <-channel:
				if !ok {
					fmt.Println("Channel is closed")
					return
				}

				fmt.Println("Receive Value:", value)
			}
		}
	}()

	time.Sleep(time.Second)
}
