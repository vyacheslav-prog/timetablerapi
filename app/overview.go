package main

import (
	"context"
	"database/sql"
	"fmt"

	"timetablerapi/overview"
)

type overviewRepo struct {
	db *sql.DB
}

const performerBoardsSchema = `
	create table if not exists performer_boards (
		id text primary key,
		created_at timestamp default current_timestamp
	);
`

func (r *overviewRepo) FetchPerformerBoard(id string) (result *overview.PerformerBoard, err error) {
	var rowCreatedAt, rowId string
	err = r.db.QueryRow("select created_at, id from performer_boards where id = $1;", id).Scan(&rowCreatedAt, &rowId)
	if rowId != "" {
		result = &overview.PerformerBoard{rowCreatedAt, rowId}
	}
	return
}

func newOverviewRepo(ctx context.Context, db *sql.DB) (*overviewRepo, error) {
	if db == nil {
		return nil, fmt.Errorf("not connection for server")
	}
	tx, err := db.Begin()
	if err != nil {
		return nil, fmt.Errorf("init a migration transaction is failed: %w", err)
	}
	defer tx.Rollback()
	existsRow := tx.QueryRowContext(ctx, "select count(*) from information_schema.tables where table_type = 'BASE TABLE' and table_name = 'performer_boards';")
	var tableExists int
	err = existsRow.Scan(&tableExists)
	if err != nil {
		return nil, fmt.Errorf("check table existence is failed: %w", err)
	}
	if 0 == tableExists {
		_, err = tx.ExecContext(ctx, performerBoardsSchema)
		if err != nil {
			return nil, fmt.Errorf("create schema for table is failed: %w", err)
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("commit a migration transaction is failed: %w", err)
	}
	return &overviewRepo{db: db}, nil
}
