// bubblesort.go
package main

import (
	"fmt"
)

func main() {
	array := []int{7, 4, 9, 2, 5, 1, 0, 6}
	BubbleSort(array)
	fmt.Println(array)
}

func BubbleSort(values []int) {
	flag := true
	length := len(values)
	for i := 0; i < length-1; i++ {
		flag = true
		for j := 0; j < length-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
			}
		}
		if flag == true {
			break
		}
	}
}
