package timetabling

type performer struct {
	emptyPeriods []period
	name         string
}

func (p *performer) findAvailablePeriod(r *period) *period {
	if 0 != len(p.emptyPeriods) {
		if firstPeriod := p.emptyPeriods[0]; r == nil || (firstPeriod.from <= r.from && r.to <= firstPeriod.to) {
			return &p.emptyPeriods[0]
		}
	}
	return nil
}
