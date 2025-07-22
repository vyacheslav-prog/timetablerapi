package timetabling

type job struct {
	from, performer string
}

type task struct {
	from, title, to string
}

func plan(performers []performer, tasks []task) []job {
	var result []job
	busyPerformers := make(map[string]int)
	for _, t := range tasks {
		tp := t.period()
		for _, p := range performers {
			if _, isBusy := busyPerformers[p.name]; isBusy {
				break
			}
			if ap := p.findAvailablePeriod(tp); nil != ap {
				result = append(result, job{tp.from, p.name})
				busyPerformers[p.name] = 1
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
