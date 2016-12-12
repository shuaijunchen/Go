package main

import "fmt"

func main() {
	/**
	 * 变量声明
	 */
	var v1 int
	var v2 string
	var v3 [10]int // 数组
	var v4 []int   //数组切片

	var v5 struct {
		f int
	}

	var v6 *int           //指针
	var v7 map[string]int // map, key为string类型，value为int类型
	var v8 func(a int) int

	// 变量结束不需要使用分号作为结束符。

	fmt.Println(v1, v2, v3, v4, v5.f, v6, v7, v8)

	// var关键字可以将若干个需要声明的变量放到一起
	// 如下：
	var (
		a int
		b string
	)
	fmt.Println(a, b)

	/**
	 * 变量初始化
	 */
	var var1 int = 10 //  正确的使用方式1
	var var2 = 10     // 正确的使用方式2，编译器可以自动推导出var2的类型
	var3 := 10        //正确的使用方式3，编译器可以自动推导出var3的类型
	fmt.Printf("var1 %t\n", var1)
	fmt.Printf("var2 %t\n", var2)
	fmt.Printf("var3 %t\n", var3)

	// 错误写法
	// var i int
	// i := 3
	// 会抛出异常：no new variables on left side of :=

	// 正确使用方法：
	var v10 int
	v10 = 123
	fmt.Println(v10)

	var i, j = 10, "string"
	fmt.Println(i, j)

	ai, bi := 10, "20"
	fmt.Println(ai, bi)

	// 匿名变量
	_,_,nickname := getName()
	fmt.Println(nickname)
}

func getName() (fistname,lastname,nickname string) {
	fistname = "zhang"
	lastname = "san"
	nickname = "zhangsan"
	return
}