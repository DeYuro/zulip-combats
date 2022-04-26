package service

import (
	"github.com/deyuro/zulip-combats/internal/zulip"
	"github.com/sirupsen/logrus"
)

type Service struct {
	bot    *zulip.Bot
	logger logrus.FieldLogger
}

func NewService(bot *zulip.Bot, logger logrus.FieldLogger) *Service {
	return &Service{bot: bot, logger: logger}
}

func (s *Service) Run() error {

	c, cancel := s.bot.GetEventChan()
	defer cancel()
	for e := range c {
		s.execute(e)
	}

	return nil
}

func (s Service) execute(message zulip.EventMessage) {

	base := Base{
		bot:     s.bot,
		message: message,
		logger:  s.logger,
	}
	switch message.Content {
	case "/help":
		runAction(Help{base})
	default:
		runAction(Skip{base})
	}
}
