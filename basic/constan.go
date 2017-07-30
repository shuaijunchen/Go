package main

import "fmt"

// 所谓常量，也就是在程序编译阶段就确定下来的值，而在程序运行时无法修改该值。

// 声明常量
const Pi float64 = 3.14

// 声明常量的例子
const s = "Hello, 世界！"
const i = 10
const Max = 1024
const b = true

func main() {
	fmt.Println(Pi, s, i, Max, b)
}
