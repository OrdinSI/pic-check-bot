package bot

import (
	"github.com/OrdinSI/pic-check-bot/internal/bot/handler/commands"
	"github.com/go-telegram/bot"
)

type Router struct {
	b *bot.Bot
}

func NewRouters(b *bot.Bot) *Router {
	return &Router{
		b: b,
	}
}

func (r *Router) Handlers() {
	r.registerCommands()

}

func (r *Router) registerCommands() {
	r.b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, commands.CommandStart)
	r.b.RegisterHandler(bot.HandlerTypeMessageText, "/help", bot.MatchTypeExact, commands.CommandHelp)
}
