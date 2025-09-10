//go:build testdb

package main

import (
	"testing"
)

func TestIdentifiesPerformerForAdding(t *testing.T) {
	repo := registrarRepo{}
	id, repoErr := repo.SaveAndIdentifyPerformer("John")
	if repoErr != nil {
		t.Error("Couldn't identify new performer:", repoErr)
	}
	if "" == id {
		t.Error("Identity must be not empty string")
	}
}
