package main

import (
	"context"
	"flag"
	"github.com/deyuro/zulip-combats/internal/config"
	"github.com/deyuro/zulip-combats/internal/logger"
	"github.com/deyuro/zulip-combats/internal/service"
	"github.com/deyuro/zulip-combats/internal/zulip"
	"go.uber.org/dig"
	"os"
	"os/signal"
	"syscall"
)

func app() error {
	appCtx, cancel := context.WithCancel(context.Background())
	go func() {
		handleSIGINT()
		cancel()
	}()

	c, err := buildContainer()
	if err != nil {
		return err
	}

	errChan := make(chan error)
	go func() {
		errChan <- run(c)
	}()
	select {
	case err := <-errChan:
		return err
	case <-appCtx.Done():
		return nil
	}
}

func run(container *dig.Container) error {

	err := container.Invoke(func(s *service.Service) error {
		return s.Run()
	})
	if err != nil {
		return err
	}

	return nil
}

func handleSIGINT() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	for range sigCh {
		signal.Stop(sigCh)
		return
	}
}

func buildContainer() (*dig.Container, error) {
	var cfgFile string

	flag.StringVar(&cfgFile, "config", "config.yml", "config file path")
	flag.Parse()

	c := dig.New()
	err := c.Provide(logger.NewLogger)
	if err != nil {
		return nil, err
	}
	err = c.Provide(func(path string) (*config.Config, error) {
		return config.NewConfig(path)
	})
	if err != nil {
		return nil, err
	}
	err = c.Provide(zulip.NewBot)
	if err != nil {
		return nil, err
	}

	err = c.Provide(service.NewService)
	if err != nil {
		return nil, err
	}

	return c, nil
}
