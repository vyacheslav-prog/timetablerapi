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

func newOverviewRepo(ctx context.Context, db *sql.DB) (repo *overviewRepo, err error) {
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
			txErr = fmt.Errorf("failed rollack for overview repo: %v", txErr)
			if err != nil {
				err = fmt.Errorf("%v, also %v", err, txErr)
			} else {
				err = txErr
			}
		}
	}()
	existsRow := tx.QueryRowContext(ctx, "select count(*) from information_schema.tables where table_type = 'BASE TABLE' and table_name = 'performer_boards';")
	var tableExists int
	if existenceErr := existsRow.Scan(&tableExists); existenceErr != nil {
		err = fmt.Errorf("check table existence is failed: %w", existenceErr)
		return
	}
	if tableExists == 0 {
		_, createSchemaErr := tx.ExecContext(ctx, performerBoardsSchema)
		if createSchemaErr != nil {
			err = fmt.Errorf("create schema for table is failed: %w", createSchemaErr)
			return
		}
	}
	if txCommitErr := tx.Commit(); txCommitErr != nil {
		err = fmt.Errorf("commit a migration transaction is failed: %w", txCommitErr)
		return
	}
	repo = &overviewRepo{db: db}
	return
}
