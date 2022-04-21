package service

import "github.com/deyuro/zulip-combats/internal/zulip"

type Service struct {
	bot *zulip.Bot
}

func NewService(bot *zulip.Bot) *Service {
	return &Service{bot: bot}
}

func (s *Service) Run() error {

	c, cancel := s.bot.GetEventChan()
	defer cancel()
	for e := range c {
		println(e.Content)
	}

	return nil
}
