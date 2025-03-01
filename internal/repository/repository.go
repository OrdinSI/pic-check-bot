package repository

import (
	"context"

	"github.com/OrdinSI/pic-check-bot/internal/model"
)

type User interface {
	GetUser(ctx context.Context, id int64) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (int64, error)
	UpdateUser(ctx context.Context, id int64, user model.User) error

	GetGroup(ctx context.Context, id int64) (*model.Group, error)
	CreateGroup(ctx context.Context, group *model.Group) (int64, error)
	UpdateGroup(ctx context.Context, id int64, group model.Group) error
}

type Message interface {
	GetImagesByHashParts(ctx context.Context, images *[]model.Image, hashPart1, hashPart2 uint64) error
	CreateImage(ctx context.Context, image *model.Image) error
	CreateRepost(ctx context.Context, repost *model.Repost) error
}
