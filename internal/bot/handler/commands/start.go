package commands

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log/slog"
)

func CommandStart(ctx context.Context, b *bot.Bot, update *models.Update) {
	res, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Привет!",
	})
	if err != nil {
		return
	}
	slog.Info("message sent", "chat_id", update.Message.Chat.ID, "message_id", res.ID)
}
