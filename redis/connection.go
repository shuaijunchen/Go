// connection.go
package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

const (
	host = "192.168.99.100:6379"
	auth = "123456"
)

func main() {
	// 连接 redis
	c, err := redis.Dial("tcp", host)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	// 关闭连接
	defer c.Close()

	// 设置 redis Auth
	_, err = c.Do("AUTH", auth)
	errorCheck("redis auth error:", err)

	// 获取 shuaijunchen 的值
	username, err := redis.String(c.Do("GET", "shuaijunchen"))
	errorCheck("redis get failed:", err)
	fmt.Printf("Got username %v \n", username)
}

func errorCheck(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
	}
}
