package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in chan int
	// done chan bool
	wg *sync.WaitGroup
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	// c := make(chan int)
	w := worker{
		in: make(chan int),
		// done: make(chan bool),
		wg: wg,
	}

	go doWorker(id, w.in, w.wg)
	return w
}

func chanDemo() {
	var workers [10]worker
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()
}

func doWorker(id int, c chan int, wg *sync.WaitGroup) {
	for n := range c {
		fmt.Printf("worker %d received %c\n", id, n)
		wg.Done()
	}
}

func main() {
	chanDemo()
}
