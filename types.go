// types.go
package main

import (
	"fmt"
	"math"
)

// Go 语言内置以下这些基础类型：
// 布尔类型：bool.
// 整形：int8、byte、int16、int、uint、uintptr等.
// 浮点类型：float32、float64.
// 复数类型：complex64、complex128.
// 字符串：string.
// 字符类型：runne.
// 错误类型：error.

// 此外，Go语言也支持以下这些符合类型
// 指针(pointer)
// 数组(array)
// 切片(slice)
// 字典(map)
// 通道（chan）
// 结构体(struct)
// 接口(interface)

// 对于常规的开发来说，用 int 和 uint 就可以，没必要用int8之类明确指定长度的类型，以免导致移植困难.

func main() {
	/**
	 * 布尔类型
	 */
	var v1 bool
	v1 = true      // v1 == true
	v2 := (1 == 2) // v2 == false
	fmt.Println(v1, v2)

	// 注意：布尔类型不能接受其他类型的赋值，不支持自动或强制的类型转换
	// var b bool
	// b = 1 	// 编译错误	异常信息(cannot use 1 (type int) as type bool in assignment)
	// b = bool(1) // 编译错误
	// fmt.Println(b)

	// 正确使用方式
	var b bool
	b = (1 != 0) // b == true
	fmt.Println("Result:", b)

	/**
	 * 整形
	 */
	// int 和 uint 类型长度和平台有关
	var value2 int32
	value1 := 64 // value1 将会自动推导为 int 类型
	// value2 = value1	// 编译错误，错误信息(cannot use value1 (type int) as type int32 in assignment)
	value2 = int32(value1) // 强制类型转换，编译通过
	fmt.Println(value2)

	// 数值运算(+、-、*、/和%)
	c := 5 % 3 // 求余运算
	fmt.Println(c)

	//比较运算(>、<、==、>=、<=、!=)
	i, j := 1, 2
	if i == j {
		fmt.Println("i and j are equal.")
	}

	// 两个不同类型的整数不能直接比较，比如int8类型的数和int类型的数不能直接比较，但各种类型的
	// 整形变量都可以直接与字面常量(literal)进行比较，比如：
	var x int32
	var y int
	x, y = 1, 2
	// if x == y { // 编译错误，错误信息(invalid operation: x == y (mismatched types int32 and int))
	// 	fmt.Println("x and y are equal.")
	// }

	if x == 1 || y == 2 { // 编译通过
		fmt.Println("x and y are equal.")
	}

	// 位运算
	var moveLeft = 1 << 10      // x << y 左移
	var moveReight = 1024 >> 10 // x >> y 右移
	fmt.Println("1 << 10 : ", moveLeft, "\n1024 >> 10 : ", moveReight)
	var dor = 1 ^ 2 // x ^ y 异或
	fmt.Println(dor)
	var versus = 20 & 20 // x & y 与
	fmt.Println(versus)
	var or = 10 | 20 // x | y 或
	fmt.Println(or)
	var negate = ^-10000 // ^x 取反
	fmt.Println(negate)

	// 浮点型
	// 浮点型用于表示小数点的数据，比如1.234就是一个浮点数据。
	var fvalue1 float32
	// fvalue1 = 12
	fvalue2 := 12.0111 // 如果不加小数点，fvalue2会被推导为整数类型而不是浮点型
	// 需要注意：fvalue2的类型是被自动推导为float64
	// fvalue1 = fvalue2	//编译错误，错误信息(cannot use fvalue2 (type float64) as type float32 in assignment)
	// 需要强制类型转换
	fvalue1 = float32(fvalue2)
	fmt.Println(fvalue1, fvalue2)

	// 浮点型比较
	// 因为浮点型不是一种精确的表达式，所以像整形那样的直接比较用==来判断两个浮点数是否相等是不可行的，这可能会导致不稳定的结果
	// 替代方式
	// p 为用户自定义的比较精确度，比如0.00001
	isV := isEqual(10.001, 10.000, 0.222)
	fmt.Println(isV)

	// 复数类型
	var comValue1 complex64
	comValue1 = 3.2 + 12i
	comValue2 := 3.2 + 12i        // comValue2 是 complex128 类型
	comValue3 := complex(3.2, 12) // comValue3 结果同 comValue2
	fmt.Println(comValue1, comValue2, comValue3)

	// 字符串
	var str string                      // 声明一个字符串变量
	fmt.Println("str int value ：", str) // 打印初始值
	str = "Hello world"                 // 字符串赋值
	ch := str[0]                        // 获取字符串的第一个字符
	fmt.Printf("The length of \"%s\" is %d \n", str, len(str))
	fmt.Printf("The first character of \"%s\" is %c.\n", str, ch)

	str2 := "Hello world" // 字符串也支持声明时进行初始化的做法
	// str2[0] = "X"         // 编译错误，错误信息(cannot assign to str2[0])
	fmt.Println(str2)

	str = "Hello, 世界"
	n := len(str)
	for i := 0; i < n; i++ {
		ch := str[i]	// 依据下标取字符串中的字符，类型为byte
		fmt.Println(i,ch)
	}
	fmt.Println()

	// 使用 Unicode 字符遍历
	for i, ch := range str {
		fmt.Println(i,ch)	// ch 的类型为rune
	}

	// 字符类型
	// 在 Go 语言中支持两个字符类型，一个是 byte (实际上是uint8的别名)，
	// 代表UTF-8字符串的单个字节的值；另一个是 rune ，代表单个Unicode 字符
 }

func isEqual(f1, f2, p float64) bool {
	return math.Abs(f1-f2) < p
}
