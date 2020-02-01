package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func longRunningTask(n int, wg *sync.WaitGroup, q chan byte) {
	// Read from buffered channel. Unblocks channel if it's full.
	defer func() { <-q }()
	defer wg.Done()

	// Wait random 2 seconds
	s := rand.Intn(2000)
	time.Sleep(time.Duration(s) * time.Millisecond)

	fmt.Printf("Operation %d complete.\n", n)
}

func main() {
	var wg sync.WaitGroup
	// Initialize buffered channel with size of 10.
	q := make(chan byte, 10)

	// Start 100 goroutines.
	fmt.Println("Will run 100 goroutines...")

	for i := 1; i <= 100; i++ {
		// Write to the buffered channel. Will block if the buffer is full.
		q <- 0
		fmt.Printf("Concurrent goroutines: %d\n", len(q))

		wg.Add(1)
		go longRunningTask(i, &wg, q)
	}

	close(q)

	// Wait until all operations are complete.
	wg.Wait()

	fmt.Println("Finished.")
}
