package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func longRunningTask(n int, c chan int) {
	// Wait random 2 seconds
	s := rand.Intn(2000)
	time.Sleep(time.Duration(s) * time.Millisecond)
	c <- n
}

func main() {
	var wg sync.WaitGroup

	// Start 10 goroutines.
	fmt.Println("Will run 10 goroutines...")

	for i := 1; i <= 10; i++ {
		wg.Add(1)

		go func(n int) {
			c := make(chan int, 1)
			go longRunningTask(n, c)

			select {
			case <-c:
				fmt.Printf("Operation %d complete.\n", n)
			case <-time.After(1 * time.Second):
				fmt.Printf("Operation %d timed out.\n", n)
			}
			wg.Done()
		}(i)
	}

	// Wait until all operations are complete.
	wg.Wait()

	fmt.Println("Finished.")
}
