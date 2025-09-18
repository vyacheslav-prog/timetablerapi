package main

import (
	"context"
	"database/sql"
	"fmt"
)

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

func newRegistrarRepo(ctx context.Context, db *sql.DB) (*registrarRepo, error) {
	if db == nil {
		return nil, fmt.Errorf("not connection for server")
	}
	tx, err := db.Begin()
	if err != nil {
		return nil, fmt.Errorf("init a migration transaction is failed: %w", err)
	}
	defer tx.Rollback()
	existsRow := tx.QueryRowContext(ctx, "select count(*) from information_schema.tables where table_type = 'BASE TABLE' and table_name = 'performers';")
	var tableExists int
	err = existsRow.Scan(&tableExists)
	if err != nil {
		return nil, fmt.Errorf("check table existence is failed: %w", err)
	}
	if 0 == tableExists {
		_, err = tx.ExecContext(ctx, performersSchema)
		if err != nil {
			return nil, fmt.Errorf("create schema for table is failed: %w", err)
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("commit a migration transaction is failed: %w", err)
	}
	return &registrarRepo{db: db}, nil
}
