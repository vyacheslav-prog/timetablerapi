//go:build testdb

package main

import (
	"testing"
)

func TestIdentifiesPerformerForAdding(t *testing.T) {
	dbConn := openDBConnect(t)
	defer dbConn.Close()
	repo := newRegistrarRepo(t.Context(), dbConn)
	id, repoErr := repo.SaveAndIdentifyPerformer("John")
	if repoErr != nil {
		t.Error("couldn't identify new performer:", repoErr)
	}
	if "" == id {
		t.Error("identity must be not empty string")
	}
	existsRow := dbConn.QueryRowContext(t.Context(), "select id from performers where id = $1", id)
	if rowErr := existsRow.Err(); rowErr != nil {
		t.Error("couldn't check existing for a saved performer:", rowErr)
	}
}
