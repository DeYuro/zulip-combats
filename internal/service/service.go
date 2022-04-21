package service

import "github.com/deyuro/zulip-combats/internal/zulip"

type Service struct {
	bot *zulip.Bot
}

func NewService(bot *zulip.Bot) *Service {
	return &Service{bot: bot}
}

func Run() error {

	for {

	}
}

func (s *Service) getMessages() error {
	queue, err := s.bot.RegisterEventQueuePrivate()
	if err != nil {
		return err
	}

}