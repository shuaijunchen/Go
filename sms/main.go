// main.go
package main

import (
	"encoding/json"
	"fmt"
	"random"
	"redis"
	"ucpaas"
)

const (
	ex         = "EX" // 设置redis过期
	expiration = 120  // 过期时间，秒
)

type PhoneInfo struct {
	PhoneNo      string `json:"phoneNo"`      // 电话号码
	Code         string `json:"code"`         // 验证码
	TemplateCode string `json:"templateCode"` // 模板Code
	expiration   string `json:"expiration"`   // 发送短信中的时间
}

func main() {
	rand := random.Random(6, 0)
	p := PhoneInfo{"13129521221", string(rand), "24103", "2"}
	obj, err := json.Marshal(p)
	_, err = ucpaas.SendSMS(p.PhoneNo, p.TemplateCode, p.Code, p.expiration)
	if err != nil {
		fmt.Println(err)
	} else {
		b, err := redis.Set(p.PhoneNo+"-"+p.TemplateCode, string(obj), ex, expiration)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(b)
	}
}
