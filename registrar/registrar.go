package registrar

type repository interface {
	SaveAndIdentifyPerformer(string) (string, error)
}

type Registrar struct {
	repo repository
}

func (r Registrar) AddPerformer(name string) string {
	identity, _ := r.repo.SaveAndIdentifyPerformer(name)
	return identity
}
