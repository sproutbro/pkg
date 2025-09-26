package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/joho/godotenv"
)

func NewSQLite(envPath string) *sql.DB {
	if err := godotenv.Load(envPath); err != nil {
		fmt.Println(err)
		return nil
	}

	db, err := sql.Open("sqlite3", os.Getenv("SQLITE_PATH"))
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
