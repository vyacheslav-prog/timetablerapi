//go:build pgsql

package services

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	countTableByNameQuery = "select count(*) from information_schema.tables where table_type = 'BASE TABLE' and table_name = $1;"
)

func openDBConnect(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("postgres connect is failed: %w", err)
	}
	return db, nil
}
