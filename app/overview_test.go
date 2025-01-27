package main

import "testing"

func TestFetchsNoPerformerBoardForEmptyRequest(t *testing.T) {
	sut := newOverviewRepo()
	result := sut.fetchPerformerBoard()
	if nil != result {
		t.Errorf("Result must be nil for empty performer request")
	}
}

func TestFetchsPerformerBoardByIdentity(t *testing.T) {
	id := "2861ff45-526f-4618-9b7a-09e581cb2113"
	sut, isSown := newOverviewRepo(), seedFakePerformerBoard(id)
	result := sut.fetchPerformerBoard(id)
	if nil == result {
		t.Errorf("Result must be not nil for [%v] performer board id, actual is [%v]", id, result)
	}
}
