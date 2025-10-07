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

func newRegistrarRepo(ctx context.Context, db *sql.DB) (repo *registrarRepo, err error) {
	if db == nil {
		err = fmt.Errorf("not connection for server")
		return
	}
	tx, txBeginErr := db.Begin()
	if txBeginErr != nil {
		err = fmt.Errorf("init a migration transaction is failed: %w", txBeginErr)
		return
	}
	defer func() {
		if txErr := tx.Rollback(); txErr != nil {
			txErr = fmt.Errorf("failed for registrar repo: %v", txErr)
			if err != nil {
				err = fmt.Errorf("%v, also %v", err, txErr)
			} else {
				err = txErr
			}
		}
	}()
	existsRow := tx.QueryRowContext(ctx, "select count(*) from information_schema.tables where table_type = 'BASE TABLE' and table_name = 'performers';")
	var tableExists int
	if checkTableErr := existsRow.Scan(&tableExists); checkTableErr != nil {
		err = fmt.Errorf("check table existence is failed: %w", checkTableErr)
		return
	}
	if tableExists == 0 {
		_, migrateErr := tx.ExecContext(ctx, performersSchema)
		if migrateErr != nil {
			err = fmt.Errorf("create schema for table is failed: %w", migrateErr)
			return
		}
	}
	if txCommitErr := tx.Commit(); txCommitErr != nil {
		err = fmt.Errorf("commit a migration transaction is failed: %w", txCommitErr)
		return
	}
	repo = &registrarRepo{db: db}
	return
}
