package commands

import (
	"context"
	"fmt"
	"strings"

	"github.com/OrdinSI/pic-check-bot/internal/model"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (h *CommandHandler) TopHandle(ctx context.Context, b *bot.Bot, update *models.Update) {

	topReposts, err := h.ucase.TopCommand(ctx)
	if err != nil {
		return
	}
	topMessage := formatTopReposts(topReposts)

	if topMessage != "" {
		res, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   topMessage,
		})
		if err != nil {
			return
		}
		h.log.Info("messages sent TOP", "chat_id", update.Message.Chat.ID, "message_id", res.ID)
	}
}

func formatTopReposts(topReposts []*model.TopRepost) string {
	if len(topReposts) == 0 {
		return "Нет данных для формирования топа репостов."
	}
	var builder strings.Builder
	builder.WriteString("🪗 Топ боянистов:\n\n")
	for i, topRepost := range topReposts {
		builder.WriteString(fmt.Sprintf(
			"%d. @%s - %d репостов\n",
			i+1, topRepost.Username, topRepost.Count,
		))
	}

	return builder.String()
}
