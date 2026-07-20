package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type dbMigrate struct {
	db              *sql.DB
	countTableQuery string
}

var (
	errMigrationCheckTable          = errors.New("check table existence is failed")
	errMigrationCreateScheme        = errors.New("create schema for table is failed")
	errMigrationNotConnection       = errors.New("not connection to sql-server")
	errMigrationTransactionIsFailed = errors.New("init a migration transaction is failed")
)

func (dm *dbMigrate) byScheme(ctx context.Context, scm, tbl string) (err error) {
	if dm.db == nil {
		err = errMigrationNotConnection
		return
	}
	tx, txBeginErr := dm.db.BeginTx(ctx, nil)
	if txBeginErr != nil {
		err = fmt.Errorf("%w: %w", errMigrationTransactionIsFailed, txBeginErr)
		return
	}
	defer func() {
		if err != nil {
			if txErr := tx.Rollback(); txErr != nil {
				err = fmt.Errorf("%w; %w", err, txErr)
			}
		}
	}()
	if migrateErr := dm.migrateIfNotExists(ctx, scm, tbl, tx); migrateErr != nil {
		err = fmt.Errorf("%w: %w", errMigrationTransactionIsFailed, migrateErr)
		return
	}
	if txCommitErr := tx.Commit(); txCommitErr != nil {
		err = fmt.Errorf("%w: %w", errMigrationTransactionIsFailed, txCommitErr)
		return
	}
	return nil
}

func (dm *dbMigrate) migrateIfNotExists(ctx context.Context, scm, tbl string, tx *sql.Tx) error {
	existsRow := tx.QueryRowContext(ctx, dm.countTableQuery, tbl)
	var tableExists int
	if checkTableErr := existsRow.Scan(&tableExists); checkTableErr != nil {
		return fmt.Errorf("%w: %w", errMigrationCheckTable, checkTableErr)
	}
	if tableExists == 0 {
		_, migrateErr := tx.ExecContext(ctx, scm)
		if migrateErr != nil {
			return fmt.Errorf("%w: %w", errMigrationCreateScheme, migrateErr)
		}
	}
	return nil
}
