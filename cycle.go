// cycle.go
package main

import (
	"fmt"
)

func main() {
	/**
	 * 循环语句
	 */
	// 与多种语言不同的是，Go语言中的循环只支持for关键字，而不支持while和do-while结构

	// 例子
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// Go 语言for后面的条件条件表达式不需要用圆括号()包含起来

	sum = 0
	for {
		sum++
		if sum > 100 {
			break
		}
	}
	fmt.Println(sum)

	// 排序
	a := []int{1, 2, 3, 4, 5, 6}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		// 交换值
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(a)

	// 使用循环语句时，需要注意以下几点：
	// 左花括号{必须与for处于同一行
	// Go语言中的for循环与C语言一样，都允许在循环条件中定义和初始化变量，唯一的区别是，
	// Go语言不支持以逗号为间隔的多个赋值语句，必须使用平行赋值的方式来初始化多个变量
	// Go语言的for循环同样支持continue和break来控制循环，但是他提供了一个更高级的brake，
	// 可以选中断哪一个循环，例子

JLoop:
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break JLoop // 中断 JLoop 标签处的循环
			}
			fmt.Println(i)
		}
	}
}
