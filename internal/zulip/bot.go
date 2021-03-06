package zulip

import (
	"github.com/deyuro/zulip-combats/internal/config"
	"github.com/deyuro/zulip-combats/internal/logger"
	"net/http"
	"strings"
)

type BotInterface interface {
	GetEventChan() (chan EventMessage, func())
	GetStreams() ([]string, error)
	RegisterEventQueuePrivate() (*Queue, error)
	RegisterEventQueue(eventList []RegisterEventType, narrow Narrow) (*Queue, error)
	GetEvents() ([]Event, error)
	SendPrivateMessage(content, email string)
	SetQueue(queue *Queue)
}

type Bot struct {
	email      string // Login for basic auth
	key        string // Password for basic auth
	entrypoint string
	client     Doer
	queue      *Queue
	logger     logger.AppLogger
}

func BindBotInterface(b *Bot) BotInterface {
	return b
}

func (b *Bot) SetQueue(queue *Queue) {
	b.queue = queue
}
func NewBot(config config.Configurator, logger logger.AppLogger) *Bot {
	bot := &Bot{
		email:      config.GetService().Bot.Email,
		key:        config.GetService().Bot.Key,
		entrypoint: config.GetService().Bot.Entrypoint,
		client:     &http.Client{},
		logger:     logger,
	}
	return bot
}

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

func (b *Bot) request(method verb, endpoint, body string) (*http.Request, error) {
	url := b.entrypoint + endpoint
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(b.email, b.key)

	return req, nil
}

func (b *Bot) GetEventChan() (chan EventMessage, func()) {
	end := false
	endFunc := func() {
		end = true
	}

	out := make(chan EventMessage)
	go func() {
		defer close(out)
		for {
			if end {
				return
			}
			ems, err := b.GetEvents()

			if err != nil {
				b.logger.WithError(err).Debug("Can not parse events")
				continue
			}

			for _, em := range ems {
				switch em.Type {
				// skip everything except Messages
				case HeartbeatQueueType:
					b.logger.Debug("Heartbeat")
				case MessageQueueType:
					out <- em.Message
				}
			}
		}
	}()

	return out, endFunc
}
