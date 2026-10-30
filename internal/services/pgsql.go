package services

const (
	countTableByNameQueryPg = "select count(*) from information_schema.tables where table_type = 'BASE TABLE' and table_name = $1;"
)
