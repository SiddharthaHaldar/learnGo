// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

func wait() {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}
	const numGreeters, max = 5, 10
	var wg sync.WaitGroup
	for x := 1; x <= max; x++ {
		for i := 0; i < numGreeters; i++ {
			wg.Add(1)
			go hello(&wg, i+1)
		}
		wg.Wait()
		fmt.Println()
	}
}

func mutex() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	// Increment
	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	// Decrement
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("Arithmetic complete.")
}

func writeLoop(m map[int]int, mux sync.Locker) {
	for {
		mux.Lock()
		for i := 0; i < 100; i++ {
			fmt.Println("Write Goroutine writing ", i, "  to map")
			m[i] = i
		}
		mux.Unlock()
	}
}

func readLoop(index int, m map[int]int, mux sync.Locker) {
	for {
		mux.Lock()
		for k, v := range m {
			fmt.Println("Read Goroutine", index, " ", k, "-", v)
		}
		mux.Unlock()
	}
}

func main() {
	m := map[int]int{}

	mux := &sync.Mutex{}

	waiter := &sync.WaitGroup{}

	(*waiter).Add(1)

	go writeLoop(m, mux)
	go readLoop(1, m, mux)

	// stop program from exiting, must be killed
	// block := make(chan struct{})
	// <-block
	(*waiter).Wait()
}
