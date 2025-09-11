package main

import "database/sql"

type registrarRepo struct {
	db *sql.DB
}

func (rr registrarRepo) SaveAndIdentifyPerformer(name string) (string, error) {
	return "p1", nil
}
