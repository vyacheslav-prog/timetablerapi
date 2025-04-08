package overview

type repository interface {
	FetchPerformerBoard(id string) (result *PerformerBoard, err error)
}

type Overview struct {
	Repo repository
}

func (s Overview) ViewPerformerBoard(id string) string {
	result, _ := s.Repo.FetchPerformerBoard(id)
	return result.id
}
