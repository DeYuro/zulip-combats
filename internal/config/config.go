package config

import (
	"github.com/jinzhu/configor"
	"github.com/pkg/errors"
)

type BotConfig interface {
}

type Config struct {
	Zulip Service `required:"true" yaml:"zulip"`
}

type Service struct {
	Bot        Bot    `required:"true" yaml:"bot"`
	Entrypoint string `required:"true" yaml:"entrypoint"`
}

type Bot struct {
	Email string `required:"true" yaml:"email"`
	Key   string `required:"true" yaml:"key"`
}

func NewConfig(filename string) (*Config, error) {
	var config Config
	if err := configor.Load(&config, filename); err != nil {
		return nil, errors.WithMessage(err, "failed to load config")
	}
	return &config, nil
}
