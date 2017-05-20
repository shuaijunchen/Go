// redis.go
package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

const (
	host = "192.168.99.100:6379" // 主机:端口
	auth = "123456"              // auth 密码
	db   = 1                     // 数据库
)

const (
	conn_err   = "Connection to redis error"
	auth_err   = "Auth redis error"
	select_err = "Select redis db error"
	set_err    = "Set redis value error"
	get_err    = "Get redis key error"
)

// 连接redis
func connection() (redis.Conn, error) {
	c, err := redis.Dial("tcp", host) // 连接reids
	isError(err, conn_err)
	_, err = c.Do("AUTH", auth) // 验证redis
	isError(err, auth_err)
	_, err = c.Do("SELECT", db) // 选择连接库
	isError(err, select_err)
	return c, nil
}

// 保存值到redis中
func Set(key, value, ex, expiration interface{}) (interface{}, error) {
	c, err := connection()
	defer c.Close()

	time, err := Query(key)
	if time == -2 {
		// Set 值
		_, err = c.Do("SET", key, value, ex, expiration) // ex 表示设置过期时间；expiration 缓存时间，秒
		isError(err, set_err)
		return true, nil
	} else {
		return time, nil
	}
}

// 根据key获取值
func Query(key interface{}) (interface{}, error) {
	c, err := connection()
	defer c.Close()
	if err != nil {
		fmt.Println(err)
	}
	// TTL 获取值的倒计时
	val, err := redis.Int(c.Do("TTL", key))
	isError(err, get_err)
	return val, nil
}

// 判断错误信息
func isError(err error, errInfo interface{}) {
	if err != nil {
		fmt.Println(errInfo)
	}
}
