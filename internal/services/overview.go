package services

import (
	"context"
	"database/sql"

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

func (r overviewRepo) FetchPerformerBoard(ctx context.Context, id string) (result *overview.PerformerBoard, err error) {
	var rowCreatedAt, rowId, rowTitle string
	err = r.db.QueryRowContext(ctx, "select created_at, id, title from performer_boards where id = $1;", id).Scan(&rowCreatedAt, &rowId, &rowTitle)
	if rowId != "" {
		result = overview.NewPerformerBoard(rowCreatedAt, rowId, rowTitle)
	}
	return
}

func newOverviewRepo(ctx context.Context, db *sql.DB) (*overviewRepo, error) {
	err := execPgSQLMigrationByScheme(ctx, performerBoardsSchema, "performer_boards", db)
	if err != nil {
		return nil, err
	}
	return &overviewRepo{db: db}, nil
}
