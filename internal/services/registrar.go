package services

import (
	"context"
	"database/sql"
)

const layoutsSchema = `
	create table if not exists layouts (
		id text primary key
	);
`

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

func (rr registrarRepo) SaveAndIdentifyLayout(mode string) (string, error) {
	return "l1", nil
}

func (rr registrarRepo) SaveAndIdentifyPerformer(name string) (string, error) {
	return "p1", nil
}

func newRegistrarRepo(ctx context.Context, db *sql.DB, dm *dbMigrate) (*registrarRepo, error) {
	err := dm.byScheme(ctx, performersSchema, "performers")
	if err != nil {
		return nil, err
	}
	err = dm.byScheme(ctx, layoutsSchema, "layouts")
	if err != nil {
		return nil, err
	}
	return &registrarRepo{db: db}, nil
}
