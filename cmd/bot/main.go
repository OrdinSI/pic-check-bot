package main

import (
	"context"
	"errors"
	mbot "github.com/OrdinSI/pic-check-bot/internal/bot"
	"github.com/OrdinSI/pic-check-bot/internal/config"
	"github.com/OrdinSI/pic-check-bot/internal/database"
	"github.com/OrdinSI/pic-check-bot/internal/log"
	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
	stdlog "log"
	"os"
	"os/signal"
	"runtime/debug"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		stdlog.Fatalf("failed to load .env file: %v", err)
	}
	cfg := config.New()

	if err := log.NewLogger(cfg.Dev); err != nil {
		stdlog.Fatalf("Не удалось создать логгер: %v", err)
	}
	logger := log.Named("main")

	defer func() {
		v := recover()
		if v != nil {
			if err, ok := v.(error); ok {
				logger.Error("panic: %v", errors.New(err.Error()+"\nstacktrace from panic:\n"+string(debug.Stack())))
			} else {
				panic(v)
			}
		}
	}()

	db, err := database.InitDB(&cfg.Database)
	if err != nil {
		logger.Fatal("failed to init database: %v", err)
	}
	if err := database.Migrate(db); err != nil {
		logger.Fatal("failed to migrate database: %v", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b, err := bot.New(cfg.Telegram.Token)
	if err != nil {
		panic(err)
	}

	router := mbot.NewRouters(b)
	router.Handlers()

	b.Start(ctx)

	if err := database.CloseDB(db); err != nil {
		logger.Error("failed to close database: %v", err)
	}

	if err := log.CloseLogger(); err != nil {
		stdlog.Fatalf("failed to close logger: %v", err)
	}

}
