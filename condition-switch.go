// condition-switch.go
package main

import (
	"fmt"
)

func main() {
	// 左花括号{必须与switch处于同一行;
	// 条件表达式不限制为常量或整数;
	// 单个case中，可以出现多个结果选项;
	// 与C语言等规则相反，Go 语言不需要用break来明确退出一个case;
	// 只有在case中明确添加 fallthrough 关键字，才会继续执行紧跟的下一个case;
	// 可以不设定 switch 之后的条件表达式，在此种情况下，整个switch结构与多个if...else...的逻辑作用同等。
	example(111)
}

func example(i int) {
	switch i {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	case 2:
		fallthrough
	case 3:
		fmt.Println("3")
	case 4, 5, 6:
		fmt.Println("4,5,6")
	default:
		fmt.Println("Default")
	}
}
