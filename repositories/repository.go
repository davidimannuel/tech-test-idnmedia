package repositories

import (
	"context"
	"database/sql"
	"time"
)

type BaseModel struct {
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}

type DBTransactional interface {
	Begin(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

type dbTx struct {
	db *sql.DB
}

func NewDBTx(db *sql.DB) DBTransactional {
	return &dbTx{
		db: db,
	}
}

func (tx *dbTx) Begin(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return tx.db.BeginTx(ctx, opts)
}
