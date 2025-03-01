package messages

import (
	"context"
	"strings"

	"github.com/OrdinSI/pic-check-bot/internal/config"
	"github.com/OrdinSI/pic-check-bot/internal/model"

	"github.com/OrdinSI/pic-check-bot/internal/log"
	"github.com/OrdinSI/pic-check-bot/internal/usecase"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type MessageHandler struct {
	ucase usecase.Message
	cfg   *config.Telegram
	log   *log.Logger
}

func NewMessageHandler(ucaese usecase.Message, cfg *config.Telegram) *MessageHandler {
	return &MessageHandler{
		ucase: ucaese,
		cfg:   cfg,
		log:   log.Named("MessageHandler"),
	}
}

func (h *MessageHandler) RegisterMessageHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil || update.Message.From == nil {
		h.log.Info("update_message in TextHandle is nil")
		return
	}

	if update.Message.Text != "" {
		h.textHandle(ctx, b, update)
		return
	}

	if len(update.Message.Photo) > 0 {
		h.imageHandle(ctx, b, update)
		return
	}
}

func (h *MessageHandler) textHandle(ctx context.Context, b *bot.Bot, update *models.Update) {
	if !strings.HasPrefix(update.Message.Text, "#") {
		req := model.UserRequest{
			UserID:    update.Message.From.ID,
			GroupID:   update.Message.Chat.ID,
			Username:  update.Message.From.Username,
			FirstName: update.Message.From.FirstName,
			LastName:  update.Message.From.LastName,
		}
		if err := h.ucase.CheckUser(ctx, req); err != nil {
			h.log.Error("Error in CheckUser in TextHandle:", err)
			return
		}
		return
	}
	res, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Ну и чушь",
	})
	if err != nil {
		return
	}
	h.log.Info("messages sent TEXT", "chat_id", update.Message.Chat.ID, "message_id", res.ID)
}
