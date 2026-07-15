//go:build testdb

package services

import (
	"testing"
)

func TestIdentifiesPerformerForAdding(t *testing.T) {
	dbConn := openDBConnect(t)
	defer dbConn.Close()
	repo, migrErr := newRegistrarRepo(t.Context(), dbConn, &dbMigrate{db, countTableByNameQuery})
	if migrErr != nil {
		t.Error("couldn't migrate for registrar repo:", migrErr)
		return
	}
	id, repoErr := repo.SaveAndIdentifyPerformer(t.Context(), "John")
	if repoErr != nil {
		t.Error("couldn't identify new performer:", repoErr)
		return
	}
	if "" == id {
		t.Error("identity must be not empty string")
		return
	}
	existsRow := dbConn.QueryRowContext(t.Context(), "select id from performers where id = $1", id)
	if rowErr := existsRow.Err(); rowErr != nil {
		t.Error("couldn't check existing for a saved performer:", rowErr)
	}
}
