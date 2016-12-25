// encode_decode_base64.go
package main

import (
	"encoding/base64"
	"fmt"
)

// base64 加密
func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

// base64 解密
func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func main() {
	// encode
	hello := "你好，世界！ hello world"
	debyte := base64Encode([]byte(hello))
	fmt.Println(debyte)

	// decode
	enbyte, err := base64Decode(debyte)
	if err != nil {
		fmt.Println(err.Error())
	}
	if hello != string(enbyte) {
		fmt.Println("hello is not equal to enbyte")
	}

	fmt.Println(string(enbyte))
}
