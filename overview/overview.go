package overview

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
