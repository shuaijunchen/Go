// constants.go
package main

import (
	"fmt"
	// "os"
)

func main() {

	/**
	 * 常量，在Go语言中，常量是只编译期间就已知且不可变的值。
	 * 常量可以包括数值类型（包括整数，浮点型，复数类型）、布尔类型、字符串类型
	 */

	// 所谓字面常量（literal），是指程序中硬编码的常量，如：
	// -12
	// 3.14159265358	浮点型常量
	// 3.2+12i			复数类型常量
	// true 	布尔型常量
	// "foo" 	字符串常量
	const PI float64 = 3.14159265358
	const zero = 0.0
	const (
		size int = 1024
		eof      = -1
	)
	const u, v float32 = 0, 3
	const a, b, c = 3, 4, "foo"

	fmt.Println(PI, zero, size, eof, u, v, a, b, c)

	const mask = 2 << 9
	fmt.Println(mask)

	// const home := os.GetEnv("HOME")
	// fmt.Println(home)

	/**
	 * 预定义常量
	 * Go语言中定义了这些常量：true、false和iota
	 */
	// iota 比较特殊，可以被认为是一个被编译器修改的常量，在每个const关键字出现时被重置为0
	// 然后在下一个const中出现之前，每出现一次iota，其所代表的数字会自动增1.

	// 如果两个 const 的赋值语句的表达式一样的，那么可以省略后一个赋值表达式。
	const ( // iota 被重设为0
		c0 = iota // c0 == 0
		// c1 = iota // c1 == 1
		// c2 = iota // c2 == 2
		c1 // c1 == 1
		c2 // c2 == 2
	)
	fmt.Println(c0, c1, c2)

	const (
		con0 = 1 << iota // con0 == 1 (iota 在每个const开头被重设为0)
		// con1 = 1 << iota // con1 == 2
		// con2 = 1 << iota // con2 == 4
		con1 // con1 == 2
		con2 // con2 == 4
	)
	fmt.Println(con0, con1, con2)

	const (
		x         = iota * 42 // x == 0
		y float64 = iota * 42 //y == 42.0
		w         = iota * 42 // w == 84
	)
	fmt.Println(x, y, w)

	const xx = iota // xx == 0 (因为iota又被重置为0了)
	const yy = iota // yy == 0 同上
	fmt.Println(xx, yy)

	/**
	 * 枚举
	 */
	//枚举指一系列相关的常量

	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays // 这个常量没有导出
	)
	// 注意：Go语言以大写字母开头的常量在包外可见
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, numberOfDays)
}
