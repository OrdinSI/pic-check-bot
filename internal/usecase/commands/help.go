package commands

import (
	"context"

	"github.com/OrdinSI/pic-check-bot/internal/usecase"
)

func (s *command) HelpCommand(ctx context.Context) (string, error) {
	return usecase.HelpMessage, nil
}
