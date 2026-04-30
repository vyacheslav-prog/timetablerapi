package registrar

import (
	"context"
	"errors"
	"fmt"
)

type Performer struct {
	From string
	Name string
	To   string
}

type repository interface {
	SaveAndIdentifyLayout(context.Context, string) (string, error)
	SaveAndIdentifyPerformer(context.Context, string) (string, error)
	SaveAndIdentifyTask(context.Context, string, string, string) (string, error)
}

type Registrar struct {
	Repo repository
}

var (
	errRegistrar = errors.New("registrar error is occurred")
)

func (r Registrar) AddLayout(mode string) (string, error) {
	identity, err := r.Repo.SaveAndIdentifyLayout(context.Background(), mode)
	if err != nil {
		return "", fmt.Errorf("%w: %w", errRegistrar, err)
	}
	return identity, nil
}

func (r Registrar) AddPerformer(ctx context.Context, prf Performer) (string, error) {
	identity, err := r.Repo.SaveAndIdentifyPerformer(ctx, prf.Name)
	if err != nil {
		return "", fmt.Errorf("%w: %w", errRegistrar, err)
	}
	return identity, nil
}

func (r Registrar) AddTask(ctx context.Context, name, from, to string) (string, error) {
	identity, err := r.Repo.SaveAndIdentifyTask(ctx, name, from, to)
	if err != nil {
		return "", fmt.Errorf("%w: %w", errRegistrar, err)
	}
	return identity, nil
}
