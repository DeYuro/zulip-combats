package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/deyuro/zulip-combats/internal/config"
	"github.com/deyuro/zulip-combats/internal/logger"
	"github.com/deyuro/zulip-combats/internal/service"
	"github.com/deyuro/zulip-combats/internal/zulip"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
)

func app() error {
	container := buildContainer()

	var bot *zulip.Bot
	var logger *logrus.Logger
	err := container.Invoke(func(l *logrus.Logger) {
		logger = l
	})

	if err != nil {
		panic(err)
	}

	appCtx, cancel := context.WithCancel(context.Background())

	errChan := make(chan error)
	go func() {
		handleSIGINT()
		cancel()
	}()

	go func() {
		errChan <- run(cancel, bot, logger)
	}()
	select {
	case err := <-errChan:
		return err
	case <-appCtx.Done():
		return nil
	}
}

func run(cancel context.CancelFunc, bot *zulip.Bot, logger logrus.FieldLogger) error {

	q, err := bot.RegisterEventQueuePrivate()
	if err != nil {
		return err
	}
	bot.SetQueue(q)

	src := service.NewService(bot, logger)

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

func buildContainer() *dig.Container {
	var cfgFile string

	flag.StringVar(&cfgFile, "config", "config.yml", "config file path")
	flag.Parse()

	c := dig.New()
	err := c.Provide(logger.NewLogger)
	if err != nil {
		panic(err)
	}
	err = c.Provide(func(path string) (*config.Config, error) {
		return config.NewConfig(path)
	})

	if err != nil {
		panic(err)
	}
	err = c.Provide(zulip.NewBot)

	return c
}
