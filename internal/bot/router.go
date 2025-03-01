package bot

import (
	"github.com/OrdinSI/pic-check-bot/internal/bot/handler/commands"
	"github.com/OrdinSI/pic-check-bot/internal/bot/handler/messages"
	"github.com/OrdinSI/pic-check-bot/internal/config"
	"github.com/OrdinSI/pic-check-bot/internal/usecase"
	"github.com/go-telegram/bot"
)

type Router struct {
	b   *bot.Bot
	ucC usecase.Command
	ucM usecase.Message
	cfg *config.Telegram
}

func NewRouters(b *bot.Bot, ucC usecase.Command, ucM usecase.Message, cfg *config.Telegram) *Router {
	return &Router{
		b:   b,
		ucC: ucC,
		ucM: ucM,
		cfg: cfg,
	}
}

func (r *Router) Handlers() {
	r.registerCommands()
}

func (r *Router) registerCommands() {
	commandHandler := commands.NewCommandHandler(r.ucC)
	messagesHandler := messages.NewMessageHandler(r.ucM, r.cfg)

	r.b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypePrefix, commandHandler.StartHandle)
	r.b.RegisterHandler(bot.HandlerTypeMessageText, "/help", bot.MatchTypePrefix, commandHandler.HelpHandle)
	r.b.RegisterHandler(bot.HandlerTypeMessageText, "/top", bot.MatchTypePrefix, commandHandler.TopHandle)

	r.b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypeContains, messagesHandler.RegisterMessageHandler)

}
