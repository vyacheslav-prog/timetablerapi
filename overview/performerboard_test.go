package overview

import "testing"

func TestShowsPerformerBoardForNoSlots(t *testing.T) {
	p := &performer{}
	result := p.getSlots()
	if len(result) != 0 {
		t.Errorf("Result must be empty slots for no periods, actual is [%v]", result)
	}
}
