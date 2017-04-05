// panic.go
package main

import (
	"os"
)

func main() {
	panic("a problem")

	_, err := os.Create("/tem/file")
	if err != nil {
		panic(err)
	}
}
