package main

import "fmt"

const f = "%T\n"

func main() {
	i:=1234
	fmt.Printf(f,i)

	j := int32(1)
	fmt.Printf(f, j)

	b:=[5]byte{'h','e','l','l','o'}
	fmt.Printf(f,b)

	params := [4]int{2,3,4,5}
	fmt.Printf(f, params)
}
