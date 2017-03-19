// coding_algorithm.go
package main

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	str := "sbsbsbsbsb"
	fmt.Println(md5Str(str))
	fmt.Println(sha512Str(str))
	en := base64EncodeStr(str)
	fmt.Println(en)
	fmt.Println(base64DecodingStr(en))
}

// md5 编码
func md5Str(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// sha512 编码
func sha512Str(str string) string {
	h := sha512.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// base 编码
func base64EncodeStr(str string) string {
	return string(base64.StdEncoding.EncodeToString([]byte(str)))
}

// base 解码
func base64DecodingStr(str string) string {
	a, err := (base64.StdEncoding.DecodeString(str))
	if err != nil {
		return "error"
	}
	return string(a)
}
