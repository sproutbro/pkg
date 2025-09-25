package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

func NewPG(envPath string) *sql.DB {
	if err := godotenv.Load(envPath); err != nil {
		fmt.Println(err)
		return nil
	}

	db, err := sql.Open("postgres", os.Getenv("PG_DNS"))
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	return db
}
