package timetabling

type job struct {
	from, performer string
}

type task struct {
	from, title, to string
}

func plan(recipients []performer, tasks []task) []job {
	var result []job
	busyPerformers := make(map[string]string)
	for _, t := range tasks {
		perf, tp := "", t.period()
		for rIndex := 0; len(recipients) != rIndex && perf == ""; rIndex += 1 {
			if r := recipients[rIndex]; nil != r.findAvailablePeriod(tp) {
				if busyFrom, isBusy := busyPerformers[r.name]; true != isBusy || busyFrom != t.from {
					perf = r.name
					busyPerformers[perf] = t.from
				}
			}
		}
		if "" != perf {
			result = append(result, job{tp.from, perf})
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
