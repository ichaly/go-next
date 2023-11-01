package sharding

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

type Connection struct {
	gorm.ConnPool
	// db, This is global db instance
	sharding *Sharding
}

func (my *Connection) String() string {
	return "gorm:sharding:conn_pool"
}

func (my *Connection) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	ftQuery, stQuery, table, err := my.sharding.resolve(query, args...)
	if err != nil {
		return nil, err
	}

	if table != "" {
		if r, ok := my.sharding.configs[table]; ok {
			if r.DoubleWrite {
				_, _ = my.ConnPool.ExecContext(ctx, ftQuery, args...)
			}
		}
	}

	return my.ConnPool.ExecContext(ctx, stQuery, args...)
}

func (my *Connection) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	ftQuery, stQuery, table, err := my.sharding.resolve(query, args...)
	if err != nil {
		return nil, err
	}
	if table != "" {
		if r, ok := my.sharding.configs[table]; ok {
			if r.DoubleWrite {
				_, _ = my.ConnPool.ExecContext(ctx, ftQuery, args...)
			}
		}
	}
	return my.ConnPool.QueryContext(ctx, stQuery, args...)
}

func (my *Connection) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	_, query, _, _ = my.sharding.resolve(query, args...)

	return my.ConnPool.QueryRowContext(ctx, query, args...)
}

func (my *Connection) BeginTx(ctx context.Context, opt *sql.TxOptions) (gorm.ConnPool, error) {
	if basePool, ok := my.ConnPool.(gorm.ConnPoolBeginner); ok {
		return basePool.BeginTx(ctx, opt)
	}
	return my, nil
}

func (my *Connection) Commit() error {
	if _, ok := my.ConnPool.(*sql.Tx); ok {
		return nil
	}

	if basePool, ok := my.ConnPool.(gorm.TxCommitter); ok {
		return basePool.Commit()
	}

	return nil
}

func (my *Connection) Rollback() error {
	if _, ok := my.ConnPool.(*sql.Tx); ok {
		return nil
	}

	if basePool, ok := my.ConnPool.(gorm.TxCommitter); ok {
		return basePool.Rollback()
	}

	return nil
}
