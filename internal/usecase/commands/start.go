package commands

import (
	"context"

	"github.com/OrdinSI/pic-check-bot/internal/model"
	"github.com/OrdinSI/pic-check-bot/internal/repository"
	"github.com/OrdinSI/pic-check-bot/internal/usecase"
)

type command struct {
	repo repository.User
}

func NewUsecase(r repository.User) usecase.Command {
	return &command{
		repo: r,
	}
}

func (s *command) StartCommand(ctx context.Context, req model.RequestStart) (string, error) {
	group, err := s.repo.GetGroup(ctx, req.GroupID)
	if err != nil {
		return "", err
	}

	if group == nil {
		if _, err := s.repo.CreateGroup(ctx, &model.Group{
			ID:        req.GroupID,
			GroupName: req.GroupName,
		}); err != nil {
			return "", err
		}
	}
	return usecase.StartMessage, nil
}
