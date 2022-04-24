package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/deyuro/zulip-combats/internal/config"
	"github.com/deyuro/zulip-combats/internal/service"
	"github.com/deyuro/zulip-combats/internal/zulip"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
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
	logger := setupLogger()
	bot := zulip.NewBot(cfg.Zulip.Bot.Email, cfg.Zulip.Bot.Key, cfg.Zulip.Entrypoint, &http.Client{}, logger)
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

	q, err := bot.RegisterEventQueuePrivate()
	if err != nil {
		return err
	}
	bot.SetQueue(q)

	src := service.NewService(bot)

	err = src.Run()

	if err != nil {
		return err
	}

	return nil
}
func setupLogger() *logrus.Logger {
	logger := logrus.New()
	logger.ReportCaller = false

	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors:    true,
		DisableTimestamp: false,
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05",
		QuoteEmptyFields: true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			file, line := frame.Func.FileLine(frame.PC)
			return frame.Function, fmt.Sprintf("%s:%d", filepath.Base(file), line)
		},
	})

	go handleSIGUSR2(logger)
	return logger
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
