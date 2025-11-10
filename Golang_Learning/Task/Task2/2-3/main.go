package main

import (
	"fmt"
	"time"
)

func main() {
	taskList := []func(){task1, task2, task1, task1, task1, task2}

	for _, tempTask := range taskList {
		go func(task func()) {
			beforeCall := time.Now()
			task()
			afterCall := time.Now()

			fmt.Println("Time spend:", afterCall.Sub(beforeCall).Seconds())
		}(tempTask)
	}

	time.Sleep(time.Second * 3)
}

func task1() {
	time.Sleep(time.Second)
}

func task2() {
	time.Sleep(time.Second * 2)
}
