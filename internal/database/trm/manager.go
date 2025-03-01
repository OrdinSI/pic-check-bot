package trm

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

type ckey string

const (
	CKey = ckey("trm")
)

type Manager interface {
	Do(context.Context, func(ctx context.Context) error, ...*sql.TxOptions) error
}

type TxGormManager struct {
	db *gorm.DB
}

func NewTxManager(db *gorm.DB) Manager {
	return &TxGormManager{db: db}
}

type CtxGetter[db interface{}, tx any] interface {
	TrOrDB(ctx context.Context) db
	Transaction(ctx context.Context, f func(context.Context) error) error
}

func (t *TxGormManager) Do(ctx context.Context, f func(ctx context.Context) error, opts ...*sql.TxOptions) error {
	return t.db.Transaction(func(tx *gorm.DB) error {
		err := f(context.WithValue(ctx, CKey, tx))
		return err
	}, opts...)
}
