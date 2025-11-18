package overview

import "context"

type repository interface {
	FetchPerformerBoard(context.Context, string) (result *PerformerBoard, err error)
}

type Overview struct {
	Repo repository
}

func (s Overview) ViewPerformerBoard(ctx context.Context, id string) string {
	result, _ := s.Repo.FetchPerformerBoard(ctx, id)
	return result.id
}
