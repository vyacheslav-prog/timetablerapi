package timetabling

type job struct {
	from string
}

type task struct {
	title string
}

func plan(p []performer, t []task) []job {
	if 0 != len(p) {
		for _, p := range p[0].emptyPeriods {
			return []job{job{p.from}}
		}
	}
	return nil
}

func (j *job) From() string {
	return j.from
}
