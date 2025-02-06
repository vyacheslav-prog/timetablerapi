package main

import (
	"database/sql"
	"os"
	"testing"
)

func TestFetchsNoPerformerBoardForEmptyRequest(t *testing.T) {
	sut := newOverviewRepo()
	result := sut.fetchPerformerBoard()
	if nil != result {
		t.Errorf("Result must be nil for empty performer request")
	}
}

func TestFetchsPerformerBoardByIdentity(t *testing.T) {
	connStr, id := os.Getenv("DATABASE_URL"), "2861ff45-526f-4618-9b7a-09e581cb2113"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Failed to connect to database [%v]", connStr)
	}
	defer db.Close()
	sut, isSown := newOverviewRepo(), seedFakePerformerBoard(id)
	result := sut.fetchPerformerBoard(id)
	if nil == result {
		t.Errorf("Result must be not nil for [%v] performer board id, actual is [%v]", id, result)
	}
}

func seedFakePerformerBoard(db *sql.DB, boardId string) bool {
	query := "insert into performer_boards (id) values ($1)"
	result, err := db.Exec(query, boardId)
	return 1 == result.RowsAffected()
}
