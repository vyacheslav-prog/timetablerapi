package overview

type PerformerBoard struct {
	createdAt, id string
}

type overviewRepo interface {
	FetchPerformerBoard(id string) (result *PerformerBoard, err error)
}

type overviewService struct {
	repo overviewRepo
}

func (s *overviewService) viewPerformerBoard(id string) string {
	result, _ := s.repo.FetchPerformerBoard(id)
	return result.id
}
