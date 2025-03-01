package repository

import (
	"context"

	"github.com/OrdinSI/pic-check-bot/internal/database/txgorm"
	"github.com/OrdinSI/pic-check-bot/internal/model"
	"gorm.io/gorm"
)

type message struct {
	*txgorm.TxGetter
}

func NewMessageRepository(db *gorm.DB) Message {
	return &message{
		TxGetter: txgorm.NewTxGetter(db),
	}
}

func (r *message) GetImagesByHashParts(ctx context.Context, images *[]model.Image, hashPart1, hashPart2 uint64) error {
	return r.TrOrDB(ctx).WithContext(ctx).
		Where("hash_part1 = ? AND hash_part2 = ?", hashPart1, hashPart2).Find(images).Error
}

func (r *message) CreateImage(ctx context.Context, image *model.Image) error {
	return r.TrOrDB(ctx).WithContext(ctx).Create(image).Error
}

func (r *message) CreateRepost(ctx context.Context, repost *model.Repost) error {
	return r.TrOrDB(ctx).WithContext(ctx).Create(repost).Error
}
