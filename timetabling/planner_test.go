package timetabling

import "testing"

func TestPlansZeroSlotsForNoPerformers(t *testing.T) {
	result := plan(nil, nil)
	if len(result) != 0 {
		t.Errorf("Result must be zero slots for no performers, actual is [%v]", result)
	}
}

func TestPlansFirstSlotForSingleTaskAndSinglePerformer(t *testing.T) {
	p, tasks := newSinglePerformer("06:00-07:00"), newSingleTask("06:00-07:00")
	result := plan(p, tasks)
	if len(result) != 1 || result[0].startAt() != "06:00" {
		t.Errorf("Result must be a single job for single available performer, actual is [%v]", result)
	}
}

func TestPlansZeroSlotsWhenPerformerDoesntHaveOpenPeriod(t *testing.T) {
	p, tasks := newSinglePerformer(""), newSingleTask("")
	result := plan(p, tasks)
	if len(result) != 0 {
		t.Errorf("Result must be zero slots for no performers, actual is [%v]", result)
	}
}

func TestPlansZeroSlotsWhenTaskIsNotFitIntoPeriod(t *testing.T) {
	p, tasks := newSinglePerformer("08:00-08:15"), newSingleTask("08:00-08:45")
	result := plan(p, tasks)
	if len(result) != 0 {
		t.Errorf("Result must be zero slots for no performers, actual is [%v]", result)
	}
}

func TestPlansSingleJobWhenPerformerPeriodIsLongerThanTaskPeriod(t *testing.T) {
	p, tasks := newSinglePerformer("06:00-07:00"), newSingleTask("06:30-07:00")
	result := plan(p, tasks)
	if len(result) != 1 || result[0].startAt() != "06:30" {
		t.Errorf("Result must be a single job for single available performer, actual is [%v]", result)
	}
}

func TestPlansSingleJobForTwoPerformersWithSameOpenPeriod(t *testing.T) {
	p, tasks := []performer{newPerformer("Dave", "06:00-07:00"), newPerformer("Kate", "06:00-07:00")}, newSingleTask("06:00-07:00")
	result := plan(p, tasks)
	if len(result) != 1 {
		t.Errorf("Result must be a single job for many performers, actual is [%v]", result)
	}
}

func TestPlansTwoJobsForSinglePerfromerWithoutOverlap(t *testing.T) {
	p, tasks := newSinglePerformer("09:00-15:00"), append(newSingleTask("09:00-10:00"), newSingleTask("14:00-15:00")[0])
	result := plan(p, tasks)
	if len(result) != 2 {
		t.Errorf("Result must contain two job for performers [%v] and tasks [%v], actual is [%v]", p, tasks, result)
	}
}

func TestPlansTwoSameTimeJobsForTwoPerformers(t *testing.T) {
	p := []performer{newPerformer("Dave", "06:00-07:00"), newPerformer("Kate", "06:00-07:00")}
	tasks := append(newSingleTask("06:00-07:00"), newSingleTask("06:00-07:00")[0])
	result := plan(p, tasks)
	if len(result) != 2 || result[0].performer == result[1].performer {
		t.Errorf("Result must contain two different job for performers [%v] and tasks [%v], actual is [%v]", p, tasks, result)
	}
}

func newPerformer(id, openPeriod string) performer {
	openPeriods := []period{}
	if len(openPeriod) == 11 {
		openPeriods = append(openPeriods, period{openPeriod[0:5], openPeriod[6:11]})
	}
	return performer{id, openPeriods}
}

func newSinglePerformer(openPeriod string) []performer {
	return []performer{newPerformer("John", openPeriod)}
}

func newSingleTask(fromTo string) []task {
	subject := "discuss nature"
	if len(fromTo) == 11 {
		return []task{{fromTo[0:5], subject, fromTo[6:11]}}
	}
	return []task{{"00:00", "00:00", subject}}
}
