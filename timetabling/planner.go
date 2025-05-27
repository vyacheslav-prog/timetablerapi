package timetabling

type job struct {
	from string
}

type task struct {
	from, title, to string
}

func plan(performers []performer, tasks []task) []job {
	var result []job
	for _, t := range tasks {
		for _, p := range performers {
			for _, ep := range p.emptyPeriods {
				if t.from == ep.from && t.to == ep.to {
					result = append(result, job{t.from})
				}
			}
		}
	}
	return result
}

func (j *job) From() string {
	return j.from
}
