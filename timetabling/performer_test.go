package timetabling

import "testing"

func TestFindsNoPeriodForEmptyAvailablePeriods(t *testing.T) {
	p := performer{}
	result := p.findAvailablePeriod()
	if result != nil {
		t.Errorf("Result must be nil for empty available periods, actual is [%v]", result)
	}
}

func TestFindsFirstAvailablePeriod(t *testing.T) {
	p := &performer{[]period{period{}}}
	result := p.findAvailablePeriod()
	if result == nil {
		t.Errorf("Result must be not nil for not empty available periods, actual is [%v]", result)
	}
}
