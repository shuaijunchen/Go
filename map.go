// map.go
package main

import (
	"fmt"
)

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	// 变量声明
	var personDB map[string]PersonInfo
	// 创建map
	// 也可以为 map 指定初始值
	// personDB = make(map[string]PersonInfo, 100)
	personDB = make(map[string]PersonInfo)

	// 插入数据
	personDB["123456"] = PersonInfo{"123456", "Tom", "Beijing"}
	personDB["zhang"] = PersonInfo{"zhang", "Jack", "Chengdu"}

	// 删除元素
	delete(personDB, "zhang")
	// 从 map 中查找键值为 "zhang" 的信息
	person, ok := personDB["zhang"]
	// fmt.Println(ok)
	// 判断是否成功找到特定的键，不需要检测取到的是否为 nil，只需要查看第二个返回值ok，这让表意清晰很多。
	if ok {
		fmt.Println("Found person", person.Name, "with id", person.ID)
	} else {
		fmt.Println("Not Found")
	}
}
