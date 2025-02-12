package main

import (
	"database/sql"
	"os"
	"testing"
)

func TestFetchsNoPerformerBoardForEmptyRequest(t *testing.T) {
	sut, err := newOverviewRepo()
	if err != nil {
		t.Errorf("failed init overview repo: %v", err)
	}
	result := sut.fetchPerformerBoard("")
	if nil != result {
		t.Errorf("Result must be nil for empty performer request")
	}
}

func TestFetchsPerformerBoardByIdentity(t *testing.T) {
	t.Skip()
	connStr, id := os.Getenv("DATABASE_URL"), "2861ff45-526f-4618-9b7a-09e581cb2113"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Failed to connect to database [%v]", connStr)
	}
	defer db.Close()
	sut, err := newOverviewRepo()
	if err != nil {
		t.Errorf("failed init overview repo: %v", err)
	}
	isSown := seedFakePerformerBoard(nil, id)
	if isSown != true {
		t.Error("Could not be seed a fake performer board")
	}
	result := sut.fetchPerformerBoard(id)
	if nil == result {
		t.Errorf("Result must be not nil for [%v] performer board id, actual is [%v]", id, result)
	}
}

func seedFakePerformerBoard(db *sql.DB, boardId string) bool {
	query := "insert into performer_boards (id) values ($1)"
	result, _ := db.Exec(query, boardId)
	touchedRows, _ := result.RowsAffected()
	return 1 == touchedRows
}
