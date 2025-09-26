package dbx

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sproutbro/pkg/config"
)

func NewPG(pg *config.PG) *sql.DB {

	db, err := sql.Open("postgres", pg.DNS)
	if err != nil {
		fmt.Println("pg.go 15", err)
		return nil
	}
	defer db.Close()
	return db
}
