package registrar

import "errors"

type repository interface {
	SaveAndIdentifyPerformer(string) (string, error)
}

type Registrar struct {
	Repo repository
}

var (
	errAddPerfomer = errors.New("unable to add a performer")
)

func (r Registrar) AddPerformer(name string) (string, error) {
	identity, err := r.Repo.SaveAndIdentifyPerformer(name)
	if err != nil {
		return "", errors.Join(errAddPerfomer, err)
	}
	return identity, nil
}
