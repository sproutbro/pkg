package dbx

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sproutbro/pkg/config"
)

func NewSQLite(sqlite *config.Sqlite) *sql.DB {
	db, err := sql.Open("sqlite3", sqlite.SQLITE1)
	if err != nil {
		fmt.Println("sqlite.go 15", err)
	}
	return db
}
