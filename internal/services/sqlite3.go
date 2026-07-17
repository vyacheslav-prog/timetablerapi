//go:build sqlite3

package services

import _ "github.com/mattn/go-sqlite3"

const (
	countTableByNameQuery = "select count(*) from sqlite_master where tbl_name = '$1';"
	sqlDriver             = "sqlite3"
)
