package timetabling

type job struct {
	from string
}

type task struct {
	from, title, to string
}

func plan(performers []performer, tasks []task) []job {
	if 1 != len(tasks) {
		return nil
	}
	taskFrom, taskTo := tasks[0].from, tasks[0].to
	for _, p := range performers {
		for _, ep := range p.emptyPeriods {
			if taskFrom == ep.from && taskTo == ep.to {
				return []job{job{taskFrom}}
			}
		}
	}
	return nil
}

func (j *job) From() string {
	return j.from
}
