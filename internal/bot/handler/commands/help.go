package commands

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (h *CommandHandler) HelpHandle(ctx context.Context, b *bot.Bot, update *models.Update) {
	message, err := h.ucase.HelpCommand(ctx)
	if err != nil {
		return
	}
	if message != "" {
		res, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   message,
		})
		if err != nil {
			return
		}
		h.log.Info("messages sent Help", "chat_id", update.Message.Chat.ID, "message_id", res.ID)
	}
}
