package main

import (
	"context"
	"flag"
	"github.com/deyuro/zulip-combats/internal/config"
	"github.com/deyuro/zulip-combats/internal/zulip"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func app() error {

	var cfgFile string

	flag.StringVar(&cfgFile, "config", "config.yml", "config file path")
	flag.Parse()

	cfg, err := config.Load(cfgFile)
	if err != nil {
		return errors.WithMessage(err, "failed to load configuration")
	}

	appCtx, cancel := context.WithCancel(context.Background())

	bot := zulip.NewBot(cfg.Zulip.Bot.Email, cfg.Zulip.Bot.Key, cfg.Zulip.Entrypoint, &http.Client{})
	errChan := make(chan error)
	go func() {
		handleSIGINT()
		cancel()
	}()

	go func() {
		errChan <- run(cancel, bot)
	}()
	select {
	case err := <-errChan:
		return err
	case <-appCtx.Done():
		return nil
	}
}

func run(cancel context.CancelFunc, bot *zulip.Bot) error {

	streamList, err := bot.GetStreams()

	if err != nil {
		return err
	}

	_ = streamList

	return nil
}

func handleSIGUSR2(logger *logrus.Logger) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGUSR2)
	for range ch {
		level := logger.GetLevel()
		switch level {
		case logrus.DebugLevel:
			logger.Warn("switching log level to INFO")
			logger.SetLevel(logrus.InfoLevel)
		default:
			logger.Warn("switching log level to DEBUG")
			logger.SetLevel(logrus.DebugLevel)
		}
	}
}

func handleSIGINT() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	for range sigCh {
		signal.Stop(sigCh)
		return
	}
}
