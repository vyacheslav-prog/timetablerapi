package timetabling

type job struct {
	from string
}

type task struct {
	from, title, to string
}

func plan(performers []performer, tasks []task) []job {
	var result []job
	for tc := 0; 0 == len(result) && tc != len(tasks); tc += 1 {
		tp := tasks[tc].period()
		for _, p := range performers {
			if ap := p.findAvailablePeriod(tp); nil != ap {
				result = append(result, job{tp.from})
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
