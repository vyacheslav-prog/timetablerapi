package timetabling

type performer struct {
	emptyPeriods []period
}

type period struct {
	from, to string
}

func (p *performer) findAvailablePeriod() *period {
	if p.emptyPeriods != nil {
		return &p.emptyPeriods[0]
	}
	return nil
}
