package main

import "fmt"

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

func newOverviewRepo() (*overviewRepo, error) {
	return nil, fmt.Errorf("not connection for server")
}

type overviewService struct {
	repo overviewRepo
}
