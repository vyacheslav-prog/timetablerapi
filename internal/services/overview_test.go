//go:build testdb

package services

import (
	"database/sql"
	"testing"

	"timetablerapi/overview"
)

func TestFetchsNoPerformerBoardForEmptyRequest(t *testing.T) {
	db, err := openDBConnect("postgres://testuser:testpassword@postgres:5432/testdb?sslmode=disable")
	if err != nil {
		t.Error("failed open connection to database:", err)
		return
	}
	defer func() {
		if err = db.Close(); err != nil {
			t.Error("failed close connection to database:", err)
		}
	}()
	sut, err := newOverviewRepo(t.Context(), db, &dbMigrate{db, countTableByNameQuery})
	if err != nil {
		t.Error("failed init overview repo:", err)
		return
	}
	result, err := sut.FetchPerformerBoard(t.Context(), "")
	if err != nil {
		t.Error("failed fetch performer board:", err)
		return
	}
	if result != nil {
		t.Errorf("Result must be nil for empty performer request, actual is [%v]", *result)
	}
}

func TestFetchsPerformerBoardByIdentity(t *testing.T) {
	db, err := openDBConnect("postgres://testuser:testpassword@postgres:5432/testdb?sslmode=disable")
	if err != nil {
		t.Error("failed open connection to database:", err)
		return
	}
	id, title := "2861ff45-526f-4618-9b7a-09e581cb2113", "my board"
	defer func() {
		if err = db.Close(); err != nil {
			t.Error("failed close connection to database:", err)
		}
	}()
	sut, err := newOverviewRepo(t.Context(), db, &dbMigrate{db, countTableByNameQuery})
	if err != nil {
		t.Error("failed init overview repo:", err)
		return
	}
	deleteBoard := seedFakePerformerBoard(t, db, id, title)
	defer deleteBoard()
	var result *overview.PerformerBoard
	result, err = sut.FetchPerformerBoard(t.Context(), id)
	if err != nil {
		t.Error("Could not fetch a performer board:", err)
		return
	}
	if title != result.Title() {
		t.Errorf("Result for board fetching must have title [%v], actual board is [%v]", title, result)
	}
}

func seedFakePerformerBoard(t *testing.T, db *sql.DB, boardId, title string) func() {
	_, err := db.ExecContext(t.Context(), "insert into performer_boards (id, title) values ($1, $2);", boardId, title)
	if err != nil {
		t.Fatal("Could not be seed a fake performer board:", err)
	}
	return func() {
		_, err = db.ExecContext(t.Context(), "delete from performer_boards where id = $1;", boardId)
		if err != nil {
			t.Fatal("Could not be delete a fake performer board:", err)
		}
	}
}
