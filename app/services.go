package main

type services struct {
	overview overviewService
}

func newServices() *services {
	return &services{}
}
