package services

import (
	"context"
	"database/sql"
	"errors"
)

const (
	pgCountTableByNameQuery      = "select count(*) from information_schema.tables where table_type = 'BASE TABLE' and table_name = '$1';"
	sqlite3CountTableByNameQuery = "select count(*) from sqlite_master where tbl_name = '$1';"
)

type dbMigrate struct {
	db              *sql.DB
	countTableQuery string
}

var (
	errMigrationCheckTable          = errors.New("check table existence is failed")
	errMigrationCreateScheme        = errors.New("create schema for table is failed")
	errMigrationNotConnection       = errors.New("not connection for server")
	errMigrationTransactionIsFailed = errors.New("init a migration transaction is failed")
)

func newDBMigrate(db *sql.DB, mode string) *dbMigrate {
	if mode == "sqlite3" {
		return &dbMigrate{db, sqlite3CountTableByNameQuery}
	}
	return &dbMigrate{db, pgCountTableByNameQuery}
}

func (dm *dbMigrate) byScheme(ctx context.Context, scm, tbl string) (err error) {
	if dm.db == nil {
		err = errMigrationNotConnection
		return
	}
	tx, txBeginErr := dm.db.BeginTx(ctx, nil)
	if txBeginErr != nil {
		err = errors.Join(errMigrationTransactionIsFailed, txBeginErr)
		return
	}
	defer func() {
		if err != nil {
			if txErr := tx.Rollback(); txErr != nil {
				err = errors.Join(err, errMigrationTransactionIsFailed, txErr)
			}
		}
	}()
	if migrateErr := dm.migrateIfNotExists(ctx, scm, tbl, tx); migrateErr != nil {
		err = migrateErr
		return
	}
	if txCommitErr := tx.Commit(); txCommitErr != nil {
		err = errors.Join(errMigrationTransactionIsFailed, txCommitErr)
		return
	}
	return nil
}

func (dm *dbMigrate) migrateIfNotExists(ctx context.Context, scm, tbl string, tx *sql.Tx) error {
	existsRow := tx.QueryRowContext(ctx, dm.countTableQuery, tbl)
	var tableExists int
	if checkTableErr := existsRow.Scan(&tableExists); checkTableErr != nil {
		return errors.Join(errMigrationCheckTable, checkTableErr)
	}
	if tableExists == 0 {
		_, migrateErr := tx.ExecContext(ctx, scm)
		if migrateErr != nil {
			return errors.Join(errMigrationCreateScheme, migrateErr)
		}
	}
	return nil
}
