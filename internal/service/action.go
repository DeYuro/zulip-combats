package service

import "github.com/deyuro/zulip-combats/internal/zulip"

type Action interface {
	Help | Skip
	create() Action
	run()
}
type Skip struct {
	message zulip.EventMessage
}

type Help struct {
	message zulip.EventMessage
}

func (h *Help) create() Action {
	//TODO implement me
	panic("implement me")
}

func (h *Help) run() {
	println(h.message.SenderFullName)
}

func (s *Skip) run() {
}

func ()()  {
	
}
