package overview

type performer struct {
	name      string
	openSlots []slot
}

type slot struct {
	from, to string
}

func (p *performer) getSlots() []slot {
	return p.openSlots
}

func (p *performer) viewBoardTitle() string {
	return "Board for " + p.name
}
