package zulip

import (
	"github.com/deyuro/zulip-combats/internal/config"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type Bot struct {
	email      string // Login for basic auth
	key        string // Password for basic auth
	entrypoint string
	client     Doer
	queue      *Queue
	logger     logrus.FieldLogger
}

const containerName = `zulip_bot`

func GetContainerName() string {
	return containerName
}

func (b *Bot) SetQueue(queue *Queue) {
	b.queue = queue
}

func NewBot(config *config.Config, logger logrus.FieldLogger) *Bot {
	bot := &Bot{
		email:      config.Zulip.Bot.Email,
		key:        config.Zulip.Bot.Key,
		entrypoint: config.Zulip.Entrypoint,
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
