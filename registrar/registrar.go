package registrar

import (
	"errors"
	"fmt"
)

type repository interface {
	SaveAndIdentifyLayout(string) (string, error)
	SaveAndIdentifyPerformer(string) (string, error)
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

func (r Registrar) AddPerformer(name string) (string, error) {
	identity, err := r.Repo.SaveAndIdentifyPerformer(name)
	if err != nil {
		return "", errors.Join(errRegistrar, err)
	}
	return identity, nil
}

func (r Registrar) AddPeriod(from, to string) (string, error) {
	return "", fmt.Errorf("%w: not implemented", errRegistrar)
}
