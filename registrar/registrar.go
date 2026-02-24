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
	errAddPerfomer = errors.New("unable to add a performer")
	errAddPeriod   = errors.New("unable to add a period")
)

func (r Registrar) AddLayout(mode string) (string, error) {
	return "", fmt.Errorf("not implemented")
}

func (r Registrar) AddPerformer(name string) (string, error) {
	identity, err := r.Repo.SaveAndIdentifyPerformer(name)
	if err != nil {
		return "", errors.Join(errAddPerfomer, err)
	}
	return identity, nil
}

func (r Registrar) AddPeriod(from, to string) (string, error) {
	return "", fmt.Errorf("%w: not implemented", errAddPeriod)
}
