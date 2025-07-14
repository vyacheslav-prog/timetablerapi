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
	if 1 != len(result) || "06:00" != result[0].startAt() {
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
	if 1 != len(result) || "06:30" != result[0].startAt() {
		t.Errorf("Result must be a single job for single available performer, actual is [%v]", result)
	}
}

func TestPlansSingleJobForTwoPerformersWithSameOpenPeriod(t *testing.T) {
	p, tasks := []performer{newPerformer("06:00-07:00"), newPerformer("06:00-07:00")}, newSingleTask("06:00-07:00")
	result := plan(p, tasks)
	if 1 != len(result) {
		t.Errorf("Result must be a single job for many performers, actual is [%v]", result)
	}
}

func TestPlansTwoJobsForSinglePerfromerWithoutOverlap(t *testing.T) {
	p, tasks := newSinglePerformer("09:00-15:00"), append(newSingleTask("09:00-10:00"), newSingleTask("14:00-15:00")[0])
	result := plan(p, tasks)
	if 2 != len(result) {
		t.Errorf("Result must contain two job for performers [%v] and tasks [%v], actual is [%v]", p, tasks, result)
	}
}

func newPerformer(openPeriod string) performer {
	openPeriods := []period{}
	if 11 == len(openPeriod) {
		openPeriods = append(openPeriods, period{openPeriod[0:5], openPeriod[6:11]})
	}
	return performer{openPeriods}
}

func newSinglePerformer(openPeriod string) []performer {
	return []performer{newPerformer(openPeriod)}
}

func newSingleTask(fromTo string) []task {
	subject := "discuss nature"
	if 11 == len(fromTo) {
		return []task{task{fromTo[0:5], subject, fromTo[6:11]}}
	}
	return []task{task{"00:00", "00:00", subject}}
}
