package timetabling

type job struct {
	from string
}

type task struct {
	from, title, to string
}

func plan(performers []performer, tasks []task) []job {
	for _, p := range performers {
		for _, ep := range p.emptyPeriods {
			return []job{job{ep.from}}
		}
	}
	return nil
}

func (j *job) From() string {
	return j.from
}
