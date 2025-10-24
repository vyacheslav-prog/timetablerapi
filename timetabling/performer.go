package timetabling

type performer struct {
	name         string
	emptyPeriods []period
}

func (p *performer) findAvailablePeriod(r *period) *period {
	if len(p.emptyPeriods) != 0 {
		if firstPeriod := p.emptyPeriods[0]; r == nil || (firstPeriod.from <= r.from && r.to <= firstPeriod.to) {
			return &p.emptyPeriods[0]
		}
	}
	return nil
}
