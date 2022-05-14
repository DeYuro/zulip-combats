package config

import (
	"github.com/jinzhu/configor"
	"github.com/pkg/errors"
)

type BotConfiger interface {
	GetEmail() string
	GetKey() string
	GetEntrypoint() string
}
type ServiceConfiger interface {
	GetBotConfig() BotConfiger
}

type AppConfiger interface {
	getService() ServiceConfiger
}

type Config struct {
	Service Service `required:"true" yaml:"zulip"`
}

func (c Config) getService() ServiceConfiger {
	return c.Service
}

type Service struct {
	Bot Bot `required:"true" yaml:"bot"`
}

func (s Service) GetBotConfig() BotConfiger {
	return s.Bot
}

type Bot struct {
	Email      string `required:"true" yaml:"email"`
	Key        string `required:"true" yaml:"key"`
	Entrypoint string `required:"true" yaml:"entrypoint"`
}

func (b Bot) GetEntrypoint() string {
	return b.Entrypoint
}

func (b Bot) GetEmail() string {
	return b.Email
}

func (b Bot) GetKey() string {
	return b.Key
}

func NewConfig(filename string) (*Config, error) {
	var config Config
	if err := configor.Load(&config, filename); err != nil {
		return nil, errors.WithMessage(err, "failed to load config")
	}
	return &config, nil
}
