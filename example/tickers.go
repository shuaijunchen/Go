// tickers.go
package main

import (
	"fmt"
	"time"
)

func main() {
	ticket := time.NewTimer(time.Millisecond * 500)
	go func() {
		for t := range ticket.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticket.Stop()
	fmt.Println("Ticker stopped")
}
