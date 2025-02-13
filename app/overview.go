package main

import (
	"database/sql"
	"fmt"
)

type overviewRepo struct {
}

const performerBoardsSchema = `
	create table if not exists performer_boards (
		id text primary key,
		created_at datetime default current_timestamp,
	);
`

func (r *overviewRepo) fetchPerformerBoard(id string) *int {
	return nil
}

func newOverviewRepo(db *sql.DB) (*overviewRepo, error) {
	if db == nil {
		return nil, fmt.Errorf("not connection for server")
	}
	return nil, nil
}

type overviewService struct {
	repo overviewRepo
}
