package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDb() {
	var err error
	dsn := "host=localhost user=postgres password=pass dbname=pgsql port=5432 sslmode=disable"
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

}
