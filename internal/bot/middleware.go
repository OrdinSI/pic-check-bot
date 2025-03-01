package bot

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

const (
	MyGroupID          = -1001270227250
	OnlyMyGroupMessage = "Этот бот работает только в нашей группе."
	OnlyGroupMessage   = "Этот бот работает только в группах."
	TypeGroup          = "group"
	TypeSupergroup     = "supergroup"
)

func GroupOnlyMiddleware(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message == nil {
			return
		}

		if update.Message.Chat.Type != TypeGroup && update.Message.Chat.Type != TypeSupergroup {
			if err := sendMessage(ctx, b, update.Message.Chat.ID, OnlyGroupMessage); err != nil {
				return
			}
		}
		if update.Message.Chat.ID != MyGroupID {
			if err := sendMessage(ctx, b, update.Message.Chat.ID, OnlyMyGroupMessage); err != nil {
				return
			}
		}

		next(ctx, b, update)
	}
}

func sendMessage(ctx context.Context, b *bot.Bot, chatID int64, text string) error {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   text,
	})
	if err != nil {
		return err
	}
	return nil
}
