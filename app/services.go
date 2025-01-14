package main

type services struct {
	storage storage
}

type storage struct {
}

type overviewRepo struct {
}

func (r *overviewRepo) fetchPerformerBoard() *int {
	return nil
}

func (s *storage) getDashboardTitle() string {
	return "Hello, world! By dashboard"
}

func newOverviewRepo() *overviewRepo {
	return nil
}

func newServices() *services {
	return &services{}
}
