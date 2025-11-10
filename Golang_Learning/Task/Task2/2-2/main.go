package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("odd number:", i*2+1)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("even number:", i*2+2)
		}
	}()

	time.Sleep(time.Second)
}
