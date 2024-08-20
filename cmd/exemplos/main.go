package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("worker %d got %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

func main() { // o main é uma go routine , thread interna / gou routine 1
	ch := make(chan int) // Esse chanel está v
	qtdWorkers := 8

	for i := range qtdWorkers {
		go worker(i, ch)
	}

	for i := range 10 {
		ch <- i
	}

}
