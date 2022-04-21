package zulip

import (
	"net/http"
	"strings"
)

type Bot struct {
	email      string // Login for basic auth
	key        string // Password for basic auth
	entrypoint string
	client     Doer
	queue      *Queue
}

func (b *Bot) SetQueue(queue *Queue) {
	b.queue = queue
}

func NewBot(email string, key string, entrypoint string, client Doer) *Bot {
	bot := &Bot{
		email:      email,
		key:        key,
		entrypoint: entrypoint,
		client:     client,
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
				// TODO: handle unknown error
				return
			}
			for _, em := range ems {
				out <- em
			}
		}
	}()

	return out, endFunc
}
