package commands

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log/slog"
)

func CommandHelp(ctx context.Context, b *bot.Bot, update *models.Update) {
	res, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "help",
	})
	if err != nil {
		return
	}
	slog.Info("message sent", "chat_id", res.Chat.ID, "message_id", res.ID)
}
