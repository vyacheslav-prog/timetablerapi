package main

import (
	"database/sql"
	"os"

	"timetablerapi/overview"
)

type overviewService interface {
	ViewPerformerBoard(string) string
}

type services struct {
	overview overviewService
}

func newServices() *services {
	dbConn, dbMode := os.Getenv("DATABASE_URL"), os.Getenv("DATABASE_MODE")
	db, err := sql.Open(dbMode, dbConn)
	return &services{
		overview.Overview{
			overviewRepo{db},
		},
	}
}
