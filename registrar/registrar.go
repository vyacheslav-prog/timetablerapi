package registrar

import (
	"errors"
	"fmt"
)

type repository interface {
	SaveAndIdentifyPerformer(string) (string, error)
}

type Registrar struct {
	Repo repository
}

var (
	errRegistrar = errors.New("registrar error is occured")
)

func (r Registrar) AddLayout(mode string) (string, error) {
	return "", fmt.Errorf("%w: not implemented", errRegistrar)
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
