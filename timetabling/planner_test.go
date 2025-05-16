package timetabling

import "testing"

func TestPlansZeroSlotsForNoPerformers(t *testing.T) {
	result := plan(nil, nil)
	if 0 != len(result) {
		t.Errorf("Result must be zero slots for no performers, actual is [%v]", result)
	}
}

func TestPlansFirstSlotForSingleTaskAndSinglePerformer(t *testing.T) {
	p, tasks := newSinglePerformer("06:00-07:00"), newSingleTask()
	result := plan(p, tasks)
	if 1 != len(result) || "06:00" != result[0].From() {
		t.Errorf("Result must be a single job for single available performer, actual is [%v]", result)
	}
}

func TestPlansZeroSlotsWhenPerformerDoesntHaveOpenPeriod(t *testing.T) {
	p, tasks := newSinglePerformer(""), newSingleTask()
	result := plan(p, tasks)
	if 0 != len(result) {
		t.Errorf("Result must be zero slots for no performers, actual is [%v]", result)
	}
}

func newSinglePerformer(openPeriod string) []performer {
	openPeriods := []period{}
	if 11 == len(openPeriod) {
		openPeriods = append(openPeriods, period{openPeriod[0:5], openPeriod[6:11]})
	}
	return []performer{performer{openPeriods}}
}

func newSingleTask() []task {
	return []task{task{"discuss nature"}}
}
