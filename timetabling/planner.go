package timetabling

type job struct {
	from string
}

type task struct {
	from, title, to string
}

func plan(performers []performer, tasks []task) []job {
	var result []job
planning:
	for _, t := range tasks {
		for _, p := range performers {
			ap := p.findAvailablePeriod(&period{t.from, t.to})
			if nil != ap {
				result = append(result, job{t.from})
				break planning
			}
		}
	}
	return result
}

func (j *job) startAt() string {
	return j.from
}
