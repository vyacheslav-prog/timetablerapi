package main

import (
	"context"
	"database/sql"
	"errors"
	"os"

	_ "github.com/lib/pq"
	"timetablerapi/overview"
	"timetablerapi/registrar"
)

type overviewService interface {
	ViewPerformerBoard(context.Context, string) (string, error)
}

type registrarService interface {
	AddPerformer(string) (string, error)
}

type services struct {
	overview  overviewService
	registrar registrarService
}

var (
	errInitServices = errors.New("failed init services")
)

func newServices(ctx context.Context) (*services, error) {
	dbConn, dbMode := os.Getenv("DATABASE_URL"), os.Getenv("DATABASE_MODE")
	db, openErr := sql.Open(dbMode, dbConn)
	if openErr != nil {
		return nil, errors.Join(errInitServices, openErr)
	}
	if pingErr := db.PingContext(ctx); pingErr != nil {
		return nil, errors.Join(errInitServices, pingErr)
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
