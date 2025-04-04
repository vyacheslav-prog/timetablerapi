package overview

type PerformerBoard struct {
	createdAt, id, title string
}

func (pb *PerformerBoard) Title() string {
	return pb.title
}

func NewPerformerBoard(createdAt, id, title string) *PerformerBoard {
	return &PerformerBoard{createdAt, id, title}
}

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
