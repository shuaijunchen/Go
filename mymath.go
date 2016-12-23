package mymath

import (
	"errors"
	"fmt"
)

/**
 * 函数定义
 */
func Add(a, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("Should be non-negative numbers!")
		return
	}
	return a + b, nil // 支持多重返回
}

// 不定参数
func Myfunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}
