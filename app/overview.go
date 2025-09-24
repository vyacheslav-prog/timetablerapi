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
		created_at timestamp default current_timestamp,
		title text
	);
`

func (r overviewRepo) FetchPerformerBoard(id string) (result *overview.PerformerBoard, err error) {
	var rowCreatedAt, rowId, rowTitle string
	err = r.db.QueryRow("select created_at, id, title from performer_boards where id = $1;", id).Scan(&rowCreatedAt, &rowId, &rowTitle)
	if rowId != "" {
		result = overview.NewPerformerBoard(rowCreatedAt, rowId, rowTitle)
	}
	return
}

func newOverviewRepo(ctx context.Context, db *sql.DB) (*overviewRepo, error) {
	if db == nil {
		return nil, fmt.Errorf("not connection for server")
	}
	tx, txBeginErr := db.Begin()
	if txBeginErr != nil {
		return nil, fmt.Errorf("init a migration transaction is failed: %w", txBeginErr)
	}
	defer tx.Rollback()
	existsRow := tx.QueryRowContext(ctx, "select count(*) from information_schema.tables where table_type = 'BASE TABLE' and table_name = 'performer_boards';")
	var tableExists int
	if existenceErr := existsRow.Scan(&tableExists); existenceErr != nil {
		return nil, fmt.Errorf("check table existence is failed: %w", existenceErr)
	}
	if 0 == tableExists {
		_, createSchemaErr := tx.ExecContext(ctx, performerBoardsSchema)
		if createSchemaErr != nil {
			return nil, fmt.Errorf("create schema for table is failed: %w", createSchemaErr)
		}
	}
	if txCommitErr := tx.Commit(); txCommitErr != nil {
		return nil, fmt.Errorf("commit a migration transaction is failed: %w", txCommitErr)
	}
	return &overviewRepo{db: db}, nil
}
