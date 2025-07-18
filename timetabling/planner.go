package timetabling

type job struct {
	from, performer string
}

type task struct {
	from, title, to string
}

func plan(performers []performer, tasks []task) []job {
	var result []job
	for _, t := range tasks {
		tp := t.period()
		for _, p := range performers {
			if ap := p.findAvailablePeriod(tp); nil != ap {
				result = append(result, job{tp.from, p.name})
				break
			}
		}
	}
	return result
}

func (t *task) period() *period {
	return &period{t.from, t.to}
}

func (j *job) startAt() string {
	return j.from
}
