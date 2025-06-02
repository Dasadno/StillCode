package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDb() {

	dsn := "host=localhost user=postgres password=pass dbname=pgsql port=5432 sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

}
