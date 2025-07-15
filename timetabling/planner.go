package timetabling

type job struct {
	from string
}

type task struct {
	from, title, to string
}

func plan(performers []performer, tasks []task) []job {
	var result []job
	lastPerformerIndex := -1
	for _, t := range tasks {
		tp := t.period()
		for index, p := range performers {
			if ap := p.findAvailablePeriod(tp); lastPerformerIndex != index && nil != ap {
				result = append(result, job{tp.from})
				lastPerformerIndex = index
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
