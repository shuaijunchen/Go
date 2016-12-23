// function.go
package main

import (
	"fmt"
	"mymath"
)

func main() {
	/**
	 * 函数
	 */
	// 函数构成代码执行和逻辑结构。
	// 在Go语言中，函数基本组成为：关键字 func、函数名、参数列表、返回值、函数体和返回语句
	// 函数调用
	count, err := mymath.Add(-10, 5)
	if err == nil {
		fmt.Println(count)
	} else {
		fmt.Println(err)
	}
	mymath.Myfunc(1, 2, 3, 4, 5)
}
