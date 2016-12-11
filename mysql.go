package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func QueryAll(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM table_a")
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("---------------------------")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
}

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(xxx.xxx.xxx.xxx:3306)/my?charset=utf8")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	QueryAll(db)
}
