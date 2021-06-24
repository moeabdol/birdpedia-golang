package models

import (
	"context"
	"database/sql"
	"fmt"
)

// DBTX interface
type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

// Store struct will provide all functions to execute db queries and
// transactions
type Store struct {
	dbtx DBTX
	db   *sql.DB
}

// New function creates a new Store instance
func New(dbtx DBTX) *Store {
	return &Store{dbtx: dbtx}
}

// WithTx function
func (store *Store) WithTx(tx *sql.Tx) *Store {
	return &Store{
		dbtx: tx,
	}
}

// execTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Store) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	s := New(tx)
	err = fn(s)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
