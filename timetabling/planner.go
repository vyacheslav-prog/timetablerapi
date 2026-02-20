package timetabling

type job struct {
	from, performer string
}

type task struct {
	from, title, to string
}

func plan(recipients []performer, tasks []task) []job {
	var result []job
	busyPrfs := make(map[string]string)
	for _, t := range tasks {
		prf, tp := "", t.period()
		for _, r := range recipients {
			busyFrom, isBusy := busyPrfs[r.name]
			if (!isBusy || busyFrom != t.from) && r.findAvailablePeriod(tp) != nil {
				prf = r.name
				busyPrfs[prf] = t.from
				break
			}
		}
		if prf != "" {
			result = append(result, job{tp.from, prf})
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
