package config

import (
	"github.com/jinzhu/configor"
	"github.com/pkg/errors"
)

type Configurator interface {
	GetService() Service
}

func BindConfigurator(s *Config) Configurator {
	return s
}

type Config struct {
	Service Service `required:"true" yaml:"zulip"`
}

func (c Config) GetService() Service {
	return c.Service
}

type Service struct {
	Bot Bot `required:"true" yaml:"bot"`
}

type Bot struct {
	Email      string `required:"true" yaml:"email"`
	Key        string `required:"true" yaml:"key"`
	Entrypoint string `required:"true" yaml:"entrypoint"`
}

func NewConfig(filename string) func() (*Config, error) {

	return func() (*Config, error) {
		var config Config
		if err := configor.Load(&config, filename); err != nil {
			return nil, errors.WithMessage(err, "failed to load config")
		}
		return &config, nil
	}
}
