// main.go
package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Gender       int    `json:"gender"`
	City         string `json:"city"`
	Country      string `json:"country"`
	Country_Code int    `json:"countryCode"`
}

type Message struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   []User `json:"data"`
}

func foo(w http.ResponseWriter, r *http.Request) {
	user1 := User{"张三", 20, 1, "成都", "CHN", 156}
	user2 := User{"李四", 21, 1, "北京", "CHN", 156}
	user3 := User{"王五", 22, 1, "上海", "CHN", 156}
	user4 := User{"赵六", 22, 1, "深圳", "CHN", 156}
	user5 := User{"Tom", 21, 1, "New York", "USA", 840}
	user6 := User{"Tony", 20, 1, "London", "GBR", 826}
	message := Message{200, "OK", []User{user1, user2, user3, user4, user5, user6}}

	js, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("server", "A Go Web Server")
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func hello(w http.ResponseWriter, r *http.Request){
	// w.Header().Set("server", "A Go Web Server")
	w.Write([]byte("Hello, 世界！"))
}


func main() {
	go http.HandleFunc("/message", foo)
	go http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
