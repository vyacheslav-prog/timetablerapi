package main

type Storage struct {
}

func (s *Storage) getDashboardTitle() string {
	return "Hello, world! By dashboard"
}

func newStorage() *Storage {
	return &Storage{}
}
