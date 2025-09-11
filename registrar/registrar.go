package registrar

type repository interface {
	SaveAndIdentifyPerformer(string) (string, error)
}

type Registrar struct {
	Repo repository
}

func (r Registrar) AddPerformer(name string) string {
	identity, _ := r.Repo.SaveAndIdentifyPerformer(name)
	return identity
}
