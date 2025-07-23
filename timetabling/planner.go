package timetabling

type job struct {
	from, performer string
}

type task struct {
	from, title, to string
}

func plan(performers []performer, tasks []task) []job {
	var result []job
	for _, t := range tasks {
		tp := t.period()
		candidate := ""
		for pIndex := 0; len(performers) != pIndex && candidate == ""; pIndex += 1 {
			if ap := performers[pIndex].findAvailablePeriod(tp); nil != ap {
				candidate = performers[pIndex].name
			}
		}
		if "" != candidate {
			result = append(result, job{tp.from, candidate})
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
