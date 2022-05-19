package sqlabst

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type executor struct {
	*sqlx.DB
}

type executorTx struct {
	executor
	*sqlx.Tx
}

type SqlAbst struct {
	executorTx
}

func NewSqlAbst(db *sqlx.DB) *SqlAbst {
	return &SqlAbst{executorTx: executorTx{
		executor: executor{
			DB: db,
		},
	}}
}

func (s SqlAbst) GetDB() *sqlx.DB {
	return s.executor.DB
}

func (s SqlAbst) Get(dest interface{}, query string, args ...interface{}) error {
	if s.Tx != nil {
		return s.Tx.Get(dest, query, args)
	}
	return s.DB.Get(dest, query, args)
}

func (s SqlAbst) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	if s.Tx != nil {
		return s.Tx.GetContext(ctx, dest, query, args)
	}
	return s.DB.GetContext(ctx, dest, query, args)
}

func (s SqlAbst) MustExec(query string, args ...interface{}) sql.Result {
	if s.Tx != nil {
		return s.Tx.MustExec(query, args)
	}
	return s.DB.MustExec(query, args)
}

func (s SqlAbst) MustExecContext(ctx context.Context, query string, args ...interface{}) sql.Result {
	if s.Tx != nil {
		return s.Tx.MustExecContext(ctx, query, args)
	}
	return s.DB.MustExecContext(ctx, query, args)
}

func (s SqlAbst) NamedExec(query string, arg interface{}) (sql.Result, error) {
	if s.Tx != nil {
		return s.Tx.NamedExec(query, arg)
	}
	return s.DB.NamedExec(query, arg)
}

func (s SqlAbst) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	if s.Tx != nil {
		return s.Tx.NamedExecContext(ctx, query, arg)
	}
	return s.DB.NamedExecContext(ctx, query, arg)
}

func (s SqlAbst) NamedQuery(query string, arg interface{}) (*sqlx.Rows, error) {
	if s.Tx != nil {
		return s.Tx.NamedQuery(query, arg)
	}
	return s.DB.NamedQuery(query, arg)
}

func (s SqlAbst) PrepareNamed(query string) (*sqlx.NamedStmt, error) {
	if s.Tx != nil {
		return s.Tx.PrepareNamed(query)
	}
	return s.DB.PrepareNamed(query)
}

func (s SqlAbst) PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error) {
	if s.Tx != nil {
		return s.Tx.PrepareNamedContext(ctx, query)
	}
	return s.DB.PrepareNamedContext(ctx, query)
}

func (s SqlAbst) Preparex(query string) (*sqlx.Stmt, error) {
	if s.Tx != nil {
		return s.Tx.Preparex(query)
	}
	return s.DB.Preparex(query)
}

func (s SqlAbst) PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error) {
	if s.Tx != nil {
		return s.Tx.PreparexContext(ctx, query)
	}
	return s.DB.PreparexContext(ctx, query)
}

func (s SqlAbst) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if s.Tx != nil {
		return s.Tx.Query(query, args...)
	}
	return s.DB.Query(query, args...)
}

func (s SqlAbst) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	if s.Tx != nil {
		return s.Tx.Query(query, args...)
	}
	return s.DB.Query(query, args...)
}

func (s SqlAbst) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	if s.Tx != nil {
		return s.Tx.QueryRowx(query, args...)
	}
	return s.DB.QueryRowx(query, args...)
}

func (s SqlAbst) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	if s.Tx != nil {
		return s.Tx.QueryRowxContext(ctx, query, args...)
	}
	return s.DB.QueryRowxContext(ctx, query, args...)
}

func (s SqlAbst) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	if s.Tx != nil {
		return s.Tx.Queryx(query, args...)
	}
	return s.DB.Queryx(query, args...)
}

func (s SqlAbst) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	if s.Tx != nil {
		return s.Tx.QueryxContext(ctx, query, args...)
	}
	return s.DB.QueryxContext(ctx, query, args...)
}

func (s SqlAbst) Rebind(query string) string {
	if s.Tx != nil {
		return s.Tx.Rebind(query)
	}
	return s.DB.Rebind(query)
}

func (s SqlAbst) Exec(query string, args ...interface{}) (sql.Result, error) {
	if s.Tx != nil {
		return s.Tx.Exec(query, args...)
	}
	return s.DB.Exec(query, args...)
}

func (s SqlAbst) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
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

func (s SqlAbst) QueryRow(query string, args ...interface{}) *sql.Row {
	if s.Tx != nil {
		return s.Tx.QueryRow(query, args...)
	}
	return s.DB.QueryRow(query, args...)
}

func (s SqlAbst) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	if s.Tx != nil {
		return s.Tx.QueryRowContext(ctx, query, args...)
	}
	return s.DB.QueryRowContext(ctx, query, args...)
}

func (s SqlAbst) Select(dest interface{}, query string, args ...interface{}) error {
	// If the transaction is not null, start from the Tx
	if s.Tx != nil {
		return s.Tx.Select(dest, query, args...)
	}
	// start from DB
	return s.DB.Select(dest, query, args...)
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
