package main

type services struct {
	overview overviewService
	storage storage
}

type storage struct {
}

func (s *storage) getDashboardTitle() string {
	return "Hello, world! By dashboard"
}

func newServices() *services {
	return &services{}
}
