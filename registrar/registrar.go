package registrar

import (
	"context"
	"errors"
	"fmt"
)

type repository interface {
	SaveAndIdentifyLayout(string) (string, error)
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
	identity, err := r.Repo.SaveAndIdentifyLayout(mode)
	if err != nil {
		return "", errors.Join(errRegistrar, err)
	}
	return identity, nil
}

func (r Registrar) AddPerformer(ctx context.Context, name string) (string, error) {
	identity, err := r.Repo.SaveAndIdentifyPerformer(ctx, name)
	if err != nil {
		return "", errors.Join(errRegistrar, err)
	}
	return identity, nil
}

func (r Registrar) AddPeriod(ctx context.Context, from, to string) (string, error) {
	return "", fmt.Errorf("%w: not implemented", errRegistrar)
}

func (r Registrar) AddTask(ctx context.Context, name, from, to string) (string, error) {
	identity, err := r.Repo.SaveAndIdentifyTask(ctx, name, from, to)
	if err != nil {
		return "", errors.Join(errRegistrar, err)
	}
	return identity, nil
}
