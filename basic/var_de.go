package main

import "fmt"

func main() {
	// 变量定义
	var i int
	fmt.Println(i)

	// 定义多个变量
	var v1, v2, v3 string
	fmt.Println(v1, v2, v3)

	// 定义变量并赋值
	var str string = "张三"
	fmt.Println(str)

	// :=是简洁的定义变量，分别初始化相应的值，类型会根据初始值自动推导出想要的类型
	ii, f, s, d := 10, 20.01, "string", true
	fmt.Println(ii, f, s, d)
	fmt.Printf("Type i type: %T\n", ii) // 推导出变量类型
	fmt.Printf("Type f type: %T\n", f)
	fmt.Printf("Type s type: %T\n", s)
	fmt.Printf("Type d type: %T\n", d)

	// _下划线，是一个特殊的变量名，任何赋予它的值都会被丢弃
	_, b := 34, 35
	fmt.Println(b)
}
