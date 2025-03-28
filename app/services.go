package main

import "timetablerapi/overview"

type overviewService interface {
	ViewPerformerBoard(string) string
}

type services struct {
	overview overviewService
}

func newServices() *services {
	return &services{
		overview.Overview{},
	}
}
