//go:build testdb

package services

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"

	"timetablerapi/overview"
)

func TestFetchsNoPerformerBoardForEmptyRequest(t *testing.T) {
	db := openDBConnect(t)
	defer db.Close()
	sut, err := newOverviewRepo(t.Context(), db)
	if err != nil {
		t.Error("failed init overview repo:", err)
	}
	result, _ := sut.FetchPerformerBoard(t.Context(), "")
	if nil != result {
		t.Errorf("Result must be nil for empty performer request, actual is [%v]", *result)
	}
}

func TestFetchsPerformerBoardByIdentity(t *testing.T) {
	db, id, title := openDBConnect(t), "2861ff45-526f-4618-9b7a-09e581cb2113", "my board"
	defer db.Close()
	sut, err := newOverviewRepo(t.Context(), db)
	if err != nil {
		t.Error("failed init overview repo:", err)
	}
	var deleteBoard func()
	deleteBoard, err = seedFakePerformerBoard(db, id, title)
	if err != nil {
		t.Error("Could not be seed a fake performer board:", err)
	}
	defer deleteBoard()
	var result *overview.PerformerBoard
	result, err = sut.FetchPerformerBoard(t.Context(), id)
	if err != nil {
		t.Error("Could not fetch a performer board:", err)
	}
	if title != result.Title() {
		t.Errorf("Result for board fetching must have title [%v], actual board is [%v]", title, result)
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

func seedFakePerformerBoard(db *sql.DB, boardId, title string) (func(), error) {
	_, err := db.Exec("insert into performer_boards (id, title) values ($1, $2);", boardId, title)
	return func() {
		db.Exec("delete from performer_boards where id = $1;", boardId)
	}, err
}
