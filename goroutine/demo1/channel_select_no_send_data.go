// channel_select_no_send_data.go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan struct{})

	var done sync.WaitGroup
	done.Add(1)

	go func() {
		select {
		case <-time.After(1 * time.Hour):
		case <-ch:
		}
		done.Done()
	}()

	t := time.Now()
	close(ch)
	done.Wait()

	fmt.Printf("Waited %v for goroutine to stop\n", time.Since(t))
}
