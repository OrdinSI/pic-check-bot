package commands

import (
	"context"

	"github.com/OrdinSI/pic-check-bot/internal/model"

	"github.com/OrdinSI/pic-check-bot/internal/log"

	"github.com/OrdinSI/pic-check-bot/internal/usecase"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type CommandHandler struct {
	ucase usecase.Command
	log   *log.Logger
}

func NewCommandHandler(ucaese usecase.Command) *CommandHandler {
	return &CommandHandler{
		ucase: ucaese,
		log:   log.Named("CommandHandler"),
	}
}

func (h *CommandHandler) StartHandle(ctx context.Context, b *bot.Bot, update *models.Update) {
	req := model.RequestStart{
		GroupID:   update.Message.Chat.ID,
		GroupName: update.Message.Chat.Title,
	}

	message, err := h.ucase.StartCommand(ctx, req)
	if err != nil {
		h.log.Error("Error in StartCommand:", err)
		return
	}
	if message != "" {
		res, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   message,
		})
		if err != nil {
			h.log.Error("Error in SendMessage in Start:", err)
			return
		}
		h.log.Info("messages sent START", "chat_id", update.Message.Chat.ID, "message_id", res.ID)
		h.log.Info("RESPONSE START:", res)
	}
}
