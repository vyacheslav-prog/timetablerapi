package timetabling

type job struct {
	from string
}

type task struct {
	from, title, to string
}

func (t *task) fit(p period) bool {
	return p.from <= t.from && t.to <= p.to
}

func plan(performers []performer, tasks []task) []job {
	var result []job
planning:
	for _, t := range tasks {
		for _, p := range performers {
			for _, ep := range p.emptyPeriods {
				if t.fit(ep) {
					result = append(result, job{t.from})
					break planning
				}
			}
		}
	}
	return result
}

func (j *job) startAt() string {
	return j.from
}
