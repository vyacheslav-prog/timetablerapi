package services

const (
	countTableByNameQuerySqlite3 = "select count(*) from sqlite_master where tbl_name = '$1';"
)
