package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"timetablerapi/overview"
	"timetablerapi/registrar"
)

type OverviewService interface {
	ViewPerformerBoard(context.Context, string) (*overview.PerformerBoard, error)
}

type RegistrarService interface {
	AddPerformer(context.Context, registrar.Performer) (string, error)
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
		return nil, fmt.Errorf("%w: %w", errInitServices, openErr)
	}
	if pingErr := db.PingContext(ctx); pingErr != nil {
		return nil, fmt.Errorf("%w: %w", errInitServices, pingErr)
	}
	dm := newDBMigrate(db, dbMode)
	or, orErr := newOverviewRepo(ctx, db, dm)
	if orErr != nil {
		return nil, orErr
	}
	rr, rrErr := newRegistrarRepo(ctx, db, dm)
	if rrErr != nil {
		return nil, rrErr
	}
	return &Services{
		overview.Overview{Repo: or},
		registrar.Registrar{Repo: rr},
	}, nil
}
