package dbx

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sproutbro/pkg/config"
)

var dns = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/seoul"

func NewPG(pg *config.PG) (*sql.DB, error) {
	dns := fmt.Sprintf(dns, pg.HOST, pg.USER, pg.PASS, pg.NAME, pg.PORT)
	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, fmt.Errorf("pg.go 16 %w", err)
	}
	defer db.Close()
	return db, nil
}
