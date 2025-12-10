package sql

import (
	"context"
	"database/sql"
	"errors"
)

var (
	errMigrationCheckTable          = errors.New("check table existence is failed")
	errMigrationCreateScheme        = errors.New("create schema for table is failed")
	errMigrationNotConnection       = errors.New("not connection for server")
	errMigrationTransactionIsFailed = errors.New("init a migration transaction is failed")
)

func execSQLMigrationByScheme(scm, tbl string, ctx context.Context, db *sql.DB) (err error) {
	if db == nil {
		err = errMigrationNotConnection
		return
	}
	tx, txBeginErr := db.BeginTx(ctx, nil)
	if txBeginErr != nil {
		err = errors.Join(errMigrationTransactionIsFailed, txBeginErr)
		return
	}
	defer func() {
		if txErr := tx.Rollback(); txErr != nil {
			txErr = errors.Join(errMigrationTransactionIsFailed, txErr)
			if err != nil {
				err = errors.Join(err, txErr)
			} else {
				err = txErr
			}
		}
	}()
	if err = checkSQLTableExistenceOrCreate(scm, tbl, ctx, tx); err != nil {
		return
	}
	if txCommitErr := tx.Commit(); txCommitErr != nil {
		err = errors.Join(errMigrationTransactionIsFailed, txCommitErr)
		return
	}
	return nil
}

func checkSQLTableExistenceOrCreate(scm, tbl string, ctx context.Context, tx *sql.Tx) error {
	existsRow := tx.QueryRowContext(ctx, "select count(*) from information_schema.tables where table_type = 'BASE TABLE' and table_name = '$1';", tbl)
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
