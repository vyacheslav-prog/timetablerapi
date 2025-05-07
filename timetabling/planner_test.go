package timetabling

import "testing"

func TestPlansZeroSlotsForNoPerformers(t *testing.T) {
	result := plan(nil)
	if 0 != len(result) {
		t.Errorf("Result must be zero slots for no performers, actual is [%v]", result)
	}
}

func TestPlansFirstSlotForSingleTaskAndSinglePerformer(t *testing.T) {
	t, p := []task{task{"discuss nature"}}, []performer{performer{[]period{period{"06:00", "07:00"}}}}
	result := plan(p, t)
	if 1 != len(result) {
		t.Errorf("Result must be a single job for single available performer, actual is [%v]", result)
	}
}
