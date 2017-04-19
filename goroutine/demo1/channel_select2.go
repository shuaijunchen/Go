// channel_select2.go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const n = 100

	ch := make(chan bool)

	var done sync.WaitGroup

	for i := 0; i < n; i++ {
		done.Add(1)
		go func() {
			select {
			case <-time.After(1 * time.Hour):
			case <-ch:
			}
			done.Done()
		}()
	}

	t := time.Now()

	close(ch)

	done.Wait()

	fmt.Printf("Waited %v for %d goroutines to stop\n", time.Since(t), n)
}
