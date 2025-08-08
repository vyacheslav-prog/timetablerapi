package main

import (
	"database/sql"
	"os"

	"timetablerapi/overview"

	_ "github.com/lib/pq"
)

type overviewService interface {
	ViewPerformerBoard(string) string
}

type timetablingService struct {
}

func (ts *timetablingService) AddPerformer() string {
	return "{\"performer_id\": \"a-a-a-a\"}"
}

type services struct {
	overview overviewService
	timetabling timetablingService
}

func newServices() (*services, error) {
	dbConn, dbMode := os.Getenv("DATABASE_URL"), os.Getenv("DATABASE_MODE")
	db, err := sql.Open(dbMode, dbConn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &services{
		overview.Overview{
			overviewRepo{db},
		},
		timetablingService{},
	}, nil
}
