package main

import (
	"context"
	"github.com/deyuro/zulip-combats/internal/di"
	"github.com/deyuro/zulip-combats/internal/service"
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

	c, err := di.BuildContainer()

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
