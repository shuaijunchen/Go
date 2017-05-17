// postgres.go
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:123456@192.168.1.1/postgres?sslmode=disable")
	errV(err)
	insert(db)
	query(db)
}

func query(db *sql.DB) {
	rows, err := db.Query("select * from books")
	errV(err)
	for rows.Next() {
		var id int
		var data string

		rows.Scan(&id, &data)

		fmt.Println(id, data)
	}
}

func insert(db *sql.DB) {
	var count int
	sql := `insert into books (id,data) values(4,'{"ref":-1}') RETURNING id`
	err := db.QueryRow(sql).Scan(&count)
	errV(err)
	fmt.Println(count)
}

func errV(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
