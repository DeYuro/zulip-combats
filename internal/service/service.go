package service

import (
	"github.com/deyuro/zulip-combats/internal/logger"
	"github.com/deyuro/zulip-combats/internal/zulip"
)

type Service struct {
	bot    *zulip.Bot
	logger logger.AppLogger
}

func NewService(bot *zulip.Bot, logger logger.AppLogger) *Service {
	return &Service{bot: bot, logger: logger}
}

func (s *Service) Run() error {
	q, err := s.bot.RegisterEventQueuePrivate()
	if err != nil {
		return err
	}
	s.bot.SetQueue(q)
	c, cancel := s.bot.GetEventChan()
	defer cancel()
	for e := range c {
		s.execute(e)
	}

	return nil
}

func initState() {

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
	case "/test":
		runAction(Test{base})
	default:
		runAction(Skip{base})
	}
}
