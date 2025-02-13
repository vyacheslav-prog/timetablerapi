//go:build testdb

package main

import (
	"database/sql"
	"os"
	"testing"
)

func TestFetchsNoPerformerBoardForEmptyRequest(t *testing.T) {
	db := openDBConnect(t)
	sut, err := newOverviewRepo(db)
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
	db, id := openDBConnect(t), "2861ff45-526f-4618-9b7a-09e581cb2113"
	defer db.Close()
	sut, err := newOverviewRepo(nil)
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

func openDBConnect(t *testing.T) *sql.DB {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("failed to connect to database")
	}
	return db
}

func seedFakePerformerBoard(db *sql.DB, boardId string) bool {
	query := "insert into performer_boards (id) values ($1)"
	result, _ := db.Exec(query, boardId)
	touchedRows, _ := result.RowsAffected()
	return 1 == touchedRows
}
