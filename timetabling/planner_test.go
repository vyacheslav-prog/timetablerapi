package timetabling

import "testing"

func TestPlansZeroSlotsForNoPerformers(t *testing.T) {
	result := plan(nil, nil)
	if 0 != len(result) {
		t.Errorf("Result must be zero slots for no performers, actual is [%v]", result)
	}
}

func TestPlansFirstSlotForSingleTaskAndSinglePerformer(t *testing.T) {
	p, tasks := newSinglePerformer("06:00-07:00"), newSingleTask("06:00-07:00")
	result := plan(p, tasks)
	if 1 != len(result) || "06:00" != result[0].From() {
		t.Errorf("Result must be a single job for single available performer, actual is [%v]", result)
	}
}

func TestPlansZeroSlotsWhenPerformerDoesntHaveOpenPeriod(t *testing.T) {
	p, tasks := newSinglePerformer(""), newSingleTask("")
	result := plan(p, tasks)
	if 0 != len(result) {
		t.Errorf("Result must be zero slots for no performers, actual is [%v]", result)
	}
}

func TestPlansZeroSlotsWhenTaskIsNotFitIntoPeriod(t *testing.T) {
	p, tasks := newSinglePerformer("08:00-08:15"), newSingleTask("08:00-08:45")
	result := plan(p, tasks)
	if 0 != len(result) {
		t.Errorf("Result must be zero slots for no performers, actual is [%v]", result)
	}
}

func TestPlansSingleJobWhenPerformerPeriodIsLongerThanTaskPeriod(t *testing.T) {
	p, tasks := newSinglePerformer("06:00-07:00"), newSingleTask("06:30-07:00")
	result := plan(p, tasks)
	if 1 != len(result) || "06:30" != result[0].From() {
		t.Errorf("Result must be a single job for single available performer, actual is [%v]", result)
	}
}

func newSinglePerformer(openPeriod string) []performer {
	openPeriods := []period{}
	if 11 == len(openPeriod) {
		openPeriods = append(openPeriods, period{openPeriod[0:5], openPeriod[6:11]})
	}
	return []performer{performer{openPeriods}}
}

func newSingleTask(fromTo string) []task {
	subject := "discuss nature"
	if 11 == len(fromTo) {
		return []task{task{fromTo[0:5], subject, fromTo[6:11]}}
	}
	return []task{task{"00:00", "00:00", subject}}
}
