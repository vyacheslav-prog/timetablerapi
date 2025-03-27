package overview

type PerformerBoard struct {
	createdAt, id string
}

func NewPerformerBoard(createdAt, id string) *PerformerBoard {
	return &PerformerBoard{createdAt, id}
}

type overviewRepo interface {
	FetchPerformerBoard(id string) (result *PerformerBoard, err error)
}

type OverviewService struct {
	repo overviewRepo
}

func (s *OverviewService) ViewPerformerBoard(id string) string {
	result, _ := s.repo.FetchPerformerBoard(id)
	return result.id
}
