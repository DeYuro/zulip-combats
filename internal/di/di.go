package di

import (
	"flag"
	"github.com/deyuro/zulip-combats/internal/config"
	"github.com/deyuro/zulip-combats/internal/logger"
	"github.com/deyuro/zulip-combats/internal/service"
	"github.com/deyuro/zulip-combats/internal/zulip"
	"go.uber.org/dig"
)

func BuildContainer() (*dig.Container, error) {
	var cfgFile string

	flag.StringVar(&cfgFile, "config", "config.yml", "config file path")
	flag.Parse()

	c := dig.New()
	err := c.Provide(logger.NewLogger)
	if err != nil {
		return nil, err
	}

	err = c.Provide(config.NewConfig(cfgFile))
	if err != nil {
		return nil, err
	}
	err = c.Provide(config.BindConfigurator)

	err = c.Provide(zulip.NewBot)
	err = c.Provide(zulip.BindBotInterface)

	if err != nil {
		return nil, err
	}

	err = c.Provide(service.NewService)
	if err != nil {
		return nil, err
	}

	return c, nil
}
