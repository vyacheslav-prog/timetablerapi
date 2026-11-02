package services

import _ "github.com/mattn/go-sqlite3"

const (
	countTableByNameQuerySqlite3 = "select count(*) from sqlite_master where tbl_name = '$1';"
)
