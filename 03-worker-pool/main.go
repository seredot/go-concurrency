package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// Counts the words with more than 5 letters.
func worker(id int, lines chan string, wg *sync.WaitGroup, counter *uint64) {
	defer wg.Done()

	for line := range lines {
		trimmed := strings.Trim(line, ",./;'[]\\<>?:\"{}|`~!@#$%^&*()_+=-")
		words := strings.Split(trimmed, " ")
		var count uint64
		for _, w := range words {
			if len(w) > 5 {
				count++
			}
		}

		// This would be wrong:
		//*counter += count
		atomic.AddUint64(counter, count)
	}
}

func readFile(fn string, lines chan string) {
	defer close(lines)

	file, err := os.Open(fn)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Use every line 10X
		lines <- scanner.Text()
		lines <- scanner.Text()
		lines <- scanner.Text()
		lines <- scanner.Text()
		lines <- scanner.Text()
		lines <- scanner.Text()
		lines <- scanner.Text()
		lines <- scanner.Text()
		lines <- scanner.Text()
		lines <- scanner.Text()
	}
}

func main() {
	lines := make(chan string, 1000)

	go readFile("input/big.txt", lines)

	start := time.Now()

	var counter uint64
	var wg sync.WaitGroup
	fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())

	for i := 0; i < runtime.NumCPU(); i++ {
		id := i
		wg.Add(1)
		go worker(id, lines, &wg, &counter)
	}
	wg.Wait()

	elapsed := time.Now().Sub(start).Seconds() * 1000
	fmt.Printf("Total count: %d, elapsed: %f", counter, elapsed)
}
