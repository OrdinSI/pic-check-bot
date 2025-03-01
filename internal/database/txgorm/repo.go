package txgorm

import (
	"context"

	"github.com/OrdinSI/pic-check-bot/internal/database/trm"

	"gorm.io/gorm"
)

type TxGetter struct {
	db *gorm.DB
}

func NewTxGetter(db *gorm.DB) *TxGetter {
	return &TxGetter{db: db}
}

func (c *TxGetter) TrOrDB(ctx context.Context) *gorm.DB {
	if tr := ctx.Value(trm.CKey); tr != nil {
		return convert(tr)
	}
	return c.db
}

func (c *TxGetter) Transaction(ctx context.Context, f func(context.Context) error) error {
	return trm.NewTxManager(c.TrOrDB(ctx)).Do(ctx, f)
}

func convert(tr any) *gorm.DB {
	if tx, ok := tr.(*gorm.DB); ok {
		return tx
	}
	return nil
}
