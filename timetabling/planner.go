package timetabling

type task struct {
	title string
}

func plan(p []performer, t []task) []int {
	if 0 != len(p) {
		return []int{0}
	}
	return nil
}
