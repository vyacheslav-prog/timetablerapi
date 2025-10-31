package timetabling

import "testing"

func TestFindsNoPeriodForEmptyAvailablePeriods(t *testing.T) {
	p := performer{}
	result := p.findAvailablePeriod(nil)
	if result != nil {
		t.Errorf("Result must be nil for empty available periods, actual is [%v]", result)
	}
}

func TestFindsFirstAvailablePeriod(t *testing.T) {
	p := &performer{"John", []period{{"06:00", "07:00"}}}
	result := p.findAvailablePeriod(nil)
	if result == nil || result.from != "06:00" {
		t.Errorf("Result must be not nil for not empty available periods, actual is [%v]", result)
	}
}

func TestFindsNoPeriodWhenRequestedPeriodIsTooLong(t *testing.T) {
	p, r := &performer{"John", []period{{"08:00", "10:00"}}}, &period{"09:00", "12:00"}
	result := p.findAvailablePeriod(r)
	if result != nil {
		t.Errorf("Result must be nil for too long request [%v], actual is [%v]", r, result)
	}
}
