package overview

type PerformerBoard struct {
	createdAt, id, title string
}

func (pb *PerformerBoard) Title() string {
	return pb.title
}

func NewPerformerBoard(createdAt, id string) *PerformerBoard {
	return &PerformerBoard{createdAt, id, ""}
}

type repository interface {
	FetchPerformerBoard(id string) (result *PerformerBoard, err error)
}

type Overview struct {
	repo repository
}

func (s Overview) ViewPerformerBoard(id string) string {
	result, _ := s.repo.FetchPerformerBoard(id)
	return result.id
}
