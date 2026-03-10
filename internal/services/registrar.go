package services

import (
	"context"
	"database/sql"
	"strconv"
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

func (rr registrarRepo) SaveAndIdentifyLayout(mode string) (string, error) {
	return "l1", nil
}

func (rr registrarRepo) SaveAndIdentifyPerformer(name string) (string, error) {
	res, err := rr.db.Exec("insert into performers (name) values (?)", name)
	if err != nil {
		return "", err
	}
	var id int64
	id, err = res.LastInsertId()
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(id, 10), nil
}

func (rr registrarRepo) SaveAndIdentifyTask(name, from, to string) (string, error) {
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
