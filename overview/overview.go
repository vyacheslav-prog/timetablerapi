package overview

import (
	"context"
	"errors"
)

type repository interface {
	FetchPerformerBoard(context.Context, string) (result *PerformerBoard, err error)
}

type Overview struct {
	Repo repository
}

var (
	errViewPerformerBoard = errors.New("unable to view a performer board")
)

func (s Overview) ViewPerformerBoard(ctx context.Context, id string) (string, error) {
	result, err := s.Repo.FetchPerformerBoard(ctx, id)
	if err != nil {
		return "", errors.Join(errViewPerformerBoard, err)
	}
	return result.id, nil
}
