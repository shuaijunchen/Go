// json.go
package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   Person `json:"data"`
}

type Person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:age`
}

func main() {

	resp := &Response{
		Code:   200,
		Status: "OK",
		Data: Person{
			ID:   "11111",
			Name: "zhangsan",
			Age:  20,
		},
	}

	obj, _ := json.Marshal(resp)
	fmt.Println(string(obj))
}
