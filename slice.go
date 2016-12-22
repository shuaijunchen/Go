// array-slice.go
package main

import (
	"fmt"
)

func main() {
	/**
	 * 数组切片
	 */

	// 前面我们提到数组的特点：数组的长度在定义后无法再次修改；数组是值类型，每次转递都会产生一个副本

	// 数组切片的数组可以抽象为以下3个变量
	// 1.一个指向原生数组的指针；
	// 2.数组切片中的元素个数；
	// 3.数组切片已分配的存储空间。

	// 穿件数组切片
	// 创建数组切片的方法主要有两种——基于数组和直接创建。

	/**
	 * 基于数组
	 */
	// 定义一个数组
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var mySlice []int = myArray[:5]
	fmt.Println("Elements of myArray: ")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}

	fmt.Println("\nElements of mySlice: ")
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// Go 语言支持myArray[first:last]这个的方式来基于数组生成一个数组切片
	// 基于myArray的所有元素创建数组切片：
	// mySlice = myArray[:]

	// 基于myArray的前五个元素创建切片
	// mySlice = myArray[:5]

	// 基于从第5个元素开始所有元素创建数组切片
	// mySlice = myArray[5:]

	fmt.Println()

	/**
	 * 直接创建
	 */
	// 创建一个初始元素个数为5的数组切片：
	slice1 := make([]int, 5)
	fmt.Println(slice1)

	//创建一个初始元素个数为5的数组切片，元素初始值为0，并留下10个元素的存储空间：
	slice2 := make([]int, 5, 10)
	fmt.Println(slice2)

	// 直接创建并初始化包含5个数组切片
	slice3 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice3)

	// 元素遍历
	// 操作数组的所有方法都适用于数组切片，比如数组切片也可以按下标读写元素，用len()函数获取元素个数，并支持使用range关键字来快速遍历所有元素

	for i := 0; i < len(slice1); i++ {
		fmt.Println("slice[", i, "]=", slice1[i])
	}
	
	// 使用range关键字操作数组切片
	for i, v := range slice3 {
		fmt.Println("slice[", i, "]=", v)
	}
}
