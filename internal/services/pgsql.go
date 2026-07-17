//go:build pgsql

package services

import _ "github.com/lib/pq"

const (
	countTableByNameQuery = "select count(*) from information_schema.tables where table_type = 'BASE TABLE' and table_name = '$1';"
	sqlDriver             = "pq"
)
