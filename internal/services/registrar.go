package services

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
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

const taskSchema = `
	create table if not exists tasks (
		id text primary key,
		created_at timestamp default current_timestamp,
		name text,
		from text,
		to text
	);
`

type registrarRepo struct {
	db *sql.DB
}

func (rr registrarRepo) SaveAndIdentifyLayout(ctx context.Context, mode string) (string, error) {
	return "l1", nil
}

func (rr registrarRepo) SaveAndIdentifyPerformer(ctx context.Context, name string) (string, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("failed for generating a performer identity: %w", err)
	}
	id := uid.String()
	_, err = rr.db.ExecContext(ctx, "insert into performers (id, name) values (?)", id, name)
	if err != nil {
		return "", fmt.Errorf("failed for storing a performer: %w", err)
	}
	return id, nil
}

func (rr registrarRepo) SaveAndIdentifyTask(ctx context.Context, name, from, to string) (string, error) {
	return "t1", nil
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
	err = dm.byScheme(ctx, taskSchema, "tasks")
	if err != nil {
		return nil, err
	}
	return &registrarRepo{db: db}, nil
}
