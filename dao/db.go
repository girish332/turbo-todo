package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// DB ...
var DB *sql.DB

// DatabaseInit func to initialize connection to DB
func DatabaseInit() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbDsn := os.Getenv("dbDsn")
	DB, err = sql.Open("postgres", dbDsn)

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to DB")
}
