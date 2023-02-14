package dbq

import (
	"context"
	"database/sql"
	"errors"
	"log"
)

type txKeyType struct{}

// DefaultTxOpts is package variable with default transaction level
var DefaultTxOpts = sql.TxOptions{
	Isolation: sql.LevelDefault,
	ReadOnly:  false,
}

// TxContext interface for DAO operations with context.
type TxContext interface {
	context.Context
	WithValue(key, value any) TxContext
	Prepare(query string) (*sql.Stmt, error)
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
}

// Tx represents transaction with context as inner object.
type Tx struct {
	context.Context //nolint:containedctx
	Tx              *sql.Tx
}

func (t *Tx) WithValue(key, value any) TxContext {
	return &Tx{
		Context: context.WithValue(t.Context, key, value),
		Tx:      t.Tx,
	}
}

// Prepare query.
func (t *Tx) Prepare(query string) (*sql.Stmt, error) {
	return t.Tx.PrepareContext(t.Context, query)
}

// Exec executes query with args.
func (t *Tx) Exec(query string, args ...any) (sql.Result, error) {
	return t.Tx.ExecContext(t.Context, query, args...)
}

// Query loads data from db.
func (t *Tx) Query(query string, args ...any) (*sql.Rows, error) {
	return t.Tx.QueryContext(t.Context, query, args...)
}

// QueryRow loads single row from db.
func (t *Tx) QueryRow(query string, args ...any) *sql.Row {
	return t.Tx.QueryRowContext(t.Context, query, args...)
}

// Commit this transaction.
func (t *Tx) Commit() error {
	return t.Tx.Commit()
}

// Rollback cancel this transaction.
func (t *Tx) Rollback() error {
	return t.Tx.Rollback()
}

// Connector for sql database.
type Connector interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

// TxProvider ...
type TxProvider struct {
	conn Connector
}

// NewTxProvider ...
func NewTxProvider(conn Connector) *TxProvider {
	return &TxProvider{
		conn: conn,
	}
}

// AcquireWithOpts transaction from db
func (t *TxProvider) AcquireWithOpts(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	tx, err := t.conn.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}

	return &Tx{
		Context: context.WithValue(ctx, txKeyType{}, Access(tx)),
		Tx:      tx,
	}, nil
}

// Acquire transaction from db
func (t *TxProvider) Acquire(ctx context.Context) (*Tx, error) {
	return t.AcquireWithOpts(ctx, &DefaultTxOpts)
}

// TxWithOpts ...
func (t *TxProvider) TxWithOpts(ctx context.Context, fn func(TxContext) error, opts *sql.TxOptions) error {
	tx, err := t.AcquireWithOpts(ctx, opts)
	if err != nil {
		return err
	}

	defer func() {
		//nolint:gocritic
		if r := recover(); r != nil {
			log.Printf("Recovering from panic in TxWithOpts error is: %v \n", r)
			_ = tx.Rollback()
			err, _ = r.(error)
		} else if err != nil {
			err = tx.Rollback()
		} else {
			err = tx.Commit()
		}

		if ctx.Err() != nil && errors.Is(err, context.DeadlineExceeded) {
			log.Printf("query response time exceeded the configured timeout")
		}
	}()

	err = fn(tx)

	return err
}

// Tx runs fn in transaction.
func (t *TxProvider) Tx(ctx context.Context, fn func(TxContext) error) error {
	return t.TxWithOpts(ctx, fn, &DefaultTxOpts)
}

// Access interface for simple DML operations.
type Access interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

// FromCtxOr returns access interface from context or data arg.
func FromCtxOr(ctx context.Context, data Access) Access {
	value, ok := ctx.Value(txKeyType{}).(Access)
	if ok {
		return value
	}
	return data
}
