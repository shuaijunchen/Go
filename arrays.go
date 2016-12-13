// arrays.go
package main

import "fmt"

func main() {
	// 数值是 Go 语言编程中最常用的数据结构之一。顾名思义，数组就是指一系列同类型数据的集合。
	// 数组中包含的每个数据被称为数组元素(element),一个数组包含的元素个数被称为数组的长度。

	// 以下为一些常规的数组声明方法：
	// [32]byte						// 长度为32的数组，每个元素为一个字节
	// [2*N] struct { x, y int32}	// 复杂类型数组
	// [10000]*float64				// 指针数组
	// [3][5]int 					// 二位数组
	// [2][2][2]float64				// 等于[2]([2]([2]float64))

	// 在 Go 语言中，数组长度在定义后就不可更改，在声明时长度可以作为一个常量或一个常量
	// 表达式(常量表达式是指在编译期即可计算的结果的表达式)。数组的长度是该数组的一个内置常量，
	// 可以用 Go 语言的内置函数len()来获取。
	// 获取一个数组 arr 元素的个数写法如下：
	// arrLength := len(arr)

	// 元素访问
	arr := [5]int{20, 13, 51, 0, 89}
	arrLength := len(arr)
	// fmt.Println(arrLength)	// 获取数组的长度
	// 可以使用数组下标来访问数组中的元素
	for i := 0; i < arrLength; i++ {
		fmt.Println("Element", i, "of array is", arr[i])
	}

	fmt.Println()

	// Go 语言还提供了一个关键字 range , 用于便捷地遍历容器中的元素。
	for i, v := range arr {
		fmt.Println("Array element[", i, "]=", v)
	}
	// range 提供了两个返回值，一个元素的数组下标，第二个是元素的值

	// 值类型
	// 需要特别注意的是，Go 语言中数组是一个值类型(value type)
	// 所有的值类型变量在赋值和作为参数传递时都将产生一次复制动作。
	// 如果将数组作为函数的参数类型，则在函数调用时该参数将发生数据复制。
	// 因此，在函数体中无法修改传入的数组内容，因为函数内操作的只是所传入数组的一个副本
	// 例子：

	fmt.Println("before array:", arr)
	modify(arr)
	fmt.Println("after array:", arr)
}

func modify(array [5]int) {
	array[0] = 10000 // 试图修改数组的第一个元素
	fmt.Println("In modify(), array values:", array)
}
