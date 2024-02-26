package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func makeConnectionString() string {
	var (
		username = os.Getenv("DATABASE_USER")
		password = os.Getenv("DATABASE_PASSWORD")
		hostname = os.Getenv("DATABASE_HOST")
		dbName = os.Getenv("DATABASE_NAME")
	)
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
	return connectionString
}

func getConnection() *sql.DB {
	db, err := sql.Open("mysql", makeConnectionString())
    if err != nil {
        log.Fatal(err)
    }
	return db
}