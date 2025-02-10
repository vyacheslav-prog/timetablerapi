package main

type overviewRepo struct {
}

func (r *overviewRepo) fetchPerformerBoard(id string) *int {
	return nil
}

func newOverviewRepo() *overviewRepo {
	return nil
}

type overviewService struct {
	repo overviewRepo
}
