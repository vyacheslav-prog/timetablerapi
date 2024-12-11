package timetabling

import "testing"

func TestSchedulesNoneForNoPerformers(t *testing.T) {
	result := schedule()
	if result != nil {
		t.Errorf("Jobs must be nil for no performers, actual is [%v]", result)
	}
}
