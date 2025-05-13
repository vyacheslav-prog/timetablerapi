package timetabling

type job struct {
	from string
}

type task struct {
	title string
}

func plan(p []performer, t []task) []job {
	if 0 != len(p) {
		return []job{job{p[0].emptyPeriods[0].from}}
	}
	return nil
}

func (j *job) From() string {
	return j.from
}
