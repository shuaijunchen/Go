// gotos.go
package main

import (
	"fmt"
)

func main() {
	// 跳转语句
	myfunc(10)
}

func myfunc(i int) {
HERE: //标签
	fmt.Println(i)
	if i < 10 {
		// 跳转到某个标签
		goto HERE
	}
}
