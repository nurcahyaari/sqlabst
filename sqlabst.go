package sqlabst

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type Executor struct {
	*sqlx.DB
}

type ExecutorTx struct {
	Executor
	*sqlx.Tx
}

type SqlAbst struct {
	ExecutorTx
}

func NewSqlAbst(db *sqlx.DB) *SqlAbst {
	return &SqlAbst{ExecutorTx: ExecutorTx{
		Executor: Executor{
			DB: db,
		},
	}}
}

func (s SqlAbst) GetDB() *sqlx.DB {
	return s.Executor.DB
}

func (s SqlAbst) Query(query string, args ...any) (*sql.Rows, error) {
	if s.Tx != nil {
		return s.Tx.Query(query, args...)
	}
	return s.DB.Query(query, args...)
}

func (s SqlAbst) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	if s.Tx != nil {
		return s.Tx.Query(query, args...)
	}
	return s.DB.Query(query, args...)
}

func (s SqlAbst) Exec(query string, args ...any) (sql.Result, error) {
	if s.Tx != nil {
		return s.Tx.Exec(query, args...)
	}
	return s.DB.Exec(query, args...)
}

func (s SqlAbst) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	if s.Tx != nil {
		return s.Tx.ExecContext(ctx, query, args...)
	}
	return s.DB.ExecContext(ctx, query, args...)
}

func (s SqlAbst) Prepare(query string) (*sql.Stmt, error) {
	if s.Tx != nil {
		return s.Tx.Prepare(query)
	}
	return s.DB.Prepare(query)
}

func (s SqlAbst) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	if s.Tx != nil {
		return s.Tx.PrepareContext(ctx, query)
	}
	return s.DB.PrepareContext(ctx, query)
}

func (s SqlAbst) QueryRow(query string, args ...any) *sql.Row {
	if s.Tx != nil {
		return s.Tx.QueryRow(query, args...)
	}
	return s.DB.QueryRow(query, args...)
}

func (s SqlAbst) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	if s.Tx != nil {
		return s.Tx.QueryRowContext(ctx, query, args...)
	}
	return s.DB.QueryRowContext(ctx, query, args...)
}

func (s SqlAbst) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	// If the transaction is not null, start from the Tx
	if s.Tx != nil {
		return s.Tx.SelectContext(ctx, dest, query, args...)
	}
	// start from DB
	return s.DB.SelectContext(ctx, dest, query, args...)
}

// Begin begins a transaction and set the Tx object
//
// This abstraction resulting sqlx.Tx cause the implementation using sqlx.Tx instead of sql.Tx
func (s *SqlAbst) Begin() error {
	// begin from the DB.Beginx
	tx, err := s.DB.Beginx()
	// assign the value to the Tx
	s.Tx = tx
	return err
}

// BeginTx begins a transaction and set the Tx object
//
// This abstraction resulting sqlx.Tx cause the implementation using sqlx.Tx instead of sql.Tx
func (s *SqlAbst) BeginTx(ctx context.Context, opts *sql.TxOptions) error {
	// begin from the DB.Beginx
	tx, err := s.DB.BeginTxx(ctx, opts)
	// assign the value to the Tx
	s.Tx = tx
	return err
}

// Beginx begins a transaction and set the Tx object
func (s *SqlAbst) Beginx() error {
	// begin from the DB.Beginx
	tx, err := s.DB.Beginx()
	// assign the value to the Tx
	s.Tx = tx
	return err
}

// BeginTxx begins a transaction and set the Tx object
func (s *SqlAbst) BeginTxx(ctx context.Context, opts *sql.TxOptions) error {
	// begin from the DB.Beginx
	tx, err := s.DB.BeginTxx(ctx, opts)
	// assign the value to the Tx
	s.Tx = tx
	return err
}

// MustBegin begins a transaction and set the Tx object
func (s *SqlAbst) MustBegin() {
	// begin from the DB.Beginx
	tx := s.DB.MustBegin()
	// assign the value to the Tx
	s.Tx = tx
}

// MustBeginTx begins a transaction and set the Tx object
func (s *SqlAbst) MustBeginTx(ctx context.Context, opts *sql.TxOptions) {
	// begin from the DB.Beginx
	tx := s.DB.MustBeginTx(ctx, opts)
	// assign the value to the Tx
	s.Tx = tx
}

// Commit commits the transaction and set nil to Tx object
func (s *SqlAbst) Commit() error {
	// check the transaction has been started or not
	if s.Tx == nil {
		return errors.New("err: the transaction is not started")
	}
	err := s.Tx.Commit()
	// nulling the Tx
	s.Tx = nil
	return err
}

// Rollback aborts the transaction  and set nil to Tx object
func (s *SqlAbst) Rollback() error {
	// check the transaction has been started or not
	if s.Tx == nil {
		return errors.New("err: the transaction is not started")
	}
	err := s.Tx.Rollback()
	// nulling the Tx
	s.Tx = nil
	return err
}
