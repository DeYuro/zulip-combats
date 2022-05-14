package service

import (
	"github.com/deyuro/zulip-combats/internal/zulip"
	"github.com/sirupsen/logrus"
)

type Action interface {
	Help | Skip | Fight | Test
	run()
}

type Base struct {
	bot     zulip.BotInterface
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

type Fight struct {
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

func (f Fight) run() {

	f.bot.SendPrivateMessage(HELP, f.message.SenderEmail)
}

func runAction[T Action](action T) {
	action.run()
}

//func (b Base) getState() (state, error) {
//	s, ok := states[b.message.SenderEmail];
//	if !ok {
//		return s, errors.New("")
//	}
//
//	return s, nil
//}
