// connection.go
package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

const (
	host = "127.0.0.1:6379"
	auth = "password"
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

	// 写入值永远不过期
	// _, err = c.Do("SET", "username", "zhangsan")
	// errorCheck("redis set failed:", err)

	key, err := c.Do("EXISTS", "username")
	fmt.Println(key)

	// 删除 username
	_, err = c.Do("DEL", "username")
	errorCheck("reids delete failed", err)
	fmt.Println("Delte redis key Success!")

	// 获取 username 的值
	username, err := redis.String(c.Do("GET", "username"))
	errorCheck("redis get failed:", err)
	fmt.Printf("Got username %v \n", username)

	// 写入值 10S 后过期
	_, err = c.Do("SET", "password", "123456", "EX", "1")
	errorCheck("redis set failed:", err)

	// 睡11S
	time.Sleep(2 * time.Second)

	password, err := redis.String(c.Do("GET", "password"))
	errorCheck("redis get failed:", err)
	fmt.Printf("Got password %v \n", password)
}

func errorCheck(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
	}
}
