package timetabling

import "testing"

func TestPlansZeroSlotsForNoPerformers(t *testing.T) {
	result := plan(nil)
	if 0 != len(result) {
		t.Errorf("Result must be zero slots for no performers, actual is [%v]", result)
	}
}
