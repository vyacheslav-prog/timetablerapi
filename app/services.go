package main

import (
	"context"
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	"timetablerapi/overview"
	"timetablerapi/registrar"
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

func newServices(ctx context.Context) (*services, error) {
	dbConn, dbMode := os.Getenv("DATABASE_URL"), os.Getenv("DATABASE_MODE")
	db, openErr := sql.Open(dbMode, dbConn)
	if openErr != nil {
		return nil, openErr
	}
	if pingErr := db.PingContext(ctx); pingErr != nil {
		return nil, pingErr
	}
	or, orErr := newOverviewRepo(ctx, db)
	if orErr != nil {
		return nil, orErr
	}
	rr, rrErr := newRegistrarRepo(ctx, db)
	if rrErr != nil {
		return nil, rrErr
	}
	return &services{
		overview.Overview{Repo: or},
		registrar.Registrar{Repo: rr},
	}, nil
}
