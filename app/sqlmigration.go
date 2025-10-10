package main

import (
	"context"
	"database/sql"
	"fmt"
)

func execSQLMigrationByScheme(scm, tbl string, ctx context.Context, db *sql.DB) (err error) {
	if db == nil {
		err = fmt.Errorf("not connection for server")
		return
	}
	tx, txBeginErr := db.Begin()
	if txBeginErr != nil {
		err = fmt.Errorf("init a migration transaction is failed: %v", txBeginErr)
		return
	}
	defer func() {
		if txErr := tx.Rollback(); txErr != nil {
			txErr = fmt.Errorf("failed for `%v` migration: %v", tbl, txErr)
			if err != nil {
				err = fmt.Errorf("%v, also %v", err, txErr)
			} else {
				err = txErr
			}
		}
	}()
	existsRow := tx.QueryRowContext(ctx, "select count(*) from information_schema.tables where table_type = 'BASE TABLE' and table_name = '$1';", tbl)
	var tableExists int
	if checkTableErr := existsRow.Scan(&tableExists); checkTableErr != nil {
		err = fmt.Errorf("check table existence is failed: %v", checkTableErr)
		return
	}
	if tableExists == 0 {
		_, migrateErr := tx.ExecContext(ctx, scm)
		if migrateErr != nil {
			err = fmt.Errorf("create schema for table is failed: %v", migrateErr)
			return
		}
	}
	if txCommitErr := tx.Commit(); txCommitErr != nil {
		err = fmt.Errorf("commit a migration transaction is failed: %v", txCommitErr)
		return
	}
	return nil
}
