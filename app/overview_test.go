//go:build testdb

package main

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

func TestFetchsNoPerformerBoardForEmptyRequest(t *testing.T) {
	db := openDBConnect(t)
	defer db.Close()
	sut, err := newOverviewRepo(t.Context(), db)
	if err != nil {
		t.Error("failed init overview repo:", err)
	}
	result, _, _ := sut.fetchPerformerBoard("")
	if nil != result {
		t.Error("Result must be nil for empty performer request")
	}
}

func TestFetchsPerformerBoardByIdentity(t *testing.T) {
	db, id := openDBConnect(t), "2861ff45-526f-4618-9b7a-09e581cb2113"
	defer db.Close()
	sut, err := newOverviewRepo(t.Context(), db)
	if err != nil {
		t.Error("failed init overview repo:", err)
	}
	var deleteBoard func()
	deleteBoard, err = seedFakePerformerBoard(db, id)
	if err != nil {
		t.Error("Could not be seed a fake performer board:", err)
	}
	defer deleteBoard()
	var result *string
	result, _, err = sut.fetchPerformerBoard(id)
	if err != nil {
		t.Error("Could not fetch a performer board:", err)
	}
	if nil == result {
		t.Errorf("Result must be not nil for [%v] performer board id, actual is [%v]", id, result)
	}
}

func openDBConnect(t *testing.T) *sql.DB {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Error("failed connection to database:", err)
	}
	return db
}

func seedFakePerformerBoard(db *sql.DB, boardId string) (func(), error) {
	_, err := db.Exec("insert into performer_boards (id) values ($1);", boardId)
	return func() {
		db.Exec("delete from performer_boards where id = $1;", boardId)
	}, err
}
