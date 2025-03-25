package overview

type performerBoard struct {
	createdAt, id string
}

type overviewRepo interface {
	FetchPerformerBoard(id string) (result *performerBoard, err error)
}

type overviewService struct {
	repo overviewRepo
}

func (s *overviewService) viewPerformerBoard(id string) string {
	return s.repo.fetchPerformerBoard(id)
}
