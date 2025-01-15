package main

import "testing"

func TestFetchsPerformerBoard(t *testing.T) {
	sut := newOverviewRepo()
	result := sut.fetchPerformerBoard()
	if nil != result {
		t.Errorf("Result must be nil for empty performer request")
	}
}
