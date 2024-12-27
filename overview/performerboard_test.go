package overview

import "testing"

func TestShowsPerformerBoardForNoSlots(t *testing.T) {
	p := &performer{}
	result := p.getSlots()
	if len(result) != 0 {
		t.Errorf("Result must be empty slots for no periods, actual is [%v]", result)
	}
}

func TestShowsPerformerBoardWithOpenSlots(t *testing.T) {
	p := &performer{[]slot{{"09:00", "10:00"}, {"10:00", "11:00"}}}
	result := p.getSlots()
	if expectedFirstFrom := "09:00"; len(result) != 2 || result[0].from != expectedFirstFrom {
		t.Errorf("Result must have two slots with first slot on [%v], actual slots is [%v]", expectedFirstFrom, result)
	}
}
