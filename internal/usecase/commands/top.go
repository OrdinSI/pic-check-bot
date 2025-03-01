package commands

import (
	"context"

	"github.com/OrdinSI/pic-check-bot/internal/model"
)

func (s *command) TopCommand(ctx context.Context) ([]*model.TopRepost, error) {
	return s.repo.TopReposts(ctx)
}
