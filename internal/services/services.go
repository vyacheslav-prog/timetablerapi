package services

import (
	"context"
	"database/sql"
	"errors"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"timetablerapi/overview"
	"timetablerapi/registrar"
)

type OverviewService interface {
	ViewPerformerBoard(context.Context, string) (string, error)
}

type RegistrarService interface {
	AddPerformer(string) (string, error)
	AddPeriod(string, string) (string, error)
}

type Services struct {
	Overview  OverviewService
	Registrar RegistrarService
}

var (
	errInitServices = errors.New("failed init services")
)

func NewServices(ctx context.Context) (*Services, error) {
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
	return &Services{
		overview.Overview{Repo: or},
		registrar.Registrar{Repo: rr},
	}, nil
}
