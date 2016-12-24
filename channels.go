// channels.go
package main

import (
	"fmt"
)

func main() {
	// 创建一个通道
	message := make(chan string)

	// 将消息存入message中
	go func() { message <- "ping" }()

	// 接收通道中的信息
	msg := <-message
	// 打印信息
	fmt.Println(msg)
}
