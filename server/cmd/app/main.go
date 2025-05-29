package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	r := gin.Default()
	r.StaticFile("/", "App/UI/Landing/landingPage.html")

	connect := "user=postgres password=PassXD dbname=UsersDB sslmode=disable"
	db, err := sql.Open("postgres", connect)
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
