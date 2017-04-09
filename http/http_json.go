// http_json.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Book struct {
	Name    string   `json:"name"`
	Author  []string `json:"author"`
	Price   float64  `json:"price"`
	Type    string   `json:"type"`
	Version string   `json:"version"`
	Sell    bool     `json:"sell"`
}

type Books struct {
	Books []Book `json:"books"`
}

func java() Book {
	return Book{
		Name:    "Java",
		Author:  []string{"Sum", "Oracle"},
		Price:   19.98,
		Type:    "编程",
		Version: "v7.1",
		Sell:    true,
	}
}

func mysql() Book {
	return Book{
		Name:    "MySQL",
		Author:  []string{"Charles", "Park", "Tom", "Jack"},
		Price:   99.98,
		Type:    "编程",
		Version: "v2.1",
		Sell:    true,
	}
}

func def() Book {
	return Book{
		Name:    "Go语言编程基础",
		Author:  []string{"google"},
		Price:   9.98,
		Type:    "编程",
		Version: "v1.2",
		Sell:    true,
	}
}

func booksHandle(w http.ResponseWriter, r *http.Request) {
	var books Books
	books.Books = append(books.Books, def())
	books.Books = append(books.Books, mysql())
	books.Books = append(books.Books, java())
	obj, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	fmt.Fprintf(w, string(obj))
}

func bookHandle(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var book Book
	switch id {
	case "java":
		book = java()
	case "mysql":
		book = mysql()
	default:
		book = def()
	}
	b, _ := json.Marshal(book)
	fmt.Fprintf(w, string(b))
}

func main() {
	http.HandleFunc("/book", bookHandle)
	http.HandleFunc("/books", booksHandle)
	http.ListenAndServe(":4000", nil)
}
