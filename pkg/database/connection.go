package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnect() (*sql.DB, error) {
	conn, err := sql.Open("mysql",
		"root:@tcp(127.0.0.1:3306)/customers")
	if err != nil {
		fmt.Println("Failed to open database connection")
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
