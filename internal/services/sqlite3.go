//go:build sqlite3

package services

import _ "github.com/mattn/go-sqlite3"

const (
	countTableByNameQuery = "select count(*) from sqlite_master where tbl_name = '$1';"
)

func openDBConnect(dsn string) (*sql.DB, error) {
	return sql.Open("sqlite3", dsn)
}
