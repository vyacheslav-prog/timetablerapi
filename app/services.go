package main

import "timetablerapi/overview"

type services struct {
	overview overview.OverviewService
}

func newServices() *services {
	return &services{
		overview.OverviewService{},
	}
}
