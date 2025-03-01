package messages

import (
	"context"

	"github.com/OrdinSI/pic-check-bot/internal/log"
	"github.com/OrdinSI/pic-check-bot/internal/model"
	"github.com/OrdinSI/pic-check-bot/internal/repository"
	"github.com/OrdinSI/pic-check-bot/internal/usecase"
)

type messages struct {
	repoU repository.User
	repoM repository.Message
	log   *log.Logger
}

func NewUsecase(rU repository.User, rM repository.Message) usecase.Message {
	return &messages{
		repoU: rU,
		repoM: rM,
		log:   log.Named("messages_usecase"),
	}
}

func (m *messages) CheckUser(ctx context.Context, req model.UserRequest) error {
	user, err := m.repoU.GetUser(ctx, req.UserID)
	if err != nil {
		return err
	}

	if user == nil {
		if _, err := m.repoU.CreateUser(ctx, &model.User{
			ID:        req.UserID,
			Username:  req.Username,
			FirstName: req.FirstName,
			LastName:  req.LastName,
		}); err != nil {
			return err
		}
	}
	return nil
}
