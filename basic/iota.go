package main

import "fmt"

// 常量定义
const (
	x = iota // x == 0
	y = iota // y == 1
	z = iota // z == 2
	w        // 常亮声明省略值时，默认和之前的值字面相同
)

// iota 的只每次遇到const时重置
const v = iota

const (
	h, i, j = iota, iota, iota // iota 在同一行的值相同
)

const (
	a       = iota // a =0
	b       = "B"
	c       = iota
	d, e, f = iota, iota, iota
	g       = iota
)

func main() {
	fmt.Println(a, b, c, d, e, f, g, h, i, j, x, y, z, w, v)
}
