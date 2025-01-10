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
	p := &performer{"", []slot{{"09:00", "10:00"}, {"10:00", "11:00"}}}
	result := p.getSlots()
	if expectedFirstFrom := "09:00"; len(result) != 2 || result[0].from != expectedFirstFrom {
		t.Errorf("Result must have two slots with first slot on [%v], actual slots is [%v]", expectedFirstFrom, result)
	}
}

func TestShowsPerformerBoardWithPersonalTitle(t *testing.T) {
	pn := "John Doe"
	p := &performer{pn, []slot{}}
	result := p.viewBoardTitle()
	if expected := "Board for John Doe"; expected != result {
		t.Errorf("Board title must be [%v] for performer [%v], actual is [%v]", expected, pn, result)
	}
}
