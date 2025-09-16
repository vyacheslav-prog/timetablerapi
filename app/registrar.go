package main

import "database/sql"

const performersSchema = `
	create table if not exists performers (
		id text primary key,
		created_at timestamp default current_timestamp,
		name text
	);
`

type registrarRepo struct {
	db *sql.DB
}

func (rr registrarRepo) SaveAndIdentifyPerformer(name string) (string, error) {
	return "p1", nil
}
