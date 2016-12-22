// condition-if.go
package main

import (
	"fmt"
)

func main() {
	// 条件语句不需要使用括号将条件包含起来();
	// 无论语句体内有几条语句，花括号{}必须存在的;
	// 左花括号{必须与if或else处于同一行;
	// 在if之后，条件语句之前，可以添加变量初始化语句，使用;间隔
	// 在有返回值的函数中，不允许将"最终的" return 语句包含在 if...else...结构中，否则会编译失败

	fmt.Println(example(10))
}

func example(x int) int {
	if x == 0 {
		return 0
	} else {
		return x / 2
	}
}
