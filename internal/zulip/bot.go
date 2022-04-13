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
}

func NewBot(email string, key string, entrypoint string, client Doer) *Bot {
	return &Bot{email: email, key: key, entrypoint: entrypoint, client: client}
}

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

func (b *Bot) request(method, endpoint, body string) (*http.Request, error) {
	url := b.entrypoint + endpoint
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(b.email, b.key)

	return req, nil
}
