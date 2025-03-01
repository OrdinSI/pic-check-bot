package usecase

import (
	"context"

	"github.com/OrdinSI/pic-check-bot/internal/model"
)

const (
	StartMessage       = "Ну привет, привет."
	HelpMessage        = "Тебе уже ничего не поможет!"
	DuplicateThreshold = 10
)

type Command interface {
	StartCommand(ctx context.Context, req model.RequestStart) (string, error)
	HelpCommand(ctx context.Context) (string, error)
	TopCommand(ctx context.Context) ([]*model.TopRepost, error)
}

type Message interface {
	CheckHashImage(ctx context.Context, fileData []byte, req model.ImageRequest) (*model.Image, string, error)
	CheckUser(ctx context.Context, req model.UserRequest) error
}
