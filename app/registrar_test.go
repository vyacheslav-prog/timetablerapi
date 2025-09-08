//go:build testdb

package main

import (
	"testing"
)

func TestIdentifiesPerformerForAdding(t *testing.T) {
	repo := registrarRepo{}
	_, repoErr := repo.SaveAndIdentifyPerformer("John")
	if repoErr != nil {
		t.Error("Couldn't identify new performer", repoErr)
	}
}
