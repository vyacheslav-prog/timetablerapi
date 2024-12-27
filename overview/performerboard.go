package overview

type performer struct {
	openSlots []slot
}

type slot struct {
	from, to string
}

func (p *performer) getSlots() []slot {
	return nil
}
