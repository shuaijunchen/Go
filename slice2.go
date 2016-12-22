// slice2.go
package main

import (
	"fmt"
)

func main() {
	/**
	 * 动态增减数组
	 */
	// 可动态增减元素是数组切片比数组强大的功能。与数组相比，数组切片多了一个存储能力(capactity)的概念，
	// 既元素个数和分配的空间可以是两个不同的值。合理地设置存储能力的值，可以大幅降低数组切片内部重新分配
	// 内存和搬送内存块的频率，从而大大提高程序性能。

	// Go 语言支持内置的cap()函数和len()函数
	// cap()函数返回的是数组切片分配的空间大小
	// len()函数返回的是数组切片中当前所存储的元素个数

	slice := make([]int, 5, 10)

	fmt.Println("len(slice) :", len(slice))
	fmt.Println("cap(slice) :", cap(slice))

	// 如果 slice 已包含的5个函数后面继续新增元素，可以使用append()函数
	// 在 slice 中加入3个元素
	slice = append(slice, 1, 2, 3)
	fmt.Println(slice)

	// 函数 append() 的第二个参数其实是一个不定参数，我们可以按自己的需求添加若干个元素，
	// 甚至可以直接添加一个数组切片追加到另一个数组切片的末端
	slice2 := []int{8, 9, 10}
	// 给 slice 后面添加另一个数组切片
	slice = append(slice, slice2...)
	// slice = append(slice, 8, 9, 10) 等同 slice = append(slice, slice2...)
	fmt.Println(slice)

	// 数组切片会自动处理存储空间不足问题，如果追加的内容长度超过当前已分配的存储空间
	// 即cap()调用返回的信息，数组切片自动会分配一块足够大的内存。

	/**
	 * 基于数组切片创建数组切片
	 */
	oldSlice := []int{1, 2, 3, 4, 5}
	newSlice := oldSlice[:3]
	fmt.Println("cap(newSlice):", cap(newSlice), newSlice, len(newSlice))

	/**
	 * 内容复制
	 */
	//数组切片支持 Go 语言的另一个内置函数copy(),用于将一个数组切片复制到另一个数组切片，
	// 如果加入的两个数组切片不一样大，就会按照其中较小的那个数组切片的元素个数进行复制

	mySlice1 := []int{1, 2, 3, 4, 5}
	mySlice2 := []int{5, 4, 3}

	fmt.Println("copy before \n", mySlice1, mySlice2)
	// copy(mySlice2, mySlice1)
	// fmt.Println(mySlice1, mySlice2)
	copy(mySlice1, mySlice2)
	fmt.Println(mySlice1, mySlice2)
}
