package main

import "fmt"
import "errors"

var isActvie bool                   // 全局变量声明
var enabled, disabled = true, false // 忽略类型声明

func test() {
	var available bool // 一般声明
	valid := false
	available = true
	fmt.Println(available, valid)
}

var hello string         // 声明变量为字符串
var emptyStr string = "" // 声明一个字符串变量，初始化值为空

func str() {
	no, yes, maybe := "no", "yes", "maybe" // 简短声明，同时声明多个变量
	jap := "jap"
	fmt.Println(no, yes, maybe, jap, hello, emptyStr)
	c := []byte(jap) // 将字符串jap转换为[]byte类型
	c[0] = 'h'
	s2 := string(c) // 再转回string类型
	fmt.Println(s2)

	// 字符串拼接
	s := "hello"
	m := " world"
	a := s + m
	fmt.Printf("%s\n", a)

	// 修改字符串
	s = "c" + s[1:]
	fmt.Printf("%s\n", s)

	hw := `hello
		world`
	fmt.Printf("%s\n", hw)

	// 错误类型
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	test()
	fmt.Println(isActvie, enabled, disabled)

	var c complex64 = 5 + 5i // 复数
	fmt.Println(c)

	str()
}
