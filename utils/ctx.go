package utils

import (
	"context"
	"database/sql"
)

type ctxKey int

const (
	CtxKeyPlayer ctxKey = iota
	CtxKeyDBTx   ctxKey = iota
)

type CtxPLayer struct {
	Id    int
	Email string
}

func SetCtxPLayer(ctx context.Context, ctxPlayer *CtxPLayer) context.Context {
	return context.WithValue(ctx, CtxKeyPlayer, ctxPlayer)
}

func GetCtxPLayer(ctx context.Context) (*CtxPLayer, error) {
	data, ok := ctx.Value(CtxKeyPlayer).(*CtxPLayer)
	if !ok {
		return nil, ErrInvalidPlayerSession
	}
	return data, nil
}

func SetCtxDBTx(ctx context.Context, ctxDBTx *sql.Tx) context.Context {
	return context.WithValue(ctx, CtxKeyDBTx, ctxDBTx)
}

func GetCtxDBTx(ctx context.Context) *sql.Tx {
	tx, ok := ctx.Value(CtxKeyDBTx).(*sql.Tx)
	if !ok {
		return nil
	}
	return tx
}
