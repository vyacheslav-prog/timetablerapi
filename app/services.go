package main

import (
	"database/sql"
	"os"

	"timetablerapi/overview"
	"timetablerapi/registrar"

	_ "github.com/lib/pq"
)

type overviewService interface {
	ViewPerformerBoard(string) string
}

type registrarService interface {
	AddPerformer(string) string
}

type services struct {
	overview    overviewService
	registrar   registrarService
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
		registrar.Registrar{},
	}, nil
}
