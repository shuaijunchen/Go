package main

import (
	"fmt"
  "math/rand"
	"time"
)

/**
 * size 随机数位数
 * r 0.纯数字 1.小写字母 2.大写字母 3.数字和大小写字母
 */
func random(size int, r int) []byte {
	random, randoms, result := r, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := r > 2 || r < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all {
			random = rand.Intn(3)
		}
		scope, base := randoms[random][0], randoms[random][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

func main() {
	random := random(6, 0)
	fmt.Println(string(random))
}
