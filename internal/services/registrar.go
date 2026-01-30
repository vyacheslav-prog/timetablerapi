package services

import (
	"context"
	"database/sql"
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
	err := execPgSQLMigrationByScheme(performersSchema, "performers", ctx, db)
	if err != nil {
		return nil, err
	}
	return &registrarRepo{db: db}, nil
}
