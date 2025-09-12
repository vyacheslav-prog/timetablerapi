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
	overview  overviewService
	registrar registrarService
}

func newServices() (*services, error) {
	dbConn, dbMode := os.Getenv("DATABASE_URL"), os.Getenv("DATABASE_MODE")
	db, openErr := sql.Open(dbMode, dbConn)
	if openErr != nil {
		return nil, openErr
	}
	if pingErr := db.Ping(); pingErr != nil {
		return nil, pingErr
	}
	return &services{
		overview.Overview{
			overviewRepo{db},
		},
		registrar.Registrar{
			registrarRepo{db},
		},
	}, nil
}
