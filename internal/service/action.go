package service

import (
	"github.com/deyuro/zulip-combats/internal/zulip"
	"github.com/sirupsen/logrus"
)

type Action interface {
	Help | Skip | Test
	run()
}

type Base struct {
	bot     *zulip.Bot
	message zulip.EventMessage
	logger  logrus.FieldLogger
}
type Skip struct {
	Base
}

type Help struct {
	Base
}
type Test struct {
	Base
}

func (h Help) run() {
	h.bot.SendPrivateMessage(HELP, h.message.SenderEmail)
}

func (t Test) run() {
	t.bot.SendPrivateMessage(TEST, t.message.SenderEmail)
}

func (s Skip) run() {
	s.logger.WithField("content", s.message.Type).Info("skipped")
}

func runAction[T Action](action T) {
	action.run()
}
