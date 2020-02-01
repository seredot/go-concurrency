package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func longRunningTask(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Printf("Operation %d complete.\n", n)

	// Wait random 5 seconds
	s := rand.Intn(5000)
	time.Sleep(time.Duration(s) * time.Millisecond)
}

func main() {
	var wg sync.WaitGroup

	// Fan-out 10 goroutines.
	fmt.Println("Will run 10 goroutines...")
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go longRunningTask(i, &wg)
	}

	// Fan-in.
	wg.Wait()
	fmt.Println("Finished.")
}
