package timetabling

type job struct {
	from string
}

type task struct {
	title string
}

func plan(performers []performer, tasks []task) []job {
	for _, p := range plist {
		for _, ep := range p.emptyPeriods {
			return []job{job{ep.from}}
		}
	}
	return nil
}

func (j *job) From() string {
	return j.from
}
