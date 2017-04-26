// main.go
package main

import (
	"fmt"
	"snowflake"
)

func main() {
	id, _ := snowflake.NewIdWorker(1)

	for i := 0; i < 100; i++ {
		id, err := id.NextId()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(id)
		}
	}
}
