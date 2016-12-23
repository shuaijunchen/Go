// function.go
package main

import (
	"errors"
	"fmt"
)

func main() {
	/**
	 * 函数
	 */
	// 函数构成代码执行和逻辑结构。
	// 在Go语言中，函数基本组成为：关键字 func、函数名、参数列表、返回值、函数体和返回语句
	// 函数调用
	count, err := Add(-10, 5)
	if err == nil {
		fmt.Println(count)
	} else {
		fmt.Println(err)
	}
}

/**
 * 函数定义
 */
func Add(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 { // 该函数不支持负数相加
		err = errors.New("Should be non-negative numbers!")
		return
	}
	return a + b, nil
}
