//go:build pgsql

package services

import (
	"database/sql"

	_ "github.com/lib/pq"
)

const (
	countTableByNameQuery = "select count(*) from information_schema.tables where table_type = 'BASE TABLE' and table_name = '$1';"
	sqlDriver             = "pq"
)

func openDBConnect(dsn string) (*sql.DB, error) {
	return sql.Open(sqlDriver, dsn)
}
